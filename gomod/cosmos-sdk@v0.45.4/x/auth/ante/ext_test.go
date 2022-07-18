package ante_test

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

func (suite *AnteTestSuite) TestRejectExtensionOptionsDecorator() {
	suite.SetupTest(true)
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	reod := ante.NewRejectExtensionOptionsDecorator()
	antehandler := sdk.ChainAnteDecorators(reod)


	theTx := suite.txBuilder.GetTx()
	_, err := antehandler(suite.ctx, theTx, false)
	suite.Require().NoError(err)

	extOptsTxBldr, ok := suite.txBuilder.(tx.ExtensionOptionsTxBuilder)
	if !ok {

		return
	}


	any, err := types.NewAnyWithValue(testdata.NewTestMsg())
	suite.Require().NoError(err)
	extOptsTxBldr.SetExtensionOptions(any)
	theTx = suite.txBuilder.GetTx()
	_, err = antehandler(suite.ctx, theTx, false)
	suite.Require().EqualError(err, "unknown extension options")
}
