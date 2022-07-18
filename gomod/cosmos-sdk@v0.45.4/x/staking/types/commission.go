package types

import (
	"time"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


func NewCommissionRates(rate, maxRate, maxChangeRate sdk.Dec) CommissionRates {
	return CommissionRates{
		Rate:          rate,
		MaxRate:       maxRate,
		MaxChangeRate: maxChangeRate,
	}
}


func NewCommission(rate, maxRate, maxChangeRate sdk.Dec) Commission {
	return Commission{
		CommissionRates: NewCommissionRates(rate, maxRate, maxChangeRate),
		UpdateTime:      time.Unix(0, 0).UTC(),
	}
}



func NewCommissionWithTime(rate, maxRate, maxChangeRate sdk.Dec, updatedAt time.Time) Commission {
	return Commission{
		CommissionRates: NewCommissionRates(rate, maxRate, maxChangeRate),
		UpdateTime:      updatedAt,
	}
}


func (c Commission) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}


func (cr CommissionRates) String() string {
	out, _ := yaml.Marshal(cr)
	return string(out)
}



func (cr CommissionRates) Validate() error {
	switch {
	case cr.MaxRate.IsNegative():

		return ErrCommissionNegative

	case cr.MaxRate.GT(sdk.OneDec()):

		return ErrCommissionHuge

	case cr.Rate.IsNegative():

		return ErrCommissionNegative

	case cr.Rate.GT(cr.MaxRate):

		return ErrCommissionGTMaxRate

	case cr.MaxChangeRate.IsNegative():

		return ErrCommissionChangeRateNegative

	case cr.MaxChangeRate.GT(cr.MaxRate):

		return ErrCommissionChangeRateGTMaxRate
	}

	return nil
}



func (c Commission) ValidateNewRate(newRate sdk.Dec, blockTime time.Time) error {
	switch {
	case blockTime.Sub(c.UpdateTime).Hours() < 24:

		return ErrCommissionUpdateTime

	case newRate.IsNegative():

		return ErrCommissionNegative

	case newRate.GT(c.MaxRate):

		return ErrCommissionGTMaxRate

	case newRate.Sub(c.Rate).GT(c.MaxChangeRate):

		return ErrCommissionGTMaxChangeRate
	}

	return nil
}
