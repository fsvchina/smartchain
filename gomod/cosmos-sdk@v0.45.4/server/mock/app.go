package mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/tendermint/tendermint/types"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)




func NewApp(rootDir string, logger log.Logger) (abci.Application, error) {
	db, err := sdk.NewLevelDB("mock", filepath.Join(rootDir, "data"))
	if err != nil {
		return nil, err
	}


	capKeyMainStore := sdk.NewKVStoreKey("main")


	baseApp := bam.NewBaseApp("kvstore", logger, db, decodeTx)


	baseApp.MountStores(capKeyMainStore)

	baseApp.SetInitChainer(InitChainer(capKeyMainStore))


	baseApp.Router().AddRoute(sdk.NewRoute("kvstore", KVStoreHandler(capKeyMainStore)))


	if err := baseApp.LoadLatestVersion(); err != nil {
		return nil, err
	}

	return baseApp, nil
}



func KVStoreHandler(storeKey sdk.StoreKey) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		dTx, ok := msg.(kvstoreTx)
		if !ok {
			return nil, errors.New("KVStoreHandler should only receive kvstoreTx")
		}


		key := dTx.key
		value := dTx.value

		store := ctx.KVStore(storeKey)
		store.Set(key, value)

		return &sdk.Result{
			Log: fmt.Sprintf("set %s=%s", key, value),
		}, nil
	}
}


type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}


type GenesisJSON struct {
	Values []KV `json:"values"`
}



func InitChainer(key sdk.StoreKey) func(sdk.Context, abci.RequestInitChain) abci.ResponseInitChain {
	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		stateJSON := req.AppStateBytes

		genesisState := new(GenesisJSON)
		err := json.Unmarshal(stateJSON, genesisState)
		if err != nil {
			panic(err)

		}

		for _, val := range genesisState.Values {
			store := ctx.KVStore(key)
			store.Set([]byte(val.Key), []byte(val.Value))
		}
		return abci.ResponseInitChain{}
	}
}



func AppGenState(_ *codec.LegacyAmino, _ types.GenesisDoc, _ []json.RawMessage) (appState json.
	RawMessage, err error) {
	appState = json.RawMessage(`{
  "values": [
    {
        "key": "hello",
        "value": "goodbye"
    },
    {
        "key": "foo",
        "value": "bar"
    }
  ]
}`)
	return
}


func AppGenStateEmpty(_ *codec.LegacyAmino, _ types.GenesisDoc, _ []json.RawMessage) (
	appState json.RawMessage, err error) {
	appState = json.RawMessage(``)
	return
}
