package rosetta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"

	auth "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/tendermint/tendermint/crypto"

	"github.com/btcsuite/btcd/btcec"
	tmcoretypes "github.com/tendermint/tendermint/rpc/core/types"

	crgtypes "github.com/cosmos/cosmos-sdk/server/rosetta/lib/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	rosettatypes "github.com/coinbase/rosetta-sdk-go/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	crgerrs "github.com/cosmos/cosmos-sdk/server/rosetta/lib/errors"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)







//

type Converter interface {


	ToSDK() ToSDKConverter


	ToRosetta() ToRosettaConverter
}




type ToRosettaConverter interface {

	BlockResponse(block *tmcoretypes.ResultBlock) crgtypes.BlockResponse

	BeginBlockTxHash(blockHash []byte) string

	EndBlockTxHash(blockHash []byte) string

	Amounts(ownedCoins []sdk.Coin, availableCoins sdk.Coins) []*rosettatypes.Amount

	Ops(status string, msg sdk.Msg) ([]*rosettatypes.Operation, error)

	OpsAndSigners(txBytes []byte) (ops []*rosettatypes.Operation, signers []*rosettatypes.AccountIdentifier, err error)

	Meta(msg sdk.Msg) (meta map[string]interface{}, err error)

	SignerData(anyAccount *codectypes.Any) (*SignerData, error)

	SigningComponents(tx authsigning.Tx, metadata *ConstructionMetadata, rosPubKeys []*rosettatypes.PublicKey) (txBytes []byte, payloadsToSign []*rosettatypes.SigningPayload, err error)

	Tx(rawTx tmtypes.Tx, txResult *abci.ResponseDeliverTx) (*rosettatypes.Transaction, error)

	TxIdentifiers(txs []tmtypes.Tx) []*rosettatypes.TransactionIdentifier

	BalanceOps(status string, events []abci.Event) []*rosettatypes.Operation

	SyncStatus(status *tmcoretypes.ResultStatus) *rosettatypes.SyncStatus

	Peers(peers []tmcoretypes.Peer) []*rosettatypes.Peer
}




type ToSDKConverter interface {

	UnsignedTx(ops []*rosettatypes.Operation) (tx authsigning.Tx, err error)


	SignedTx(txBytes []byte, signatures []*rosettatypes.Signature) (signedTxBytes []byte, err error)

	Msg(meta map[string]interface{}, msg sdk.Msg) (err error)


	HashToTxType(hashBytes []byte) (txType TransactionType, realHash []byte)

	PubKey(pk *rosettatypes.PublicKey) (cryptotypes.PubKey, error)
}

type converter struct {
	newTxBuilder    func() sdkclient.TxBuilder
	txBuilderFromTx func(tx sdk.Tx) (sdkclient.TxBuilder, error)
	txDecode        sdk.TxDecoder
	txEncode        sdk.TxEncoder
	bytesToSign     func(tx authsigning.Tx, signerData authsigning.SignerData) (b []byte, err error)
	ir              codectypes.InterfaceRegistry
	cdc             *codec.ProtoCodec
}

func NewConverter(cdc *codec.ProtoCodec, ir codectypes.InterfaceRegistry, cfg sdkclient.TxConfig) Converter {
	return converter{
		newTxBuilder:    cfg.NewTxBuilder,
		txBuilderFromTx: cfg.WrapTxBuilder,
		txDecode:        cfg.TxDecoder(),
		txEncode:        cfg.TxEncoder(),
		bytesToSign: func(tx authsigning.Tx, signerData authsigning.SignerData) (b []byte, err error) {
			bytesToSign, err := cfg.SignModeHandler().GetSignBytes(signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signerData, tx)
			if err != nil {
				return nil, err
			}

			return crypto.Sha256(bytesToSign), nil
		},
		ir:  ir,
		cdc: cdc,
	}
}

func (c converter) ToSDK() ToSDKConverter {
	return c
}

func (c converter) ToRosetta() ToRosettaConverter {
	return c
}


