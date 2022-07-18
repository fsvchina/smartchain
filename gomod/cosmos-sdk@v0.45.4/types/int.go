package types

import (
	"encoding"
	"encoding/json"
	"fmt"
	"testing"

	"math/big"
)

const maxBitLen = 256

func newIntegerFromString(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 0)
}

func equal(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == 0 }

func gt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == 1 }

func gte(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) >= 0 }

func lt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == -1 }

func lte(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) <= 0 }

func add(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Add(i, i2) }

func sub(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Sub(i, i2) }

func mul(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mul(i, i2) }

func div(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Quo(i, i2) }

func mod(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mod(i, i2) }

func neg(i *big.Int) *big.Int { return new(big.Int).Neg(i) }

func abs(i *big.Int) *big.Int { return new(big.Int).Abs(i) }

func min(i *big.Int, i2 *big.Int) *big.Int {
	if i.Cmp(i2) == 1 {
		return new(big.Int).Set(i2)
	}

	return new(big.Int).Set(i)
}

func max(i *big.Int, i2 *big.Int) *big.Int {
	if i.Cmp(i2) == -1 {
		return new(big.Int).Set(i2)
	}

	return new(big.Int).Set(i)
}

func unmarshalText(i *big.Int, text string) error {
	if err := i.UnmarshalText([]byte(text)); err != nil {
		return err
	}

	if i.BitLen() > maxBitLen {
		return fmt.Errorf("integer out of range: %s", text)
	}

	return nil
}

var _ CustomProtobufType = (*Int)(nil)




type Int struct {
	i *big.Int
}


func (i Int) BigInt() *big.Int {
	if i.IsNil() {
		return nil
	}
	return new(big.Int).Set(i.i)
}


func (i Int) IsNil() bool {
	return i.i == nil
}


func NewInt(n int64) Int {
	return Int{big.NewInt(n)}
}


func NewIntFromUint64(n uint64) Int {
	b := big.NewInt(0)
	b.SetUint64(n)
	return Int{b}
}



func NewIntFromBigInt(i *big.Int) Int {
	if i == nil {
		return Int{}
	}

	if i.BitLen() > maxBitLen {
		panic("NewIntFromBigInt() out of bound")
	}
	return Int{i}
}


func NewIntFromString(s string) (res Int, ok bool) {
	i, ok := newIntegerFromString(s)
	if !ok {
		return
	}

	if i.BitLen() > maxBitLen {
		ok = false
		return
	}
	return Int{i}, true
}



func NewIntWithDecimal(n int64, dec int) Int {
	if dec < 0 {
		panic("NewIntWithDecimal() decimal is negative")
	}
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(dec)), nil)
	i := new(big.Int)
	i.Mul(big.NewInt(n), exp)


	if i.BitLen() > maxBitLen {
		panic("NewIntWithDecimal() out of bound")
	}
	return Int{i}
}


func ZeroInt() Int { return Int{big.NewInt(0)} }


func OneInt() Int { return Int{big.NewInt(1)} }


func (i Int) ToDec() Dec {
	return NewDecFromInt(i)
}



func (i Int) Int64() int64 {
	if !i.i.IsInt64() {
		panic("Int64() out of bound")
	}
	return i.i.Int64()
}


func (i Int) IsInt64() bool {
	return i.i.IsInt64()
}



func (i Int) Uint64() uint64 {
	if !i.i.IsUint64() {
		panic("Uint64() out of bounds")
	}
	return i.i.Uint64()
}


func (i Int) IsUint64() bool {
	return i.i.IsUint64()
}


func (i Int) IsZero() bool {
	return i.i.Sign() == 0
}


func (i Int) IsNegative() bool {
	return i.i.Sign() == -1
}


func (i Int) IsPositive() bool {
	return i.i.Sign() == 1
}


func (i Int) Sign() int {
	return i.i.Sign()
}


func (i Int) Equal(i2 Int) bool {
	return equal(i.i, i2.i)
}


func (i Int) GT(i2 Int) bool {
	return gt(i.i, i2.i)
}



func (i Int) GTE(i2 Int) bool {
	return gte(i.i, i2.i)
}


func (i Int) LT(i2 Int) bool {
	return lt(i.i, i2.i)
}


