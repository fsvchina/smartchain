package legacytx

import (
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)


var (
	_ sdk.Tx                             = (*StdTx)(nil)
	_ sdk.TxWithMemo                     = (*StdTx)(nil)
	_ sdk.FeeTx                          = (*StdTx)(nil)
	_ codectypes.UnpackInterfacesMessage = (*StdTx)(nil)

	_ codectypes.UnpackInterfacesMessage = (*StdSignature)(nil)
)





type StdFee struct {
	Amount sdk.Coins `json:"amount" yaml:"amount"`
	Gas    uint64    `json:"gas" yaml:"gas"`
}


func NewStdFee(gas uint64, amount sdk.Coins) StdFee {
	return StdFee{
		Amount: amount,
		Gas:    gas,
	}
}


func (fee StdFee) GetGas() uint64 {
	return fee.Gas
}


func (fee StdFee) GetAmount() sdk.Coins {
	return fee.Amount
}


func (fee StdFee) Bytes() []byte {
	if len(fee.Amount) == 0 {
		fee.Amount = sdk.NewCoins()
	}

	bz, err := legacy.Cdc.MarshalJSON(fee)
	if err != nil {
		panic(err)
	}

	return bz
}


//



func (fee StdFee) GasPrices() sdk.DecCoins {
	return sdk.NewDecCoinsFromCoins(fee.Amount...).QuoDec(sdk.NewDec(int64(fee.Gas)))
}





type StdTx struct {
	Msgs          []sdk.Msg      `json:"msg" yaml:"msg"`
	Fee           StdFee         `json:"fee" yaml:"fee"`
	Signatures    []StdSignature `json:"signatures" yaml:"signatures"`
	Memo          string         `json:"memo" yaml:"memo"`
	TimeoutHeight uint64         `json:"timeout_height" yaml:"timeout_height"`
}


func NewStdTx(msgs []sdk.Msg, fee StdFee, sigs []StdSignature, memo string) StdTx {
	return StdTx{
		Msgs:       msgs,
		Fee:        fee,
		Signatures: sigs,
		Memo:       memo,
	}
}


func (tx StdTx) GetMsgs() []sdk.Msg { return tx.Msgs }



func (tx StdTx) ValidateBasic() error {
	stdSigs := tx.GetSignatures()

	if tx.Fee.Gas > txtypes.MaxGasWanted {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"invalid gas supplied; %d > %d", tx.Fee.Gas, txtypes.MaxGasWanted,
		)
	}
	if tx.Fee.Amount.IsAnyNegative() {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFee,
			"invalid fee provided: %s", tx.Fee.Amount,
		)
	}
	if len(stdSigs) == 0 {
		return sdkerrors.ErrNoSignatures
	}
	if len(stdSigs) != len(tx.GetSigners()) {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnauthorized,
			"wrong number of signers; expected %d, got %d", len(tx.GetSigners()), len(stdSigs),
		)
	}

	return nil
}




func (tx *StdTx) AsAny() *codectypes.Any {
	return codectypes.UnsafePackAny(tx)
}






func (tx StdTx) GetSigners() []sdk.AccAddress {
	var signers []sdk.AccAddress
	seen := map[string]bool{}

	for _, msg := range tx.GetMsgs() {
		for _, addr := range msg.GetSigners() {
			if !seen[addr.String()] {
				signers = append(signers, addr)
				seen[addr.String()] = true
			}
		}
	}

	return signers
}


func (tx StdTx) GetMemo() string { return tx.Memo }


func (tx StdTx) GetTimeoutHeight() uint64 {
	return tx.TimeoutHeight
}








func (tx StdTx) GetSignatures() [][]byte {
	sigs := make([][]byte, len(tx.Signatures))
	for i, stdSig := range tx.Signatures {
		sigs[i] = stdSig.Signature
	}
	return sigs
}


func (tx StdTx) GetSignaturesV2() ([]signing.SignatureV2, error) {
	res := make([]signing.SignatureV2, len(tx.Signatures))

	for i, sig := range tx.Signatures {
		var err error
		res[i], err = StdSignatureToSignatureV2(legacy.Cdc, sig)
		if err != nil {
			return nil, sdkerrors.Wrapf(err, "Unable to convert signature %v to V2", sig)
		}
	}

	return res, nil
}



func (tx StdTx) GetPubKeys() ([]cryptotypes.PubKey, error) {
	pks := make([]cryptotypes.PubKey, len(tx.Signatures))

	for i, stdSig := range tx.Signatures {
		pks[i] = stdSig.GetPubKey()
	}

	return pks, nil
}


func (tx StdTx) GetGas() uint64 { return tx.Fee.Gas }


func (tx StdTx) GetFee() sdk.Coins { return tx.Fee.Amount }




func (tx StdTx) FeePayer() sdk.AccAddress {
	if tx.GetSigners() != nil {
		return tx.GetSigners()[0]
	}
	return sdk.AccAddress{}
}


func (tx StdTx) FeeGranter() sdk.AccAddress {
	return nil
}

func (tx StdTx) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, m := range tx.Msgs {
		err := codectypes.UnpackInterfaces(m, unpacker)
		if err != nil {
			return err
		}
	}


	for _, s := range tx.Signatures {
		err := s.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}

	return nil
}