func (c converter) UnsignedTx(ops []*rosettatypes.Operation) (tx authsigning.Tx, err error) {
	builder := c.newTxBuilder()

	var msgs []sdk.Msg

	for i := 0; i < len(ops); i++ {
		op := ops[i]

		protoMessage, err := c.ir.Resolve(op.Type)
		if err != nil {
			return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, "operation not found: "+op.Type)
		}

		msg, ok := protoMessage.(sdk.Msg)
		if !ok {
			return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, "operation is not a valid supported sdk.Msg: "+op.Type)
		}

		err = c.Msg(op.Metadata, msg)
		if err != nil {
			return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
		}


		if err = msg.ValidateBasic(); err != nil {
			return nil, crgerrs.WrapError(
				crgerrs.ErrBadArgument,
				fmt.Sprintf("validation of operation at index %d failed: %s", op.OperationIdentifier.Index, err),
			)
		}
		signers := msg.GetSigners()

		if len(signers) == 0 {
			return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, fmt.Sprintf("operation at index %d got no signers", op.OperationIdentifier.Index))
		}

		msgs = append(msgs, msg)

		if len(signers) == 1 {
			continue
		}






		for j := 0; j < len(signers)-1; j++ {
			skipOp := ops[i+j]

			if skipOp.Type != op.Type {
				return nil, crgerrs.WrapError(
					crgerrs.ErrBadArgument,
					fmt.Sprintf("operation at index %d should have had type %s got: %s", i+j, op.Type, skipOp.Type),
				)
			}

			if !reflect.DeepEqual(op.Metadata, skipOp.Metadata) {
				return nil, crgerrs.WrapError(
					crgerrs.ErrBadArgument,
					fmt.Sprintf("operation at index %d should have had metadata equal to %#v, got: %#v", i+j, op.Metadata, skipOp.Metadata))
			}

			i++
		}
	}

	if err := builder.SetMsgs(msgs...); err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, err.Error())
	}

	return builder.GetTx(), nil

}


func (c converter) Msg(meta map[string]interface{}, msg sdk.Msg) error {
	metaBytes, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	return c.cdc.UnmarshalJSON(metaBytes, msg)
}

func (c converter) Meta(msg sdk.Msg) (meta map[string]interface{}, err error) {
	b, err := c.cdc.MarshalJSON(msg)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	err = json.Unmarshal(b, &meta)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	return
}




func (c converter) Ops(status string, msg sdk.Msg) ([]*rosettatypes.Operation, error) {
	opName := sdk.MsgTypeURL(msg)

	meta, err := c.Meta(msg)
	if err != nil {
		return nil, err
	}

	ops := make([]*rosettatypes.Operation, len(msg.GetSigners()))
	for i, signer := range msg.GetSigners() {
		op := &rosettatypes.Operation{
			Type:     opName,
			Status:   &status,
			Account:  &rosettatypes.AccountIdentifier{Address: signer.String()},
			Metadata: meta,
		}

		ops[i] = op
	}

	return ops, nil
}


func (c converter) Tx(rawTx tmtypes.Tx, txResult *abci.ResponseDeliverTx) (*rosettatypes.Transaction, error) {

	tx, err := c.txDecode(rawTx)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}



	status := StatusTxSuccess
	switch txResult {



	case nil:
		status = ""

	default:
		if txResult.Code != abci.CodeTypeOK {
			status = StatusTxReverted
		}
	}

	msgs := tx.GetMsgs()
	var rawTxOps []*rosettatypes.Operation

	for _, msg := range msgs {
		ops, err := c.Ops(status, msg)
		if err != nil {
			return nil, err
		}
		rawTxOps = append(rawTxOps, ops...)
	}


	var balanceOps []*rosettatypes.Operation

	if txResult != nil {
		balanceOps = c.BalanceOps(status, txResult.Events)
	}


	totalOps := AddOperationIndexes(rawTxOps, balanceOps)

	return &rosettatypes.Transaction{
		TransactionIdentifier: &rosettatypes.TransactionIdentifier{Hash: fmt.Sprintf("%X", rawTx.Hash())},
		Operations:            totalOps,
	}, nil
}

func (c converter) BalanceOps(status string, events []abci.Event) []*rosettatypes.Operation {
	var ops []*rosettatypes.Operation

	for _, e := range events {
		balanceOps, ok := sdkEventToBalanceOperations(status, e)
		if !ok {
			continue
		}
		ops = append(ops, balanceOps...)
	}

	return ops
}





