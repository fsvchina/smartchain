package types

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/bits"
	"regexp"
	"strings"
)











func NewCompactBitArray(bits int) *CompactBitArray {
	if bits <= 0 {
		return nil
	}
	nElems := (bits + 7) / 8
	if nElems <= 0 || nElems > math.MaxInt32 {



		return nil
	}
	return &CompactBitArray{
		ExtraBitsStored: uint32(bits % 8),
		Elems:           make([]byte, nElems),
	}
}


func (bA *CompactBitArray) Count() int {
	if bA == nil {
		return 0
	} else if bA.ExtraBitsStored == 0 {
		return len(bA.Elems) * 8
	}

	return (len(bA.Elems)-1)*8 + int(bA.ExtraBitsStored)
}



func (bA *CompactBitArray) GetIndex(i int) bool {
	if bA == nil {
		return false
	}
	if i < 0 || i >= bA.Count() {
		return false
	}

	return bA.Elems[i>>3]&(1<<uint8(7-(i%8))) > 0
}



func (bA *CompactBitArray) SetIndex(i int, v bool) bool {
	if bA == nil {
		return false
	}

	if i < 0 || i >= bA.Count() {
		return false
	}

	if v {
		bA.Elems[i>>3] |= (1 << uint8(7-(i%8)))
	} else {
		bA.Elems[i>>3] &= ^(1 << uint8(7-(i%8)))
	}

	return true
}




func (bA *CompactBitArray) NumTrueBitsBefore(index int) int {
	onesCount := 0
	max := bA.Count()
	if index > max {
		index = max
	}

	for elem := 0; ; elem++ {
		if elem*8+7 >= index {
			onesCount += bits.OnesCount8(bA.Elems[elem] >> (7 - (index % 8) + 1))
			return onesCount
		}
		onesCount += bits.OnesCount8(bA.Elems[elem])
	}
}


func (bA *CompactBitArray) Copy() *CompactBitArray {
	if bA == nil {
		return nil
	}

	c := make([]byte, len(bA.Elems))
	copy(c, bA.Elems)

	return &CompactBitArray{
		ExtraBitsStored: bA.ExtraBitsStored,
		Elems:           c,
	}
}


func (bA *CompactBitArray) Equal(other *CompactBitArray) bool {
	if bA == other {
		return true
	}
	if bA == nil || other == nil {
		return false
	}
	return bA.ExtraBitsStored == other.ExtraBitsStored &&
		bytes.Equal(bA.Elems, other.Elems)
}







func (bA *CompactBitArray) String() string { return bA.StringIndented("") }



func (bA *CompactBitArray) StringIndented(indent string) string {
	if bA == nil {
		return "nil-BitArray"
	}
	lines := []string{}
	bits := ""
	size := bA.Count()
	for i := 0; i < size; i++ {
		if bA.GetIndex(i) {
			bits += "x"
		} else {
			bits += "_"
		}

		if i%100 == 99 {
			lines = append(lines, bits)
			bits = ""
		}

		if i%10 == 9 {
			bits += indent
		}

		if i%50 == 49 {
			bits += indent
		}
	}

	if len(bits) > 0 {
		lines = append(lines, bits)
	}

	return fmt.Sprintf("BA{%v:%v}", size, strings.Join(lines, indent))
}



func (bA *CompactBitArray) MarshalJSON() ([]byte, error) {
	if bA == nil {
		return []byte("null"), nil
	}

	bits := `"`
	size := bA.Count()
	for i := 0; i < size; i++ {
		if bA.GetIndex(i) {
			bits += `x`
		} else {
			bits += `_`
		}
	}

	bits += `"`

	return []byte(bits), nil
}

var bitArrayJSONRegexp = regexp.MustCompile(`\A"([_x]*)"\z`)



func (bA *CompactBitArray) UnmarshalJSON(bz []byte) error {
	b := string(bz)
	if b == "null" {


		bA.ExtraBitsStored = 0
		bA.Elems = nil

		return nil
	}

	match := bitArrayJSONRegexp.FindStringSubmatch(b)
	if match == nil {
		return fmt.Errorf("bitArray in JSON should be a string of format %q but got %s", bitArrayJSONRegexp.String(), b)
	}

	bits := match[1]


	numBits := len(bits)
	bA2 := NewCompactBitArray(numBits)
	for i := 0; i < numBits; i++ {
		if bits[i] == 'x' {
			bA2.SetIndex(i, true)
		}
	}
	*bA = *bA2

	return nil
}



func (bA *CompactBitArray) CompactMarshal() []byte {
	size := bA.Count()
	if size <= 0 {
		return []byte("null")
	}

	bz := make([]byte, 0, size/8)



	bz = appendUvarint(bz, uint64(size))
	bz = append(bz, bA.Elems...)

	return bz
}



func CompactUnmarshal(bz []byte) (*CompactBitArray, error) {
	if len(bz) < 2 {
		return nil, errors.New("compact bit array: invalid compact unmarshal size")
	} else if bytes.Equal(bz, []byte("null")) {
		return NewCompactBitArray(0), nil
	}

	size, n := binary.Uvarint(bz)
	if n < 0 || n >= len(bz) {
		return nil, fmt.Errorf("compact bit array: n=%d is out of range of len(bz)=%d", n, len(bz))
	}
	bz = bz[n:]

	if len(bz) != int(size+7)/8 {
		return nil, errors.New("compact bit array: invalid compact unmarshal size")
	}

	bA := &CompactBitArray{uint32(size % 8), bz}

	return bA, nil
}

func appendUvarint(b []byte, x uint64) []byte {
	var a [binary.MaxVarintLen64]byte
	n := binary.PutUvarint(a[:], x)

	return append(b, a[:n]...)
}
