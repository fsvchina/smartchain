package ante_test

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

func (suite *AnteTestSuite) TestSetup() {
	suite.SetupTest(true) 
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	
	msg := testdata.NewTestMsg(addr1)
	feeAmount := testdata.NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	suite.Require().NoError(suite.txBuilder.SetMsgs(msg))
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
	suite.Require().NoError(err)

	sud := ante.NewSetUpContextDecorator()
	antehandler := sdk.ChainAnteDecorators(sud)

	
	suite.ctx = suite.ctx.WithBlockHeight(1)

	
	suite.Require().Equal(uint64(0), suite.ctx.GasMeter().Limit(), "GasMeter set with limit before setup")

	newCtx, err := antehandler(suite.ctx, tx, false)
	suite.Require().Nil(err, "SetUpContextDecorator returned error")

	
	suite.Require().Equal(gasLimit, newCtx.GasMeter().Limit(), "GasMeter not set correctly")
}

func (suite *AnteTestSuite) TestRecoverPanic() {
	suite.SetupTest(true) 
	suite.txBuilder = suite.clientCtx.TxConfig.NewTxBuilder()

	
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	
	msg := testdata.NewTestMsg(addr1)
	feeAmount := testdata.NewTestFeeAmount()
	gasLimit := testdata.NewTestGasLimit()
	suite.Require().NoError(suite.txBuilder.SetMsgs(msg))
	suite.txBuilder.SetFeeAmount(feeAmount)
	suite.txBuilder.SetGasLimit(gasLimit)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	tx, err := suite.CreateTestTx(privs, accNums, accSeqs, suite.ctx.ChainID())
	suite.Require().NoError(err)

	sud := ante.NewSetUpContextDecorator()
	antehandler := sdk.ChainAnteDecorators(sud, OutOfGasDecorator{})

	
	suite.ctx = suite.ctx.WithBlockHeight(1)

	newCtx, err := antehandler(suite.ctx, tx, false)

	suite.Require().NotNil(err, "Did not return error on OutOfGas panic")

	suite.Require().True(sdkerrors.ErrOutOfGas.Is(err), "Returned error is not an out of gas error")
	suite.Require().Equal(gasLimit, newCtx.GasMeter().Limit())

	antehandler = sdk.ChainAnteDecorators(sud, PanicDecorator{})
	suite.Require().Panics(func() { antehandler(suite.ctx, tx, false) }, "Recovered from non-Out-of-Gas panic") 
}

type OutOfGasDecorator struct{}


func (ogd OutOfGasDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	overLimit := ctx.GasMeter().Limit() + 1

	
	ctx.GasMeter().ConsumeGas(overLimit, "test panic")

	
	return next(ctx, tx, simulate)
}

type PanicDecorator struct{}

func (pd PanicDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	panic("random error")
}