func sdkEventToBalanceOperations(status string, event abci.Event) (operations []*rosettatypes.Operation, isBalanceEvent bool) {

	var (
		accountIdentifier string
		coinChange        sdk.Coins
		isSub             bool
	)

	switch event.Type {
	default:
		return nil, false
	case banktypes.EventTypeCoinSpent:
		spender, err := sdk.AccAddressFromBech32((string)(event.Attributes[0].Value))
		if err != nil {
			panic(err)
		}
		coins, err := sdk.ParseCoinsNormalized((string)(event.Attributes[1].Value))
		if err != nil {
			panic(err)
		}

		isSub = true
		coinChange = coins
		accountIdentifier = spender.String()

	case banktypes.EventTypeCoinReceived:
		receiver, err := sdk.AccAddressFromBech32((string)(event.Attributes[0].Value))
		if err != nil {
			panic(err)
		}
		coins, err := sdk.ParseCoinsNormalized((string)(event.Attributes[1].Value))
		if err != nil {
			panic(err)
		}

		isSub = false
		coinChange = coins
		accountIdentifier = receiver.String()



	case banktypes.EventTypeCoinBurn:
		coins, err := sdk.ParseCoinsNormalized((string)(event.Attributes[1].Value))
		if err != nil {
			panic(err)
		}

		coinChange = coins
		accountIdentifier = BurnerAddressIdentifier
	}

	operations = make([]*rosettatypes.Operation, len(coinChange))

	for i, coin := range coinChange {

		value := coin.Amount.String()


		if isSub {
			value = "-" + value
		}

		op := &rosettatypes.Operation{
			Type:    event.Type,
			Status:  &status,
			Account: &rosettatypes.AccountIdentifier{Address: accountIdentifier},
			Amount: &rosettatypes.Amount{
				Value: value,
				Currency: &rosettatypes.Currency{
					Symbol:   coin.Denom,
					Decimals: 0,
				},
			},
		}

		operations[i] = op
	}
	return operations, true
}


func (c converter) Amounts(ownedCoins []sdk.Coin, availableCoins sdk.Coins) []*rosettatypes.Amount {
	amounts := make([]*rosettatypes.Amount, len(availableCoins))
	ownedCoinsMap := make(map[string]sdk.Int, len(availableCoins))

	for _, ownedCoin := range ownedCoins {
		ownedCoinsMap[ownedCoin.Denom] = ownedCoin.Amount
	}

	for i, coin := range availableCoins {
		value, owned := ownedCoinsMap[coin.Denom]
		if !owned {
			amounts[i] = &rosettatypes.Amount{
				Value: sdk.NewInt(0).String(),
				Currency: &rosettatypes.Currency{
					Symbol: coin.Denom,
				},
			}
			continue
		}
		amounts[i] = &rosettatypes.Amount{
			Value: value.String(),
			Currency: &rosettatypes.Currency{
				Symbol: coin.Denom,
			},
		}
	}

	return amounts
}



func AddOperationIndexes(msgOps []*rosettatypes.Operation, balanceOps []*rosettatypes.Operation) (finalOps []*rosettatypes.Operation) {
	lenMsgOps := len(msgOps)
	lenBalanceOps := len(balanceOps)
	finalOps = make([]*rosettatypes.Operation, 0, lenMsgOps+lenBalanceOps)

	var currentIndex int64

	for _, op := range msgOps {
		op.OperationIdentifier = &rosettatypes.OperationIdentifier{
			Index: currentIndex,
		}

		finalOps = append(finalOps, op)
		currentIndex++
	}


	for _, op := range balanceOps {
		op.OperationIdentifier = &rosettatypes.OperationIdentifier{
			Index: currentIndex,
		}

		finalOps = append(finalOps, op)
		currentIndex++
	}

	return finalOps
}




func (c converter) EndBlockTxHash(hash []byte) string {
	final := append([]byte{EndBlockHashStart}, hash...)
	return fmt.Sprintf("%X", final)
}




func (c converter) BeginBlockTxHash(hash []byte) string {
	final := append([]byte{BeginBlockHashStart}, hash...)
	return fmt.Sprintf("%X", final)
}



