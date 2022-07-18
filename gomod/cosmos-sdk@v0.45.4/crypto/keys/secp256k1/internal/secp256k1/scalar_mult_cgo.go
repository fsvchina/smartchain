






package secp256k1

import (
	"math/big"
	"unsafe"
)

include/secp256k1.h"

extern int secp256k1_ext_scalar_mul(const secp256k1_context* ctx, const unsigned char *point, const unsigned char *scalar);

*/
import "C"

func (BitCurve *BitCurve) ScalarMult(Bx, By *big.Int, scalar []byte) (*big.Int, *big.Int) {


	if len(scalar) > 32 {
		panic("can't handle scalars > 256 bits")
	}

	padded := make([]byte, 32)
	copy(padded[32-len(scalar):], scalar)
	scalar = padded


	point := make([]byte, 64)
	readBits(Bx, point[:32])
	readBits(By, point[32:])

	pointPtr := (*C.uchar)(unsafe.Pointer(&point[0]))
	scalarPtr := (*C.uchar)(unsafe.Pointer(&scalar[0]))
	res := C.secp256k1_ext_scalar_mul(context, pointPtr, scalarPtr)


	x := new(big.Int).SetBytes(point[:32])
	y := new(big.Int).SetBytes(point[32:])
	for i := range point {
		point[i] = 0
	}
	for i := range padded {
		scalar[i] = 0
	}
	if res != 1 {
		return nil, nil
	}
	return x, y
}
