package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"testing"
)

var _ CustomProtobufType = (*Dec)(nil)



type Dec struct {
	i *big.Int
}

const (

	Precision = 18



	DecimalPrecisionBits = 60

	maxDecBitLen = maxBitLen + DecimalPrecisionBits


	maxApproxRootIterations = 100
)

var (
	precisionReuse       = new(big.Int).Exp(big.NewInt(10), big.NewInt(Precision), nil)
	fivePrecision        = new(big.Int).Quo(precisionReuse, big.NewInt(2))
	precisionMultipliers []*big.Int
	zeroInt              = big.NewInt(0)
	oneInt               = big.NewInt(1)
	tenInt               = big.NewInt(10)
)


var (
	ErrEmptyDecimalStr      = errors.New("decimal string cannot be empty")
	ErrInvalidDecimalLength = errors.New("invalid decimal length")
	ErrInvalidDecimalStr    = errors.New("invalid decimal string")
)


func init() {
	precisionMultipliers = make([]*big.Int, Precision+1)
	for i := 0; i <= Precision; i++ {
		precisionMultipliers[i] = calcPrecisionMultiplier(int64(i))
	}
}

func precisionInt() *big.Int {
	return new(big.Int).Set(precisionReuse)
}

func ZeroDec() Dec     { return Dec{new(big.Int).Set(zeroInt)} }
func OneDec() Dec      { return Dec{precisionInt()} }
func SmallestDec() Dec { return Dec{new(big.Int).Set(oneInt)} }


func calcPrecisionMultiplier(prec int64) *big.Int {
	if prec > Precision {
		panic(fmt.Sprintf("too much precision, maximum %v, provided %v", Precision, prec))
	}
	zerosToAdd := Precision - prec
	multiplier := new(big.Int).Exp(tenInt, big.NewInt(zerosToAdd), nil)
	return multiplier
}


func precisionMultiplier(prec int64) *big.Int {
	if prec > Precision {
		panic(fmt.Sprintf("too much precision, maximum %v, provided %v", Precision, prec))
	}
	return precisionMultipliers[prec]
}


func NewDec(i int64) Dec {
	return NewDecWithPrec(i, 0)
}



func NewDecWithPrec(i, prec int64) Dec {
	return Dec{
		new(big.Int).Mul(big.NewInt(i), precisionMultiplier(prec)),
	}
}



func NewDecFromBigInt(i *big.Int) Dec {
	return NewDecFromBigIntWithPrec(i, 0)
}



func NewDecFromBigIntWithPrec(i *big.Int, prec int64) Dec {
	return Dec{
		new(big.Int).Mul(i, precisionMultiplier(prec)),
	}
}



func NewDecFromInt(i Int) Dec {
	return NewDecFromIntWithPrec(i, 0)
}



func NewDecFromIntWithPrec(i Int, prec int64) Dec {
	return Dec{
		new(big.Int).Mul(i.BigInt(), precisionMultiplier(prec)),
	}
}









//


//

func NewDecFromStr(str string) (Dec, error) {
	if len(str) == 0 {
		return Dec{}, ErrEmptyDecimalStr
	}


	neg := false
	if str[0] == '-' {
		neg = true
		str = str[1:]
	}

	if len(str) == 0 {
		return Dec{}, ErrEmptyDecimalStr
	}

	strs := strings.Split(str, ".")
	lenDecs := 0
	combinedStr := strs[0]

	if len(strs) == 2 {
		lenDecs = len(strs[1])
		if lenDecs == 0 || len(combinedStr) == 0 {
			return Dec{}, ErrInvalidDecimalLength
		}
		combinedStr += strs[1]
	} else if len(strs) > 2 {
		return Dec{}, ErrInvalidDecimalStr
	}

	if lenDecs > Precision {
		return Dec{}, fmt.Errorf("invalid precision; max: %d, got: %d", Precision, lenDecs)
	}


	zerosToAdd := Precision - lenDecs
	zeros := fmt.Sprintf(`%0`+strconv.Itoa(zerosToAdd)+`s`, "")
	combinedStr += zeros

	combined, ok := new(big.Int).SetString(combinedStr, 10)
	if !ok {
		return Dec{}, fmt.Errorf("failed to set decimal string: %s", combinedStr)
	}
	if combined.BitLen() > maxDecBitLen {
		return Dec{}, fmt.Errorf("decimal out of range; bitLen: got %d, max %d", combined.BitLen(), maxDecBitLen)
	}
	if neg {
		combined = new(big.Int).Neg(combined)
	}

	return Dec{combined}, nil
}