func (c converter) HashToTxType(hashBytes []byte) (txType TransactionType, realHash []byte) {
	switch len(hashBytes) {
	case DeliverTxSize:
		return DeliverTxTx, hashBytes

	case BeginEndBlockTxSize:
		switch hashBytes[0] {
		case BeginBlockHashStart:
			return BeginBlockTx, hashBytes[1:]
		case EndBlockHashStart:
			return EndBlockTx, hashBytes[1:]
		default:
			return UnrecognizedTx, nil
		}

	default:
		return UnrecognizedTx, nil
	}
}


func (c converter) SyncStatus(status *tmcoretypes.ResultStatus) *rosettatypes.SyncStatus {

	var stage = StatusPeerSynced
	if status.SyncInfo.CatchingUp {
		stage = StatusPeerSyncing
	}

	return &rosettatypes.SyncStatus{
		CurrentIndex: &status.SyncInfo.LatestBlockHeight,
		TargetIndex:  nil,
		Stage:        &stage,
	}
}


func (c converter) TxIdentifiers(txs []tmtypes.Tx) []*rosettatypes.TransactionIdentifier {
	converted := make([]*rosettatypes.TransactionIdentifier, len(txs))
	for i, tx := range txs {
		converted[i] = &rosettatypes.TransactionIdentifier{Hash: fmt.Sprintf("%X", tx.Hash())}
	}

	return converted
}


func (c converter) BlockResponse(block *tmcoretypes.ResultBlock) crgtypes.BlockResponse {
	var parentBlock *rosettatypes.BlockIdentifier

	switch block.Block.Height {
	case 1:
		parentBlock = &rosettatypes.BlockIdentifier{
			Index: 1,
			Hash:  fmt.Sprintf("%X", block.BlockID.Hash.Bytes()),
		}
	default:
		parentBlock = &rosettatypes.BlockIdentifier{
			Index: block.Block.Height - 1,
			Hash:  fmt.Sprintf("%X", block.Block.LastBlockID.Hash.Bytes()),
		}
	}
	return crgtypes.BlockResponse{
		Block: &rosettatypes.BlockIdentifier{
			Index: block.Block.Height,
			Hash:  block.Block.Hash().String(),
		},
		ParentBlock:          parentBlock,
		MillisecondTimestamp: timeToMilliseconds(block.Block.Time),
		TxCount:              int64(len(block.Block.Txs)),
	}
}


func (c converter) Peers(peers []tmcoretypes.Peer) []*rosettatypes.Peer {
	converted := make([]*rosettatypes.Peer, len(peers))

	for i, peer := range peers {
		converted[i] = &rosettatypes.Peer{
			PeerID: peer.NodeInfo.Moniker,
			Metadata: map[string]interface{}{
				"addr": peer.NodeInfo.ListenAddr,
			},
		}
	}

	return converted
}