func (i Int) LTE(i2 Int) bool {
	return lte(i.i, i2.i)
}


func (i Int) Add(i2 Int) (res Int) {
	res = Int{add(i.i, i2.i)}

	if res.i.BitLen() > maxBitLen {
		panic("Int overflow")
	}
	return
}


func (i Int) AddRaw(i2 int64) Int {
	return i.Add(NewInt(i2))
}


func (i Int) Sub(i2 Int) (res Int) {
	res = Int{sub(i.i, i2.i)}

	if res.i.BitLen() > maxBitLen {
		panic("Int overflow")
	}
	return
}


func (i Int) SubRaw(i2 int64) Int {
	return i.Sub(NewInt(i2))
}


func (i Int) Mul(i2 Int) (res Int) {

	if i.i.BitLen()+i2.i.BitLen()-1 > maxBitLen {
		panic("Int overflow")
	}
	res = Int{mul(i.i, i2.i)}

	if res.i.BitLen() > maxBitLen {
		panic("Int overflow")
	}
	return
}


func (i Int) MulRaw(i2 int64) Int {
	return i.Mul(NewInt(i2))
}


func (i Int) Quo(i2 Int) (res Int) {

	if i2.i.Sign() == 0 {
		panic("Division by zero")
	}
	return Int{div(i.i, i2.i)}
}


func (i Int) QuoRaw(i2 int64) Int {
	return i.Quo(NewInt(i2))
}


func (i Int) Mod(i2 Int) Int {
	if i2.Sign() == 0 {
		panic("division-by-zero")
	}
	return Int{mod(i.i, i2.i)}
}


func (i Int) ModRaw(i2 int64) Int {
	return i.Mod(NewInt(i2))
}


func (i Int) Neg() (res Int) {
	return Int{neg(i.i)}
}


func (i Int) Abs() Int {
	return Int{abs(i.i)}
}


func MinInt(i1, i2 Int) Int {
	return Int{min(i1.BigInt(), i2.BigInt())}
}


func MaxInt(i, i2 Int) Int {
	return Int{max(i.BigInt(), i2.BigInt())}
}


func (i Int) String() string {
	return i.i.String()
}


func (i Int) MarshalJSON() ([]byte, error) {
	if i.i == nil {
		i.i = new(big.Int)
	}
	return marshalJSON(i.i)
}


func (i *Int) UnmarshalJSON(bz []byte) error {
	if i.i == nil {
		i.i = new(big.Int)
	}
	return unmarshalJSON(i.i, bz)
}



func marshalJSON(i encoding.TextMarshaler) ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}

	return json.Marshal(string(text))
}



func unmarshalJSON(i *big.Int, bz []byte) error {
	var text string
	if err := json.Unmarshal(bz, &text); err != nil {
		return err
	}

	return unmarshalText(i, text)
}


func (i Int) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}


func (i Int) Marshal() ([]byte, error) {
	if i.i == nil {
		i.i = new(big.Int)
	}
	return i.i.MarshalText()
}


func (i *Int) MarshalTo(data []byte) (n int, err error) {
	if i.i == nil {
		i.i = new(big.Int)
	}
	if i.i.BitLen() == 0 {
		copy(data, []byte{0x30})
		return 1, nil
	}

	bz, err := i.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}


func (i *Int) Unmarshal(data []byte) error {
	if len(data) == 0 {
		i = nil
		return nil
	}

	if i.i == nil {
		i.i = new(big.Int)
	}

	if err := i.i.UnmarshalText(data); err != nil {
		return err
	}

	if i.i.BitLen() > maxBitLen {
		return fmt.Errorf("integer out of range; got: %d, max: %d", i.i.BitLen(), maxBitLen)
	}

	return nil
}


func (i *Int) Size() int {
	bz, _ := i.Marshal()
	return len(bz)
}


func (i Int) MarshalAmino() ([]byte, error)   { return i.Marshal() }
func (i *Int) UnmarshalAmino(bz []byte) error { return i.Unmarshal(bz) }


func IntEq(t *testing.T, exp, got Int) (*testing.T, bool, string, string, string) {
	return t, exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func (ip IntProto) String() string {
	return ip.Int.String()
}
