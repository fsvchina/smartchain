package v043_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v040slashing "github.com/cosmos/cosmos-sdk/x/slashing/legacy/v040"
	v043slashing "github.com/cosmos/cosmos-sdk/x/slashing/legacy/v043"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)

func TestStoreMigration(t *testing.T) {
	slashingKey := sdk.NewKVStoreKey("slashing")
	ctx := testutil.DefaultContext(slashingKey, sdk.NewTransientStoreKey("transient_test"))
	store := ctx.KVStore(slashingKey)

	_, _, addr1 := testdata.KeyTestPubAddr()
	consAddr := sdk.ConsAddress(addr1)

	value := []byte("foo")

	testCases := []struct {
		name   string
		oldKey []byte
		newKey []byte
	}{
		{
			"ValidatorSigningInfoKey",
			v040slashing.ValidatorSigningInfoKey(consAddr),
			types.ValidatorSigningInfoKey(consAddr),
		},
		{
			"ValidatorMissedBlockBitArrayKey",
			v040slashing.ValidatorMissedBlockBitArrayKey(consAddr, 2),
			types.ValidatorMissedBlockBitArrayKey(consAddr, 2),
		},
		{
			"AddrPubkeyRelationKey",
			v040slashing.AddrPubkeyRelationKey(consAddr),
			types.AddrPubkeyRelationKey(consAddr),
		},
	}


	for _, tc := range testCases {
		store.Set(tc.oldKey, value)
	}


	err := v043slashing.MigrateStore(ctx, slashingKey)
	require.NoError(t, err)


	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if !bytes.Equal(tc.oldKey, tc.newKey) {
				require.Nil(t, store.Get(tc.oldKey))
			}
			require.Equal(t, value, store.Get(tc.newKey))
		})
	}
}
