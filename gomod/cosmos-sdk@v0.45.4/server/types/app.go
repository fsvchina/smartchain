package types

import (
	"encoding/json"
	"io"
	"time"

	"github.com/gogo/protobuf/grpc"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
)



const ServerStartTime = 5 * time.Second

type (







	AppOptions interface {
		Get(string) interface{}
	}




	Application interface {
		abci.Application

		RegisterAPIRoutes(*api.Server, config.APIConfig)



		RegisterGRPCServer(grpc.Server)



		RegisterTxService(clientCtx client.Context)


		RegisterTendermintService(clientCtx client.Context)
	}



	AppCreator func(log.Logger, dbm.DB, io.Writer, AppOptions) Application


	ModuleInitFlags func(startCmd *cobra.Command)



	ExportedApp struct {

		AppState json.RawMessage

		Validators []tmtypes.GenesisValidator

		Height int64

		ConsensusParams *abci.ConsensusParams
	}



	AppExporter func(log.Logger, dbm.DB, io.Writer, int64, bool, []string, AppOptions) (ExportedApp, error)
)
