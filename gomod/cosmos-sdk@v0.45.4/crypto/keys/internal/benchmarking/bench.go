package benchmarking

import (
	"crypto/rand"
	"io"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/types"
)









func BenchmarkKeyGeneration(b *testing.B, generateKey func(reader io.Reader) types.PrivKey) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		generateKey(rand.Reader)
	}
}



func BenchmarkSigning(b *testing.B, priv types.PrivKey) {
	message := []byte("Hello, world!")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := priv.Sign(message)

		if err != nil {
			b.FailNow()
		}
	}
}



func BenchmarkVerification(b *testing.B, priv types.PrivKey) {
	pub := priv.PubKey()

	message := []byte("Hello, world!")
	signature, err := priv.Sign(message)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pub.VerifySignature(message, signature)
	}
}






























