package types

import (
	"errors"
	"fmt"
	"math/big"
)




type Uint struct {
	i *big.Int
}


func (u Uint) BigInt() *big.Int {
	return new(big.Int).Set(u.i)
}


func NewUintFromBigInt(i *big.Int) Uint {
	u, err := checkNewUint(i)
	if err != nil {
		panic(fmt.Errorf("overflow: %s", err))
	}
	return u
}


func NewUint(n uint64) Uint {
	i := new(big.Int)
	i.SetUint64(n)
	return NewUintFromBigInt(i)
}


func NewUintFromString(s string) Uint {
	u, err := ParseUint(s)
	if err != nil {
		panic(err)
	}
	return u
}


func ZeroUint() Uint { return Uint{big.NewInt(0)} }


func OneUint() Uint { return Uint{big.NewInt(1)} }

var _ CustomProtobufType = (*Uint)(nil)



func (u Uint) Uint64() uint64 {
	if !u.i.IsUint64() {
		panic("Uint64() out of bound")
	}
	return u.i.Uint64()
}


func (u Uint) IsZero() bool { return u.Equal(ZeroUint()) }


func (u Uint) Equal(u2 Uint) bool { return equal(u.i, u2.i) }


func (u Uint) GT(u2 Uint) bool { return gt(u.i, u2.i) }


func (u Uint) GTE(u2 Uint) bool { return u.GT(u2) || u.Equal(u2) }


func (u Uint) LT(u2 Uint) bool { return lt(u.i, u2.i) }


func (u Uint) LTE(u2 Uint) bool { return !u.GT(u2) }


func (u Uint) Add(u2 Uint) Uint { return NewUintFromBigInt(new(big.Int).Add(u.i, u2.i)) }


func (u Uint) AddUint64(u2 uint64) Uint { return u.Add(NewUint(u2)) }


func (u Uint) Sub(u2 Uint) Uint { return NewUintFromBigInt(new(big.Int).Sub(u.i, u2.i)) }


func (u Uint) SubUint64(u2 uint64) Uint { return u.Sub(NewUint(u2)) }


func (u Uint) Mul(u2 Uint) (res Uint) {
	return NewUintFromBigInt(new(big.Int).Mul(u.i, u2.i))
}


func (u Uint) MulUint64(u2 uint64) (res Uint) { return u.Mul(NewUint(u2)) }


func (u Uint) Quo(u2 Uint) (res Uint) { return NewUintFromBigInt(div(u.i, u2.i)) }


func (u Uint) Mod(u2 Uint) Uint {
	if u2.IsZero() {
		panic("division-by-zero")
	}
	return Uint{mod(u.i, u2.i)}
}


func (u Uint) Incr() Uint {
	return u.Add(OneUint())
}



func (u Uint) Decr() Uint {
	return u.Sub(OneUint())
}


func (u Uint) QuoUint64(u2 uint64) Uint { return u.Quo(NewUint(u2)) }


func MinUint(u1, u2 Uint) Uint { return NewUintFromBigInt(min(u1.i, u2.i)) }


func MaxUint(u1, u2 Uint) Uint { return NewUintFromBigInt(max(u1.i, u2.i)) }


func (u Uint) String() string { return u.i.String() }


func (u Uint) MarshalJSON() ([]byte, error) {
	if u.i == nil {
		u.i = new(big.Int)
	}
	return marshalJSON(u.i)
}


func (u *Uint) UnmarshalJSON(bz []byte) error {
	if u.i == nil {
		u.i = new(big.Int)
	}
	return unmarshalJSON(u.i, bz)
}


func (u Uint) Marshal() ([]byte, error) {
	if u.i == nil {
		u.i = new(big.Int)
	}
	return u.i.MarshalText()
}


func (u *Uint) MarshalTo(data []byte) (n int, err error) {
	if u.i == nil {
		u.i = new(big.Int)
	}
	if u.i.BitLen() == 0 {
		copy(data, []byte{0x30})
		return 1, nil
	}

	bz, err := u.Marshal()
	if err != nil {
		return 0, err
	}

	copy(data, bz)
	return len(bz), nil
}


func (u *Uint) Unmarshal(data []byte) error {
	if len(data) == 0 {
		u = nil
		return nil
	}

	if u.i == nil {
		u.i = new(big.Int)
	}

	if err := u.i.UnmarshalText(data); err != nil {
		return err
	}

	if u.i.BitLen() > maxBitLen {
		return fmt.Errorf("integer out of range; got: %d, max: %d", u.i.BitLen(), maxBitLen)
	}

	return nil
}


func (u *Uint) Size() int {
	bz, _ := u.Marshal()
	return len(bz)
}


func (u Uint) MarshalAmino() ([]byte, error)   { return u.Marshal() }
func (u *Uint) UnmarshalAmino(bz []byte) error { return u.Unmarshal(bz) }



func UintOverflow(i *big.Int) error {
	if i.Sign() < 0 {
		return errors.New("non-positive integer")
	}
	if i.BitLen() > 256 {
		return fmt.Errorf("bit length %d greater than 256", i.BitLen())
	}
	return nil
}


func ParseUint(s string) (Uint, error) {
	i, ok := new(big.Int).SetString(s, 0)
	if !ok {
		return Uint{}, fmt.Errorf("cannot convert %q to big.Int", s)
	}
	return checkNewUint(i)
}

func checkNewUint(i *big.Int) (Uint, error) {
	if err := UintOverflow(i); err != nil {
		return Uint{}, err
	}
	return Uint{i}, nil
}



func RelativePow(x Uint, n Uint, b Uint) (z Uint) {
	if x.IsZero() {
		if n.IsZero() {
			z = b
			return
		}
		z = ZeroUint()
		return
	}

	z = x
	if n.Mod(NewUint(2)).Equal(ZeroUint()) {
		z = b
	}

	halfOfB := b.Quo(NewUint(2))
	n = n.Quo(NewUint(2))

	for n.GT(ZeroUint()) {
		xSquared := x.Mul(x)
		xSquaredRounded := xSquared.Add(halfOfB)

		x = xSquaredRounded.Quo(b)

		if n.Mod(NewUint(2)).Equal(OneUint()) {
			zx := z.Mul(x)
			zxRounded := zx.Add(halfOfB)
			z = zxRounded.Quo(b)
		}
		n = n.Quo(NewUint(2))
	}
	return z
}
