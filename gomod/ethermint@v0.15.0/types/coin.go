package types

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (

	//





	AttoPhoton string = "afsv"



	BaseDenomUnit = 18


	DefaultGasPrice = 20
)


var PowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))



func NewPhotonCoin(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(AttoPhoton, amount)
}



func NewPhotonDecCoin(amount sdk.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoPhoton, amount)
}



func NewPhotonCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoPhoton, amount)
}