func (c converter) OpsAndSigners(txBytes []byte) (ops []*rosettatypes.Operation, signers []*rosettatypes.AccountIdentifier, err error) {

	rosTx, err := c.ToRosetta().Tx(txBytes, nil)
	if err != nil {
		return nil, nil, err
	}
	ops = rosTx.Operations


	sdkTx, err := c.txDecode(txBytes)
	if err != nil {
		return nil, nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	txBuilder, err := c.txBuilderFromTx(sdkTx)
	if err != nil {
		return nil, nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	for _, signer := range txBuilder.GetTx().GetSigners() {
		signers = append(signers, &rosettatypes.AccountIdentifier{
			Address: signer.String(),
		})
	}

	return
}

func (c converter) SignedTx(txBytes []byte, signatures []*rosettatypes.Signature) (signedTxBytes []byte, err error) {
	rawTx, err := c.txDecode(txBytes)
	if err != nil {
		return nil, err
	}

	txBuilder, err := c.txBuilderFromTx(rawTx)
	if err != nil {
		return nil, err
	}

	notSignedSigs, err := txBuilder.GetTx().GetSignaturesV2() //
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	if len(notSignedSigs) != len(signatures) {
		return nil, crgerrs.WrapError(
			crgerrs.ErrInvalidTransaction,
			fmt.Sprintf("expected transaction to have signers data matching the provided signatures: %d <-> %d", len(notSignedSigs), len(signatures)))
	}

	signedSigs := make([]signing.SignatureV2, len(notSignedSigs))
	for i, signature := range signatures {

		signedSigs[i] = signing.SignatureV2{
			PubKey: notSignedSigs[i].PubKey,
			Data: &signing.SingleSignatureData{
				SignMode:  signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
				Signature: signature.Bytes,
			},
			Sequence: notSignedSigs[i].Sequence,
		}
	}

	if err = txBuilder.SetSignatures(signedSigs...); err != nil {
		return nil, err
	}

	txBytes, err = c.txEncode(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil
}

func (c converter) PubKey(pubKey *rosettatypes.PublicKey) (cryptotypes.PubKey, error) {
	if pubKey.CurveType != "secp256k1" {
		return nil, crgerrs.WrapError(crgerrs.ErrUnsupportedCurve, "only secp256k1 supported")
	}

	cmp, err := btcec.ParsePubKey(pubKey.Bytes, btcec.S256())
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, err.Error())
	}

	compressedPublicKey := make([]byte, secp256k1.PubKeySize)
	copy(compressedPublicKey, cmp.SerializeCompressed())

	pk := &secp256k1.PubKey{Key: compressedPublicKey}

	return pk, nil
}


func (c converter) SigningComponents(tx authsigning.Tx, metadata *ConstructionMetadata, rosPubKeys []*rosettatypes.PublicKey) (txBytes []byte, payloadsToSign []*rosettatypes.SigningPayload, err error) {


	feeAmount, err := sdk.ParseCoinsNormalized(metadata.GasPrice)
	if err != nil {
		return nil, nil, crgerrs.WrapError(crgerrs.ErrBadArgument, err.Error())
	}

	signers := tx.GetSigners()


	if len(metadata.SignersData) != len(signers) || len(signers) != len(rosPubKeys) {
		return nil, nil, crgerrs.WrapError(crgerrs.ErrBadArgument, "signers data and account identifiers mismatch")
	}


	builder, err := c.txBuilderFromTx(tx)
	if err != nil {
		return nil, nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}
	builder.SetFeeAmount(feeAmount)
	builder.SetGasLimit(metadata.GasLimit)
	builder.SetMemo(metadata.Memo)


	partialSignatures := make([]signing.SignatureV2, len(signers))
	payloadsToSign = make([]*rosettatypes.SigningPayload, len(signers))


	for i, signer := range signers {


		pubKey, err := c.ToSDK().PubKey(rosPubKeys[0])
		if err != nil {
			return nil, nil, err
		}
		if !bytes.Equal(pubKey.Address().Bytes(), signer.Bytes()) {
			return nil, nil, crgerrs.WrapError(
				crgerrs.ErrBadArgument,
				fmt.Sprintf("public key at index %d does not match the expected transaction signer: %X <-> %X", i, rosPubKeys[i].Bytes, signer.Bytes()),
			)
		}


		signerData := authsigning.SignerData{
			ChainID:       metadata.ChainID,
			AccountNumber: metadata.SignersData[i].AccountNumber,
			Sequence:      metadata.SignersData[i].Sequence,
		}


		signBytes, err := c.bytesToSign(tx, signerData)
		if err != nil {
			return nil, nil, crgerrs.WrapError(crgerrs.ErrUnknown, fmt.Sprintf("unable to sign tx: %s", err.Error()))
		}


		payloadsToSign[i] = &rosettatypes.SigningPayload{
			AccountIdentifier: &rosettatypes.AccountIdentifier{Address: signer.String()},
			Bytes:             signBytes,
			SignatureType:     rosettatypes.Ecdsa,
		}


		partialSignatures[i] = signing.SignatureV2{
			PubKey:   pubKey,
			Data:     &signing.SingleSignatureData{},
			Sequence: metadata.SignersData[i].Sequence,
		}

	}




	err = builder.SetSignatures(partialSignatures...)
	if err != nil {
		return nil, nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}


	txBytes, err = c.txEncode(builder.GetTx())
	if err != nil {
		return nil, nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	return txBytes, payloadsToSign, nil
}


func (c converter) SignerData(anyAccount *codectypes.Any) (*SignerData, error) {
	var acc auth.AccountI
	err := c.ir.UnpackAny(anyAccount, &acc)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, err.Error())
	}

	return &SignerData{
		AccountNumber: acc.GetAccountNumber(),
		Sequence:      acc.GetSequence(),
	}, nil
}
