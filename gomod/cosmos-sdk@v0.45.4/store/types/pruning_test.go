package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPruningOptions_Validate(t *testing.T) {
	testCases := []struct {
		keepRecent uint64
		keepEvery  uint64
		interval   uint64
		expectErr  bool
	}{
		{100, 500, 10, false},
		{0, 0, 10, false},
		{0, 1, 0, false},
		{0, 10, 10, false},
		{100, 0, 0, true},
		{0, 1, 5, true},
	}

	for _, tc := range testCases {
		po := NewPruningOptions(tc.keepRecent, tc.keepEvery, tc.interval)
		err := po.Validate()
		require.Equal(t, tc.expectErr, err != nil, "options: %v, err: %s", po, err)
	}
}