func MustNewDecFromStr(s string) Dec {
	dec, err := NewDecFromStr(s)
	if err != nil {
		panic(err)
	}
	return dec
}

func (d Dec) IsNil() bool       { return d.i == nil }
func (d Dec) IsZero() bool      { return (d.i).Sign() == 0 }
func (d Dec) IsNegative() bool  { return (d.i).Sign() == -1 }
func (d Dec) IsPositive() bool  { return (d.i).Sign() == 1 }
func (d Dec) Equal(d2 Dec) bool { return (d.i).Cmp(d2.i) == 0 }
func (d Dec) GT(d2 Dec) bool    { return (d.i).Cmp(d2.i) > 0 }
func (d Dec) GTE(d2 Dec) bool   { return (d.i).Cmp(d2.i) >= 0 }
func (d Dec) LT(d2 Dec) bool    { return (d.i).Cmp(d2.i) < 0 }
func (d Dec) LTE(d2 Dec) bool   { return (d.i).Cmp(d2.i) <= 0 }
func (d Dec) Neg() Dec          { return Dec{new(big.Int).Neg(d.i)} }
func (d Dec) Abs() Dec          { return Dec{new(big.Int).Abs(d.i)} }


func (d Dec) BigInt() *big.Int {
	if d.IsNil() {
		return nil
	}

	cp := new(big.Int)
	return cp.Set(d.i)
}


func (d Dec) Add(d2 Dec) Dec {
	res := new(big.Int).Add(d.i, d2.i)

	if res.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{res}
}


func (d Dec) Sub(d2 Dec) Dec {
	res := new(big.Int).Sub(d.i, d2.i)

	if res.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{res}
}


func (d Dec) Mul(d2 Dec) Dec {
	mul := new(big.Int).Mul(d.i, d2.i)
	chopped := chopPrecisionAndRound(mul)

	if chopped.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{chopped}
}


func (d Dec) MulTruncate(d2 Dec) Dec {
	mul := new(big.Int).Mul(d.i, d2.i)
	chopped := chopPrecisionAndTruncate(mul)

	if chopped.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{chopped}
}


func (d Dec) MulInt(i Int) Dec {
	mul := new(big.Int).Mul(d.i, i.i)

	if mul.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{mul}
}


func (d Dec) MulInt64(i int64) Dec {
	mul := new(big.Int).Mul(d.i, big.NewInt(i))

	if mul.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{mul}
}


func (d Dec) Quo(d2 Dec) Dec {

	mul := new(big.Int).Mul(d.i, precisionReuse)
	mul.Mul(mul, precisionReuse)

	quo := new(big.Int).Quo(mul, d2.i)
	chopped := chopPrecisionAndRound(quo)

	if chopped.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{chopped}
}


func (d Dec) QuoTruncate(d2 Dec) Dec {

	mul := new(big.Int).Mul(d.i, precisionReuse)
	mul.Mul(mul, precisionReuse)

	quo := mul.Quo(mul, d2.i)
	chopped := chopPrecisionAndTruncate(quo)

	if chopped.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{chopped}
}


func (d Dec) QuoRoundUp(d2 Dec) Dec {

	mul := new(big.Int).Mul(d.i, precisionReuse)
	mul.Mul(mul, precisionReuse)

	quo := new(big.Int).Quo(mul, d2.i)
	chopped := chopPrecisionAndRoundUp(quo)

	if chopped.BitLen() > maxDecBitLen {
		panic("Int overflow")
	}
	return Dec{chopped}
}


func (d Dec) QuoInt(i Int) Dec {
	mul := new(big.Int).Quo(d.i, i.i)
	return Dec{mul}
}


func (d Dec) QuoInt64(i int64) Dec {
	mul := new(big.Int).Quo(d.i, big.NewInt(i))
	return Dec{mul}
}







func (d Dec) ApproxRoot(root uint64) (guess Dec, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = errors.New("out of bounds")
			}
		}
	}()

	if d.IsNegative() {
		absRoot, err := d.MulInt64(-1).ApproxRoot(root)
		return absRoot.MulInt64(-1), err
	}

	if root == 1 || d.IsZero() || d.Equal(OneDec()) {
		return d, nil
	}

	if root == 0 {
		return OneDec(), nil
	}

	rootInt := NewIntFromUint64(root)
	guess, delta := OneDec(), OneDec()

	for iter := 0; delta.Abs().GT(SmallestDec()) && iter < maxApproxRootIterations; iter++ {
		prev := guess.Power(root - 1)
		if prev.IsZero() {
			prev = SmallestDec()
		}
		delta = d.Quo(prev)
		delta = delta.Sub(guess)
		delta = delta.QuoInt(rootInt)

		guess = guess.Add(delta)
	}

	return guess, nil
}


