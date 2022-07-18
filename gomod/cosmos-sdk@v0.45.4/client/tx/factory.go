package tx

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)



type Factory struct {
	keybase            keyring.Keyring
	txConfig           client.TxConfig
	accountRetriever   client.AccountRetriever
	accountNumber      uint64
	sequence           uint64
	gas                uint64
	timeoutHeight      uint64
	gasAdjustment      float64
	chainID            string
	memo               string
	fees               sdk.Coins
	gasPrices          sdk.DecCoins
	signMode           signing.SignMode
	simulateAndExecute bool
}


func NewFactoryCLI(clientCtx client.Context, flagSet *pflag.FlagSet) Factory {
	signModeStr := clientCtx.SignModeStr

	signMode := signing.SignMode_SIGN_MODE_UNSPECIFIED
	switch signModeStr {
	case flags.SignModeDirect:
		signMode = signing.SignMode_SIGN_MODE_DIRECT
	case flags.SignModeLegacyAminoJSON:
		signMode = signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	case flags.SignModeEIP191:
		signMode = signing.SignMode_SIGN_MODE_EIP_191
	}

	accNum, _ := flagSet.GetUint64(flags.FlagAccountNumber)
	accSeq, _ := flagSet.GetUint64(flags.FlagSequence)
	gasAdj, _ := flagSet.GetFloat64(flags.FlagGasAdjustment)
	memo, _ := flagSet.GetString(flags.FlagNote)
	timeoutHeight, _ := flagSet.GetUint64(flags.FlagTimeoutHeight)

	gasStr, _ := flagSet.GetString(flags.FlagGas)
	gasSetting, _ := flags.ParseGasSetting(gasStr)

	f := Factory{
		txConfig:           clientCtx.TxConfig,
		accountRetriever:   clientCtx.AccountRetriever,
		keybase:            clientCtx.Keyring,
		chainID:            clientCtx.ChainID,
		gas:                gasSetting.Gas,
		simulateAndExecute: gasSetting.Simulate,
		accountNumber:      accNum,
		sequence:           accSeq,
		timeoutHeight:      timeoutHeight,
		gasAdjustment:      gasAdj,
		memo:               memo,
		signMode:           signMode,
	}

	feesStr, _ := flagSet.GetString(flags.FlagFees)
	f = f.WithFees(feesStr)

	gasPricesStr, _ := flagSet.GetString(flags.FlagGasPrices)
	f = f.WithGasPrices(gasPricesStr)

	return f
}

func (f Factory) AccountNumber() uint64                     { return f.accountNumber }
func (f Factory) Sequence() uint64                          { return f.sequence }
func (f Factory) Gas() uint64                               { return f.gas }
func (f Factory) GasAdjustment() float64                    { return f.gasAdjustment }
func (f Factory) Keybase() keyring.Keyring                  { return f.keybase }
func (f Factory) ChainID() string                           { return f.chainID }
func (f Factory) Memo() string                              { return f.memo }
func (f Factory) Fees() sdk.Coins                           { return f.fees }
func (f Factory) GasPrices() sdk.DecCoins                   { return f.gasPrices }
func (f Factory) AccountRetriever() client.AccountRetriever { return f.accountRetriever }
func (f Factory) TimeoutHeight() uint64                     { return f.timeoutHeight }



func (f Factory) SimulateAndExecute() bool { return f.simulateAndExecute }


func (f Factory) WithTxConfig(g client.TxConfig) Factory {
	f.txConfig = g
	return f
}


func (f Factory) WithAccountRetriever(ar client.AccountRetriever) Factory {
	f.accountRetriever = ar
	return f
}


func (f Factory) WithChainID(chainID string) Factory {
	f.chainID = chainID
	return f
}


func (f Factory) WithGas(gas uint64) Factory {
	f.gas = gas
	return f
}


func (f Factory) WithFees(fees string) Factory {
	parsedFees, err := sdk.ParseCoinsNormalized(fees)
	if err != nil {
		panic(err)
	}

	f.fees = parsedFees
	return f
}


func (f Factory) WithGasPrices(gasPrices string) Factory {
	parsedGasPrices, err := sdk.ParseDecCoins(gasPrices)
	if err != nil {
		panic(err)
	}

	f.gasPrices = parsedGasPrices
	return f
}


