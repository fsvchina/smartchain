package types

import (
	"fmt"
	"testing"
)

func BenchmarkParseChainID(b *testing.B) {
	b.ReportAllocs()
	
	for i := 1; i < b.N; i++ {
		chainID := fmt.Sprintf("ethermint_1-%d", i)
		if _, err := ParseChainID(chainID); err != nil {
			b.Fatal(err)
		}
	}
}
