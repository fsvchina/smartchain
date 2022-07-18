package capability_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	"github.com/cosmos/cosmos-sdk/x/capability/keeper"
	"github.com/cosmos/cosmos-sdk/x/capability/types"
)

type CapabilityTestSuite struct {
	suite.Suite

	cdc    codec.Codec
	ctx    sdk.Context
	app    *simapp.SimApp
	keeper *keeper.Keeper
	module module.AppModule
}

func (suite *CapabilityTestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(checkTx)
	cdc := app.AppCodec()


	keeper := keeper.NewKeeper(cdc, app.GetKey(types.StoreKey), app.GetMemKey(types.MemStoreKey))

	suite.app = app
	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1})
	suite.keeper = keeper
	suite.cdc = cdc
	suite.module = capability.NewAppModule(cdc, *keeper)
}



func (suite *CapabilityTestSuite) TestInitializeMemStore() {
	sk1 := suite.keeper.ScopeToModule(banktypes.ModuleName)

	cap1, err := sk1.NewCapability(suite.ctx, "transfer")
	suite.Require().NoError(err)
	suite.Require().NotNil(cap1)


	newKeeper := keeper.NewKeeper(suite.cdc, suite.app.GetKey(types.StoreKey), suite.app.GetMemKey("testingkey"))
	newSk1 := newKeeper.ScopeToModule(banktypes.ModuleName)


	ctx := suite.app.BaseApp.NewUncachedContext(false, tmproto.Header{})
	newKeeper.Seal()
	suite.Require().False(newKeeper.IsInitialized(ctx), "memstore initialized flag set before BeginBlock")


	ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{}).WithBlockGasMeter(sdk.NewGasMeter(50))
	prevGas := ctx.BlockGasMeter().GasConsumed()
	restartedModule := capability.NewAppModule(suite.cdc, *newKeeper)
	restartedModule.BeginBlock(ctx, abci.RequestBeginBlock{})
	suite.Require().True(newKeeper.IsInitialized(ctx), "memstore initialized flag not set")
	gasUsed := ctx.BlockGasMeter().GasConsumed()

	suite.Require().Equal(prevGas, gasUsed, "beginblocker consumed gas during execution")



	cacheCtx, _ := ctx.CacheContext()
	_, ok := newSk1.GetCapability(cacheCtx, "transfer")
	suite.Require().True(ok)


	ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})

	cap1, ok = newSk1.GetCapability(ctx, "transfer")
	suite.Require().True(ok)




	restartedModule.BeginBlock(ctx, abci.RequestBeginBlock{})
	recap, ok := newSk1.GetCapability(ctx, "transfer")
	suite.Require().True(ok)
	suite.Require().Equal(cap1, recap, "capabilities got reinitialized after second BeginBlock")
	suite.Require().True(newKeeper.IsInitialized(ctx), "memstore initialized flag not set")
}

func TestCapabilityTestSuite(t *testing.T) {
	suite.Run(t, new(CapabilityTestSuite))
}
