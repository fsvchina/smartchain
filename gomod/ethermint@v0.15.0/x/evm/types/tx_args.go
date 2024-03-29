package types

import (
	"errors"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)





type TransactionArgs struct {
	From                 *common.Address `json:"from"`
	To                   *common.Address `json:"to"`
	Gas                  *hexutil.Uint64 `json:"gas"`
	GasPrice             *hexutil.Big    `json:"gasPrice"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas"`
	Value                *hexutil.Big    `json:"value"`
	Nonce                *hexutil.Uint64 `json:"nonce"`




	Data  *hexutil.Bytes `json:"data"`
	Input *hexutil.Bytes `json:"input"`


	AccessList *ethtypes.AccessList `json:"accessList,omitempty"`
	ChainID    *hexutil.Big         `json:"chainId,omitempty"`
}


func (args *TransactionArgs) String() string {

	return fmt.Sprintf("TransactionArgs{From:%v, To:%v, Gas:%v,"+
		" Nonce:%v, Data:%v, Input:%v, AccessList:%v}",
		args.From,
		args.To,
		args.Gas,
		args.Nonce,
		args.Data,
		args.Input,
		args.AccessList)
}



func (args *TransactionArgs) ToTransaction() *MsgEthereumTx {
	var (
		chainID, value, gasPrice, maxFeePerGas, maxPriorityFeePerGas sdk.Int
		gas, nonce                                                   uint64
		from, to                                                     string
	)


	if args.ChainID != nil {
		chainID = sdk.NewIntFromBigInt(args.ChainID.ToInt())
	}

	if args.Nonce != nil {
		nonce = uint64(*args.Nonce)
	}

	if args.Gas != nil {
		gas = uint64(*args.Gas)
	}

	if args.GasPrice != nil {
		gasPrice = sdk.NewIntFromBigInt(args.GasPrice.ToInt())
	}

	if args.MaxFeePerGas != nil {
		maxFeePerGas = sdk.NewIntFromBigInt(args.MaxFeePerGas.ToInt())
	}

	if args.MaxPriorityFeePerGas != nil {
		maxPriorityFeePerGas = sdk.NewIntFromBigInt(args.MaxPriorityFeePerGas.ToInt())
	}

	if args.Value != nil {
		value = sdk.NewIntFromBigInt(args.Value.ToInt())
	}

	if args.To != nil {
		to = args.To.Hex()
	}

	var data TxData
	switch {
	case args.MaxFeePerGas != nil:
		al := AccessList{}
		if args.AccessList != nil {
			al = NewAccessList(args.AccessList)
		}

		data = &DynamicFeeTx{
			To:        to,
			ChainID:   &chainID,
			Nonce:     nonce,
			GasLimit:  gas,
			GasFeeCap: &maxFeePerGas,
			GasTipCap: &maxPriorityFeePerGas,
			Amount:    &value,
			Data:      args.GetData(),
			Accesses:  al,
		}
	case args.AccessList != nil:
		data = &AccessListTx{
			To:       to,
			ChainID:  &chainID,
			Nonce:    nonce,
			GasLimit: gas,
			GasPrice: &gasPrice,
			Amount:   &value,
			Data:     args.GetData(),
			Accesses: NewAccessList(args.AccessList),
		}
	default:
		data = &LegacyTx{
			To:       to,
			Nonce:    nonce,
			GasLimit: gas,
			GasPrice: &gasPrice,
			Amount:   &value,
			Data:     args.GetData(),
		}
	}

	any, err := PackTxData(data)
	if err != nil {
		return nil
	}

	if args.From != nil {
		from = args.From.Hex()
	}

	return &MsgEthereumTx{
		Data: any,
		From: from,
	}
}



func (args *TransactionArgs) ToMessage(globalGasCap uint64, baseFee *big.Int) (ethtypes.Message, error) {

	if args.GasPrice != nil && (args.MaxFeePerGas != nil || args.MaxPriorityFeePerGas != nil) {
		return ethtypes.Message{}, errors.New("both gasPrice and (maxFeePerGas or maxPriorityFeePerGas) specified")
	}


	addr := args.GetFrom()


	gas := globalGasCap
	if gas == 0 {
		gas = uint64(math.MaxUint64 / 2)
	}
	if args.Gas != nil {
		gas = uint64(*args.Gas)
	}
	if globalGasCap != 0 && globalGasCap < gas {
		gas = globalGasCap
	}
	var (
		gasPrice  *big.Int
		gasFeeCap *big.Int
		gasTipCap *big.Int
	)
	if baseFee == nil {

		gasPrice = new(big.Int)
		if args.GasPrice != nil {
			gasPrice = args.GasPrice.ToInt()
		}
		gasFeeCap, gasTipCap = gasPrice, gasPrice
	} else {

		if args.GasPrice != nil {

			gasPrice = args.GasPrice.ToInt()
			gasFeeCap, gasTipCap = gasPrice, gasPrice
		} else {

			gasFeeCap = new(big.Int)
			if args.MaxFeePerGas != nil {
				gasFeeCap = args.MaxFeePerGas.ToInt()
			}
			gasTipCap = new(big.Int)
			if args.MaxPriorityFeePerGas != nil {
				gasTipCap = args.MaxPriorityFeePerGas.ToInt()
			}

			gasPrice = new(big.Int)
			if gasFeeCap.BitLen() > 0 || gasTipCap.BitLen() > 0 {
				gasPrice = math.BigMin(new(big.Int).Add(gasTipCap, baseFee), gasFeeCap)
			}
		}
	}
	value := new(big.Int)
	if args.Value != nil {
		value = args.Value.ToInt()
	}
	data := args.GetData()
	var accessList ethtypes.AccessList
	if args.AccessList != nil {
		accessList = *args.AccessList
	}

	nonce := uint64(0)
	if args.Nonce != nil {
		nonce = uint64(*args.Nonce)
	}

	msg := ethtypes.NewMessage(addr, args.To, nonce, value, gas, gasPrice, gasFeeCap, gasTipCap, data, accessList, true)
	return msg, nil
}


func (args *TransactionArgs) GetFrom() common.Address {
	if args.From == nil {
		return common.Address{}
	}
	return *args.From
}


func (args *TransactionArgs) GetData() []byte {
	if args.Input != nil {
		return *args.Input
	}
	if args.Data != nil {
		return *args.Data
	}
	return nil
}
