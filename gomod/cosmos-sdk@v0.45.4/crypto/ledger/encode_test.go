package ledger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

type byter interface {
	Bytes() []byte
}

func checkAminoJSON(t *testing.T, src interface{}, dst interface{}, isNil bool) {

	js, err := cdc.MarshalJSON(src)
	require.Nil(t, err, "%+v", err)
	if isNil {
		require.Equal(t, string(js), `null`)
	} else {
		require.Contains(t, string(js), `"type":`)
		require.Contains(t, string(js), `"value":`)
	}

	err = cdc.UnmarshalJSON(js, dst)
	require.Nil(t, err, "%+v", err)
}


func ExamplePrintRegisteredTypes() {
	cdc.PrintTypes(os.Stdout)










}

func TestNilEncodings(t *testing.T) {


	var a, b []byte
	checkAminoJSON(t, &a, &b, true)
	require.EqualValues(t, a, b)


	var c, d cryptotypes.PubKey
	checkAminoJSON(t, &c, &d, true)
	require.EqualValues(t, c, d)


	var e, f cryptotypes.PrivKey
	checkAminoJSON(t, &e, &f, true)
	require.EqualValues(t, e, f)

}
