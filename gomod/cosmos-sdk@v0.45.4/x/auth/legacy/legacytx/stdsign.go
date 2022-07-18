package legacytx

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)




type LegacyMsg interface {
	sdk.Msg


	GetSignBytes() []byte



	Route() string



	Type() string
}






type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number" yaml:"account_number"`
	Sequence      uint64            `json:"sequence" yaml:"sequence"`
	TimeoutHeight uint64            `json:"timeout_height,omitempty" yaml:"timeout_height"`
	ChainID       string            `json:"chain_id" yaml:"chain_id"`
	Memo          string            `json:"memo" yaml:"memo"`
	Fee           json.RawMessage   `json:"fee" yaml:"fee"`
	Msgs          []json.RawMessage `json:"msgs" yaml:"msgs"`
}


func StdSignBytes(chainID string, accnum, sequence, timeout uint64, fee StdFee, msgs []sdk.Msg, memo string) []byte {
	msgsBytes := make([]json.RawMessage, 0, len(msgs))
	for _, msg := range msgs {
		legacyMsg, ok := msg.(LegacyMsg)
		if !ok {
			panic(fmt.Errorf("expected %T when using amino JSON", (*LegacyMsg)(nil)))
		}

		msgsBytes = append(msgsBytes, json.RawMessage(legacyMsg.GetSignBytes()))
	}

	bz, err := legacy.Cdc.MarshalJSON(StdSignDoc{
		AccountNumber: accnum,
		ChainID:       chainID,
		Fee:           json.RawMessage(fee.Bytes()),
		Memo:          memo,
		Msgs:          msgsBytes,
		Sequence:      sequence,
		TimeoutHeight: timeout,
	})
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(bz)
}


type StdSignature struct {
	cryptotypes.PubKey `json:"pub_key" yaml:"pub_key"`
	Signature          []byte                          `json:"signature" yaml:"signature"`
}


func NewStdSignature(pk cryptotypes.PubKey, sig []byte) StdSignature {
	return StdSignature{PubKey: pk, Signature: sig}
}


func (ss StdSignature) GetSignature() []byte {
	return ss.Signature
}



func (ss StdSignature) GetPubKey() cryptotypes.PubKey {
	return ss.PubKey
}


func (ss StdSignature) MarshalYAML() (interface{}, error) {
	pk := ""
	if ss.PubKey != nil {
		pk = ss.PubKey.String()
	}

	bz, err := yaml.Marshal(struct {
		PubKey    string
		Signature string
	}{
		pk,
		fmt.Sprintf("%X", ss.Signature),
	})
	if err != nil {
		return nil, err
	}

	return string(bz), err
}

func (ss StdSignature) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return codectypes.UnpackInterfaces(ss.PubKey, unpacker)
}


func StdSignatureToSignatureV2(cdc *codec.LegacyAmino, sig StdSignature) (signing.SignatureV2, error) {
	pk := sig.GetPubKey()
	data, err := pubKeySigToSigData(cdc, pk, sig.Signature)
	if err != nil {
		return signing.SignatureV2{}, err
	}

	return signing.SignatureV2{
		PubKey: pk,
		Data:   data,
	}, nil
}

func pubKeySigToSigData(cdc *codec.LegacyAmino, key cryptotypes.PubKey, sig []byte) (signing.SignatureData, error) {
	multiPK, ok := key.(multisig.PubKey)
	if !ok {
		return &signing.SingleSignatureData{
			SignMode:  signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
			Signature: sig,
		}, nil
	}
	var multiSig multisig.AminoMultisignature
	err := cdc.Unmarshal(sig, &multiSig)
	if err != nil {
		return nil, err
	}

	sigs := multiSig.Sigs
	sigDatas := make([]signing.SignatureData, len(sigs))
	pubKeys := multiPK.GetPubKeys()
	bitArray := multiSig.BitArray
	n := multiSig.BitArray.Count()
	signatures := multisig.NewMultisig(n)
	sigIdx := 0
	for i := 0; i < n; i++ {
		if bitArray.GetIndex(i) {
			data, err := pubKeySigToSigData(cdc, pubKeys[i], multiSig.Sigs[sigIdx])
			if err != nil {
				return nil, sdkerrors.Wrapf(err, "Unable to convert Signature to SigData %d", sigIdx)
			}

			sigDatas[sigIdx] = data
			multisig.AddSignature(signatures, data, sigIdx)
			sigIdx++
		}
	}

	return signatures, nil
}
