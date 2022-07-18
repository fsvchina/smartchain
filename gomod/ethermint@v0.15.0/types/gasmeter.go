package types

import (
	fmt "fmt"
	math "math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)




type ErrorNegativeGasConsumed struct {
	Descriptor string
}



type ErrorGasOverflow struct {
	Descriptor string
}

type infiniteGasMeterWithLimit struct {
	consumed sdk.Gas
	limit    sdk.Gas
}


func NewInfiniteGasMeterWithLimit(limit sdk.Gas) sdk.GasMeter {
	return &infiniteGasMeterWithLimit{
		consumed: 0,
		limit:    limit,
	}
}

func (g *infiniteGasMeterWithLimit) GasConsumed() sdk.Gas {
	return g.consumed
}

func (g *infiniteGasMeterWithLimit) GasConsumedToLimit() sdk.Gas {
	return g.consumed
}

func (g *infiniteGasMeterWithLimit) Limit() sdk.Gas {
	return g.limit
}



func addUint64Overflow(a, b uint64) (uint64, bool) {
	if math.MaxUint64-a < b {
		return 0, true
	}

	return a + b, false
}

func (g *infiniteGasMeterWithLimit) ConsumeGas(amount sdk.Gas, descriptor string) {
	var overflow bool

	g.consumed, overflow = addUint64Overflow(g.consumed, amount)
	if overflow {
		panic(ErrorGasOverflow{descriptor})
	}
}



//



func (g *infiniteGasMeterWithLimit) RefundGas(amount sdk.Gas, descriptor string) {
	if g.consumed < amount {
		panic(ErrorNegativeGasConsumed{Descriptor: descriptor})
	}

	g.consumed -= amount
}

func (g *infiniteGasMeterWithLimit) IsPastLimit() bool {
	return false
}

func (g *infiniteGasMeterWithLimit) IsOutOfGas() bool {
	return false
}

func (g *infiniteGasMeterWithLimit) String() string {
	return fmt.Sprintf("InfiniteGasMeter:\n  consumed: %d", g.consumed)
}