func (f Factory) WithKeybase(keybase keyring.Keyring) Factory {
	f.keybase = keybase
	return f
}


func (f Factory) WithSequence(sequence uint64) Factory {
	f.sequence = sequence
	return f
}


func (f Factory) WithMemo(memo string) Factory {
	f.memo = memo
	return f
}


func (f Factory) WithAccountNumber(accnum uint64) Factory {
	f.accountNumber = accnum
	return f
}


func (f Factory) WithGasAdjustment(gasAdj float64) Factory {
	f.gasAdjustment = gasAdj
	return f
}



func (f Factory) WithSimulateAndExecute(sim bool) Factory {
	f.simulateAndExecute = sim
	return f
}


func (f Factory) SignMode() signing.SignMode {
	return f.signMode
}


func (f Factory) WithSignMode(mode signing.SignMode) Factory {
	f.signMode = mode
	return f
}


func (f Factory) WithTimeoutHeight(height uint64) Factory {
	f.timeoutHeight = height
	return f
}



func (f Factory) BuildUnsignedTx(msgs ...sdk.Msg) (client.TxBuilder, error) {
	if f.chainID == "" {
		return nil, fmt.Errorf("chain ID required but not specified")
	}

	fees := f.fees

	if !f.gasPrices.IsZero() {
		if !fees.IsZero() {
			return nil, errors.New("cannot provide both fees and gas prices")
		}

		glDec := sdk.NewDec(int64(f.gas))

		
		
		fees = make(sdk.Coins, len(f.gasPrices))

		for i, gp := range f.gasPrices {
			fee := gp.Amount.Mul(glDec)
			fees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
		}
	}

	tx := f.txConfig.NewTxBuilder()

	if err := tx.SetMsgs(msgs...); err != nil {
		return nil, err
	}

	tx.SetMemo(f.memo)
	tx.SetFeeAmount(fees)
	tx.SetGasLimit(f.gas)
	tx.SetTimeoutHeight(f.TimeoutHeight())

	return tx, nil
}





func (f Factory) PrintUnsignedTx(clientCtx client.Context, msgs ...sdk.Msg) error {
	if f.SimulateAndExecute() {
		if clientCtx.Offline {
			return errors.New("cannot estimate gas in offline mode")
		}

		_, adjusted, err := CalculateGas(clientCtx, f, msgs...)
		if err != nil {
			return err
		}

		f = f.WithGas(adjusted)
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", GasEstimateResponse{GasEstimate: f.Gas()})
	}

	unsignedTx, err := f.BuildUnsignedTx(msgs...)
	if err != nil {
		return err
	}

	json, err := clientCtx.TxConfig.TxJSONEncoder()(unsignedTx.GetTx())
	if err != nil {
		return err
	}

	return clientCtx.PrintString(fmt.Sprintf("%s\n", json))
}




func (f Factory) BuildSimTx(msgs ...sdk.Msg) ([]byte, error) {
	txb, err := f.BuildUnsignedTx(msgs...)
	if err != nil {
		return nil, err
	}

	
	

	var pk cryptotypes.PubKey = &secp256k1.PubKey{} 

	if f.keybase != nil {
		infos, _ := f.keybase.List()
		if len(infos) == 0 {
			return nil, errors.New("cannot build signature for simulation, key infos slice is empty")
		}

		
		pk = infos[0].GetPubKey()
	}

	
	
	sig := signing.SignatureV2{
		PubKey: pk,
		Data: &signing.SingleSignatureData{
			SignMode: f.signMode,
		},
		Sequence: f.Sequence(),
	}
	if err := txb.SetSignatures(sig); err != nil {
		return nil, err
	}

	return f.txConfig.TxEncoder()(txb.GetTx())
}





func (f Factory) Prepare(clientCtx client.Context) (Factory, error) {
	fc := f

	from := clientCtx.GetFromAddress()

	if err := fc.accountRetriever.EnsureExists(clientCtx, from); err != nil {
		return fc, err
	}

	initNum, initSeq := fc.accountNumber, fc.sequence
	if initNum == 0 || initSeq == 0 {
		num, seq, err := fc.accountRetriever.GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return fc, err
		}

		if initNum == 0 {
			fc = fc.WithAccountNumber(num)
		}

		if initSeq == 0 {
			fc = fc.WithSequence(seq)
		}
	}

	return fc, nil
}
