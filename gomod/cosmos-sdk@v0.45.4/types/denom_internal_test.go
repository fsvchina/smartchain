package types

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

var (
	atom  = "atom"
	matom = "matom"
	uatom = "uatom"
	natom = "natom"
)

type internalDenomTestSuite struct {
	suite.Suite
}

func TestInternalDenomTestSuite(t *testing.T) {
	suite.Run(t, new(internalDenomTestSuite))
}

func (s *internalDenomTestSuite) TestRegisterDenom() {
	atomUnit := OneDec()

	s.Require().NoError(RegisterDenom(atom, atomUnit))
	s.Require().Error(RegisterDenom(atom, atomUnit))

	res, ok := GetDenomUnit(atom)
	s.Require().True(ok)
	s.Require().Equal(atomUnit, res)

	res, ok = GetDenomUnit(matom)
	s.Require().False(ok)
	s.Require().Equal(ZeroDec(), res)


	baseDenom = ""
	denomUnits = map[string]Dec{}
}

func (s *internalDenomTestSuite) TestConvertCoins() {
	atomUnit := OneDec()
	s.Require().NoError(RegisterDenom(atom, atomUnit))

	matomUnit := NewDecWithPrec(1, 3)
	s.Require().NoError(RegisterDenom(matom, matomUnit))

	uatomUnit := NewDecWithPrec(1, 6)
	s.Require().NoError(RegisterDenom(uatom, uatomUnit))

	natomUnit := NewDecWithPrec(1, 9)
	s.Require().NoError(RegisterDenom(natom, natomUnit))

	res, err := GetBaseDenom()
	s.Require().NoError(err)
	s.Require().Equal(res, natom)
	s.Require().Equal(NormalizeCoin(NewCoin(uatom, NewInt(1))), NewCoin(natom, NewInt(1000)))
	s.Require().Equal(NormalizeCoin(NewCoin(matom, NewInt(1))), NewCoin(natom, NewInt(1000000)))
	s.Require().Equal(NormalizeCoin(NewCoin(atom, NewInt(1))), NewCoin(natom, NewInt(1000000000)))

	coins, err := ParseCoinsNormalized("1atom,1matom,1uatom")
	s.Require().NoError(err)
	s.Require().Equal(coins, Coins{
		Coin{natom, NewInt(1000000000)},
		Coin{natom, NewInt(1000000)},
		Coin{natom, NewInt(1000)},
	})

	testCases := []struct {
		input  Coin
		denom  string
		result Coin
		expErr bool
	}{
		{NewCoin("foo", ZeroInt()), atom, Coin{}, true},
		{NewCoin(atom, ZeroInt()), "foo", Coin{}, true},
		{NewCoin(atom, ZeroInt()), "FOO", Coin{}, true},

		{NewCoin(atom, NewInt(5)), matom, NewCoin(matom, NewInt(5000)), false},
		{NewCoin(atom, NewInt(5)), uatom, NewCoin(uatom, NewInt(5000000)), false},
		{NewCoin(atom, NewInt(5)), natom, NewCoin(natom, NewInt(5000000000)), false},

		{NewCoin(uatom, NewInt(5000000)), matom, NewCoin(matom, NewInt(5000)), false},
		{NewCoin(uatom, NewInt(5000000)), natom, NewCoin(natom, NewInt(5000000000)), false},
		{NewCoin(uatom, NewInt(5000000)), atom, NewCoin(atom, NewInt(5)), false},

		{NewCoin(matom, NewInt(5000)), natom, NewCoin(natom, NewInt(5000000000)), false},
		{NewCoin(matom, NewInt(5000)), uatom, NewCoin(uatom, NewInt(5000000)), false},
	}

	for i, tc := range testCases {
		res, err := ConvertCoin(tc.input, tc.denom)
		s.Require().Equal(
			tc.expErr, err != nil,
			"unexpected error; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
		s.Require().Equal(
			tc.result, res,
			"invalid result; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
	}


	baseDenom = ""
	denomUnits = map[string]Dec{}
}

func (s *internalDenomTestSuite) TestConvertDecCoins() {
	atomUnit := OneDec()
	s.Require().NoError(RegisterDenom(atom, atomUnit))

	matomUnit := NewDecWithPrec(1, 3)
	s.Require().NoError(RegisterDenom(matom, matomUnit))

	uatomUnit := NewDecWithPrec(1, 6)
	s.Require().NoError(RegisterDenom(uatom, uatomUnit))

	natomUnit := NewDecWithPrec(1, 9)
	s.Require().NoError(RegisterDenom(natom, natomUnit))

	res, err := GetBaseDenom()
	s.Require().NoError(err)
	s.Require().Equal(res, natom)
	s.Require().Equal(NormalizeDecCoin(NewDecCoin(uatom, NewInt(1))), NewDecCoin(natom, NewInt(1000)))
	s.Require().Equal(NormalizeDecCoin(NewDecCoin(matom, NewInt(1))), NewDecCoin(natom, NewInt(1000000)))
	s.Require().Equal(NormalizeDecCoin(NewDecCoin(atom, NewInt(1))), NewDecCoin(natom, NewInt(1000000000)))

	coins, err := ParseCoinsNormalized("0.1atom,0.1matom,0.1uatom")
	s.Require().NoError(err)
	s.Require().Equal(coins, Coins{
		Coin{natom, NewInt(100000000)},
		Coin{natom, NewInt(100000)},
		Coin{natom, NewInt(100)},
	})

	testCases := []struct {
		input  DecCoin
		denom  string
		result DecCoin
		expErr bool
	}{
		{NewDecCoin("foo", ZeroInt()), atom, DecCoin{}, true},
		{NewDecCoin(atom, ZeroInt()), "foo", DecCoin{}, true},
		{NewDecCoin(atom, ZeroInt()), "FOO", DecCoin{}, true},


		{NewDecCoinFromDec(atom, NewDecWithPrec(5, 1)), matom, NewDecCoin(matom, NewInt(500)), false},
		{NewDecCoinFromDec(atom, NewDecWithPrec(5, 1)), uatom, NewDecCoin(uatom, NewInt(500000)), false},
		{NewDecCoinFromDec(atom, NewDecWithPrec(5, 1)), natom, NewDecCoin(natom, NewInt(500000000)), false},

		{NewDecCoin(uatom, NewInt(5000000)), matom, NewDecCoin(matom, NewInt(5000)), false},
		{NewDecCoin(uatom, NewInt(5000000)), natom, NewDecCoin(natom, NewInt(5000000000)), false},
		{NewDecCoin(uatom, NewInt(5000000)), atom, NewDecCoin(atom, NewInt(5)), false},

		{NewDecCoin(matom, NewInt(5000)), natom, NewDecCoin(natom, NewInt(5000000000)), false},
		{NewDecCoin(matom, NewInt(5000)), uatom, NewDecCoin(uatom, NewInt(5000000)), false},
	}

	for i, tc := range testCases {
		res, err := ConvertDecCoin(tc.input, tc.denom)
		s.Require().Equal(
			tc.expErr, err != nil,
			"unexpected error; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
		s.Require().Equal(
			tc.result, res,
			"invalid result; tc: #%d, input: %s, denom: %s", i+1, tc.input, tc.denom,
		)
	}


	baseDenom = ""
	denomUnits = map[string]Dec{}
}

func (s *internalDenomTestSuite) TestDecOperationOrder() {
	dec, err := NewDecFromStr("11")
	s.Require().NoError(err)
	s.Require().NoError(RegisterDenom("unit1", dec))
	dec, err = NewDecFromStr("100000011")
	s.Require().NoError(RegisterDenom("unit2", dec))

	coin, err := ConvertCoin(NewCoin("unit1", NewInt(100000011)), "unit2")
	s.Require().NoError(err)
	s.Require().Equal(coin, NewCoin("unit2", NewInt(11)))


	baseDenom = ""
	denomUnits = map[string]Dec{}
}
