package simulation

import (
	"errors"
	"math/big"
	"math/rand"
	"time"
	"unsafe"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    
	letterIdxMask = 1<<letterIdxBits - 1 
	letterIdxMax  = 63 / letterIdxBits   
)





func RandStringOfLength(r *rand.Rand, n int) string {
	b := make([]byte, n)
	
	for i, cache, remain := n-1, r.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}


func RandPositiveInt(r *rand.Rand, max sdk.Int) (sdk.Int, error) {
	if !max.GTE(sdk.OneInt()) {
		return sdk.Int{}, errors.New("max too small")
	}

	max = max.Sub(sdk.OneInt())

	return sdk.NewIntFromBigInt(new(big.Int).Rand(r, max.BigInt())).Add(sdk.OneInt()), nil
}



func RandomAmount(r *rand.Rand, max sdk.Int) sdk.Int {
	var randInt = big.NewInt(0)

	switch r.Intn(10) {
	case 0:
		
	case 1:
		randInt = max.BigInt()
	default: 
		randInt = big.NewInt(0).Rand(r, max.BigInt()) 
	}

	return sdk.NewIntFromBigInt(randInt)
}



func RandomDecAmount(r *rand.Rand, max sdk.Dec) sdk.Dec {
	var randInt = big.NewInt(0)

	switch r.Intn(10) {
	case 0:
		
	case 1:
		randInt = max.BigInt() 
	default: 
		randInt = big.NewInt(0).Rand(r, max.BigInt())
	}

	return sdk.NewDecFromBigIntWithPrec(randInt, sdk.Precision)
}


func RandTimestamp(r *rand.Rand) time.Time {
	
	unixTime := r.Int63n(253373529600)
	return time.Unix(unixTime, 0)
}


func RandIntBetween(r *rand.Rand, min, max int) int {
	return r.Intn(max-min) + min
}




func RandSubsetCoins(r *rand.Rand, coins sdk.Coins) sdk.Coins {
	if len(coins) == 0 {
		return sdk.Coins{}
	}
	
	denomIdx := r.Intn(len(coins))
	coin := coins[denomIdx]
	amt, err := RandPositiveInt(r, coin.Amount)
	
	if err != nil {
		return sdk.Coins{}
	}

	subset := sdk.Coins{sdk.NewCoin(coin.Denom, amt)}

	for i, c := range coins {
		
		if i == denomIdx {
			continue
		}
		
		
		if r.Intn(2) == 0 && len(coins) != 1 {
			continue
		}

		amt, err := RandPositiveInt(r, c.Amount)
		
		if err != nil {
			continue
		}

		subset = append(subset, sdk.NewCoin(c.Denom, amt))
	}

	return subset.Sort()
}




//

func DeriveRand(r *rand.Rand) *rand.Rand {
	const num = 8 
	ms := multiSource(make([]rand.Source, num))

	for i := 0; i < num; i++ {
		ms[i] = rand.NewSource(r.Int63())
	}

	return rand.New(ms)
}

type multiSource []rand.Source

func (ms multiSource) Int63() (r int64) {
	for _, source := range ms {
		r ^= source.Int63()
	}

	return r
}

func (ms multiSource) Seed(seed int64) {
	panic("multiSource Seed should not be called")
}
