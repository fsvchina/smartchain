package mock

import (
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"github.com/tendermint/tendermint/types"
)


func TestInitApp(t *testing.T) {

	app, closer, err := SetupApp()


	if closer != nil {
		defer closer()
	}
	require.NoError(t, err)


	appState, err := AppGenState(nil, types.GenesisDoc{}, nil)
	require.NoError(t, err)


	req := abci.RequestInitChain{
		AppStateBytes: appState,
	}
	app.InitChain(req)
	app.Commit()


	query := abci.RequestQuery{
		Path: "/store/main/key",
		Data: []byte("foo"),
	}
	qres := app.Query(query)
	require.Equal(t, uint32(0), qres.Code, qres.Log)
	require.Equal(t, []byte("bar"), qres.Value)
}


func TestDeliverTx(t *testing.T) {

	app, closer, err := SetupApp()

	if closer != nil {
		defer closer()
	}
	require.NoError(t, err)

	key := "my-special-key"
	value := "top-secret-data!!"
	tx := NewTx(key, value)
	txBytes := tx.GetSignBytes()

	header := tmproto.Header{
		AppHash: []byte("apphash"),
		Height:  1,
	}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})
	dres := app.DeliverTx(abci.RequestDeliverTx{Tx: txBytes})
	require.Equal(t, uint32(0), dres.Code, dres.Log)
	app.EndBlock(abci.RequestEndBlock{})
	cres := app.Commit()
	require.NotEmpty(t, cres.Data)


	query := abci.RequestQuery{
		Path: "/store/main/key",
		Data: []byte(key),
	}
	qres := app.Query(query)
	require.Equal(t, uint32(0), qres.Code, qres.Log)
	require.Equal(t, []byte(value), qres.Value)
}
