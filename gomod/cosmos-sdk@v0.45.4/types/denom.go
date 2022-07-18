package types

import (
	"fmt"
)



var denomUnits = map[string]Dec{}


var baseDenom string



func RegisterDenom(denom string, unit Dec) error {
	if err := ValidateDenom(denom); err != nil {
		return err
	}

	if _, ok := denomUnits[denom]; ok {
		return fmt.Errorf("denom %s already registered", denom)
	}

	denomUnits[denom] = unit

	if baseDenom == "" || unit.LT(denomUnits[baseDenom]) {
		baseDenom = denom
	}
	return nil
}



func GetDenomUnit(denom string) (Dec, bool) {
	if err := ValidateDenom(denom); err != nil {
		return ZeroDec(), false
	}

	unit, ok := denomUnits[denom]
	if !ok {
		return ZeroDec(), false
	}

	return unit, true
}


func GetBaseDenom() (string, error) {
	if baseDenom == "" {
		return "", fmt.Errorf("no denom is registered")
	}
	return baseDenom, nil
}




func ConvertCoin(coin Coin, denom string) (Coin, error) {
	if err := ValidateDenom(denom); err != nil {
		return Coin{}, err
	}

	srcUnit, ok := GetDenomUnit(coin.Denom)
	if !ok {
		return Coin{}, fmt.Errorf("source denom not registered: %s", coin.Denom)
	}

	dstUnit, ok := GetDenomUnit(denom)
	if !ok {
		return Coin{}, fmt.Errorf("destination denom not registered: %s", denom)
	}

	if srcUnit.Equal(dstUnit) {
		return NewCoin(denom, coin.Amount), nil
	}

	return NewCoin(denom, coin.Amount.ToDec().Mul(srcUnit).Quo(dstUnit).TruncateInt()), nil
}




func ConvertDecCoin(coin DecCoin, denom string) (DecCoin, error) {
	if err := ValidateDenom(denom); err != nil {
		return DecCoin{}, err
	}

	srcUnit, ok := GetDenomUnit(coin.Denom)
	if !ok {
		return DecCoin{}, fmt.Errorf("source denom not registered: %s", coin.Denom)
	}

	dstUnit, ok := GetDenomUnit(denom)
	if !ok {
		return DecCoin{}, fmt.Errorf("destination denom not registered: %s", denom)
	}

	if srcUnit.Equal(dstUnit) {
		return NewDecCoinFromDec(denom, coin.Amount), nil
	}

	return NewDecCoinFromDec(denom, coin.Amount.Mul(srcUnit).Quo(dstUnit)), nil
}



func NormalizeCoin(coin Coin) Coin {
	base, err := GetBaseDenom()
	if err != nil {
		return coin
	}
	newCoin, err := ConvertCoin(coin, base)
	if err != nil {
		return coin
	}
	return newCoin
}



func NormalizeDecCoin(coin DecCoin) DecCoin {
	base, err := GetBaseDenom()
	if err != nil {
		return coin
	}
	newCoin, err := ConvertDecCoin(coin, base)
	if err != nil {
		return coin
	}
	return newCoin
}


func NormalizeCoins(coins []DecCoin) Coins {
	if coins == nil {
		return nil
	}
	result := make([]Coin, 0, len(coins))

	for _, coin := range coins {
		newCoin, _ := NormalizeDecCoin(coin).TruncateDecimal()
		result = append(result, newCoin)
	}

	return result
}