func (d Dec) Power(power uint64) Dec {
	if power == 0 {
		return OneDec()
	}
	tmp := OneDec()

	for i := power; i > 1; {
		if i%2 != 0 {
			tmp = tmp.Mul(d)
		}
		i /= 2
		d = d.Mul(d)
	}

	return d.Mul(tmp)
}



func (d Dec) ApproxSqrt() (Dec, error) {
	return d.ApproxRoot(2)
}


func (d Dec) IsInteger() bool {
	return new(big.Int).Rem(d.i, precisionReuse).Sign() == 0
}


func (d Dec) Format(s fmt.State, verb rune) {
	_, err := s.Write([]byte(d.String()))
	if err != nil {
		panic(err)
	}
}

func (d Dec) String() string {
	if d.i == nil {
		return d.i.String()
	}

	isNeg := d.IsNegative()

	if isNeg {
		d = d.Neg()
	}

	bzInt, err := d.i.MarshalText()
	if err != nil {
		return ""
	}
	inputSize := len(bzInt)

	var bzStr []byte



	if inputSize <= Precision {
		bzStr = make([]byte, Precision+2)


		bzStr[0] = byte('0')
		bzStr[1] = byte('.')


		for i := 0; i < Precision-inputSize; i++ {
			bzStr[i+2] = byte('0')
		}


		copy(bzStr[2+(Precision-inputSize):], bzInt)
	} else {

		bzStr = make([]byte, inputSize+1)
		decPointPlace := inputSize - Precision

		copy(bzStr, bzInt[:decPointPlace])
		bzStr[decPointPlace] = byte('.')
		copy(bzStr[decPointPlace+1:], bzInt[decPointPlace:])
	}

	if isNeg {
		return "-" + string(bzStr)
	}

	return string(bzStr)
}



func (d Dec) Float64() (float64, error) {
	return strconv.ParseFloat(d.String(), 64)
}



func (d Dec) MustFloat64() float64 {
	if value, err := strconv.ParseFloat(d.String(), 64); err != nil {
		panic(err)
	} else {
		return value
	}
}












//

func chopPrecisionAndRound(d *big.Int) *big.Int {

	if d.Sign() == -1 {

		d = d.Neg(d)
		d = chopPrecisionAndRound(d)
		d = d.Neg(d)
		return d
	}


	quo, rem := d, big.NewInt(0)
	quo, rem = quo.QuoRem(d, precisionReuse, rem)

	if rem.Sign() == 0 {
		return quo
	}

	switch rem.Cmp(fivePrecision) {
	case -1:
		return quo
	case 1:
		return quo.Add(quo, oneInt)
	default:

		if quo.Bit(0) == 0 {
			return quo
		}
		return quo.Add(quo, oneInt)
	}
}

func chopPrecisionAndRoundUp(d *big.Int) *big.Int {

	if d.Sign() == -1 {

		d = d.Neg(d)

		d = chopPrecisionAndTruncate(d)
		d = d.Neg(d)
		return d
	}


	quo, rem := d, big.NewInt(0)
	quo, rem = quo.QuoRem(d, precisionReuse, rem)

	if rem.Sign() == 0 {
		return quo
	}

	return quo.Add(quo, oneInt)
}

func chopPrecisionAndRoundNonMutative(d *big.Int) *big.Int {
	tmp := new(big.Int).Set(d)
	return chopPrecisionAndRound(tmp)
}


func (d Dec) RoundInt64() int64 {
	chopped := chopPrecisionAndRoundNonMutative(d.i)
	if !chopped.IsInt64() {
		panic("Int64() out of bound")
	}
	return chopped.Int64()
}


func (d Dec) RoundInt() Int {
	return NewIntFromBigInt(chopPrecisionAndRoundNonMutative(d.i))
}



func chopPrecisionAndTruncate(d *big.Int) *big.Int {
	return new(big.Int).Quo(d, precisionReuse)
}


