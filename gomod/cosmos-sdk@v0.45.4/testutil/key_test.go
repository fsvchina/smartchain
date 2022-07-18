package testutil

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types"
)

func TestGenerateCoinKey(t *testing.T) {
	t.Parallel()
	addr, mnemonic, err := GenerateCoinKey(hd.Secp256k1)
	require.NoError(t, err)


	info, err := keyring.NewInMemory().NewAccount("xxx", mnemonic, "", hd.NewFundraiserParams(0, types.GetConfig().GetCoinType(), 0).String(), hd.Secp256k1)
	require.NoError(t, err)
	require.Equal(t, addr, info.GetAddress())
}

func TestGenerateSaveCoinKey(t *testing.T) {
	t.Parallel()

	kb, err := keyring.New(t.Name(), "test", t.TempDir(), nil)
	require.NoError(t, err)

	addr, mnemonic, err := GenerateSaveCoinKey(kb, "keyname", "", false, hd.Secp256k1)
	require.NoError(t, err)


	info, err := kb.Key("keyname")
	require.NoError(t, err)
	require.Equal(t, addr, info.GetAddress())


	info, err = keyring.NewInMemory().NewAccount("xxx", mnemonic, "", hd.NewFundraiserParams(0, types.GetConfig().GetCoinType(), 0).String(), hd.Secp256k1)
	require.NoError(t, err)
	require.Equal(t, addr, info.GetAddress())
}

func TestGenerateSaveCoinKeyOverwriteFlag(t *testing.T) {
	t.Parallel()

	kb, err := keyring.New(t.Name(), "test", t.TempDir(), nil)
	require.NoError(t, err)

	keyname := "justakey"
	addr1, _, err := GenerateSaveCoinKey(kb, keyname, "", false, hd.Secp256k1)
	require.NoError(t, err)


	_, _, err = GenerateSaveCoinKey(kb, keyname, "", false, hd.Secp256k1)
	require.Error(t, err)


	addr2, _, err := GenerateSaveCoinKey(kb, keyname, "", true, hd.Secp256k1)
	require.NoError(t, err)

	require.NotEqual(t, addr1, addr2)
}