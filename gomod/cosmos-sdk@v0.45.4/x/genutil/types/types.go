package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)



type (

	AppMap map[string]json.RawMessage



	//

	MigrationCallback func(AppMap, client.Context) AppMap


	MigrationMap map[string]MigrationCallback
)


const ModuleName = "genutil"


type InitConfig struct {
	ChainID   string
	GenTxsDir string
	NodeID    string
	ValPubKey cryptotypes.PubKey
}


func NewInitConfig(chainID, genTxsDir, nodeID string, valPubKey cryptotypes.PubKey) InitConfig {
	return InitConfig{
		ChainID:   chainID,
		GenTxsDir: genTxsDir,
		NodeID:    nodeID,
		ValPubKey: valPubKey,
	}
}