func (d Dec) TruncateInt64() int64 {
	chopped := chopPrecisionAndTruncate(d.i)
	if !chopped.IsInt64() {
		panic("Int64() out of bound")
	}
	return chopped.Int64()
}


func (d Dec) TruncateInt() Int {
	return NewIntFromBigInt(chopPrecisionAndTruncate(d.i))
}


func (d Dec) TruncateDec() Dec {
	return NewDecFromBigInt(chopPrecisionAndTruncate(d.i))
}



func (d Dec) Ceil() Dec {
	tmp := new(big.Int).Set(d.i)

	quo, rem := tmp, big.NewInt(0)
	quo, rem = quo.QuoRem(tmp, precisionReuse, rem)


	if rem.Cmp(zeroInt) == 0 {
		return NewDecFromBigInt(quo)
	}

	if rem.Sign() == -1 {
		return NewDecFromBigInt(quo)
	}

	return NewDecFromBigInt(quo.Add(quo, oneInt))
}



var MaxSortableDec = OneDec().Quo(SmallestDec())




func ValidSortableDec(dec Dec) bool {
	return dec.Abs().LTE(MaxSortableDec)
}




func SortableDecBytes(dec Dec) []byte {
	if !ValidSortableDec(dec) {
		panic("dec must be within bounds")
	}


	if dec.Equal(MaxSortableDec) {
		return []byte("max")
	}

	if dec.Equal(MaxSortableDec.Neg()) {
		return []byte("--")
	}

	if dec.IsNegative() {
		return append([]byte("-"), []byte(fmt.Sprintf(fmt.Sprintf("%%0%ds", Precision*2+1), dec.Abs().String()))...)
	}
	return []byte(fmt.Sprintf(fmt.Sprintf("%%0%ds", Precision*2+1), dec.String()))
}


var nilJSON []byte

func init() {
	empty := new(big.Int)
	bz, _ := empty.MarshalText()
	nilJSON, _ = json.Marshal(string(bz))
}


func (d Dec) MarshalJSON() ([]byte, error) {
	if d.i == nil {
		return nilJSON, nil
	}
	return json.Marshal(d.String())
}


func (d *Dec) UnmarshalJSON(bz []byte) error {
	if d.i == nil {
		d.i = new(big.Int)
	}

	var text string
	err := json.Unmarshal(bz, &text)
	if err != nil {
		return err
	}


	newDec, err := NewDecFromStr(text)
	if err != nil {
		return err
	}

	d.i = newDec.i
	return nil
}


func (d Dec) MarshalYAML() (interface{}, error) {
	return d.String(), nil
}


func (d Dec) Marshal() ([]byte, error) {
	if d.i == nil {
		d.i = new(big.Int)
	}
	return d.i.MarshalText()
}


func (d *Dec) MarshalTo(data []byte) (n int, err error) {
	if d.i == nil {
		d.i = new(big.Int)
	}

	if d.i.Cmp(zeroInt) == 0 {
		copy(data, []byte{0x30})
		return 1, nil
	}

	bz, err := d.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}


func (d *Dec) Unmarshal(data []byte) error {
	if len(data) == 0 {
		d = nil
		return nil
	}

	if d.i == nil {
		d.i = new(big.Int)
	}

	if err := d.i.UnmarshalText(data); err != nil {
		return err
	}

	if d.i.BitLen() > maxDecBitLen {
		return fmt.Errorf("decimal out of range; got: %d, max: %d", d.i.BitLen(), maxDecBitLen)
	}

	return nil
}


func (d *Dec) Size() int {
	bz, _ := d.Marshal()
	return len(bz)
}


func (d Dec) MarshalAmino() ([]byte, error)   { return d.Marshal() }
func (d *Dec) UnmarshalAmino(bz []byte) error { return d.Unmarshal(bz) }

func (dp DecProto) String() string {
	return dp.Dec.String()
}




func DecsEqual(d1s, d2s []Dec) bool {
	if len(d1s) != len(d2s) {
		return false
	}

	for i, d1 := range d1s {
		if !d1.Equal(d2s[i]) {
			return false
		}
	}
	return true
}


func MinDec(d1, d2 Dec) Dec {
	if d1.LT(d2) {
		return d1
	}
	return d2
}


func MaxDec(d1, d2 Dec) Dec {
	if d1.LT(d2) {
		return d2
	}
	return d1
}


func DecEq(t *testing.T, exp, got Dec) (*testing.T, bool, string, string, string) {
	return t, exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}
