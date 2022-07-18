package v4_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmversion "github.com/tendermint/tendermint/proto/tendermint/version"
	"github.com/tendermint/tendermint/version"

	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	tmclient "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"

	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
	feemarkettypes "github.com/tharsis/ethermint/x/feemarket/types"

	"fs.video/smartchain/app"
	v4 "fs.video/smartchain/app/upgrades/v4"
)

type UpgradeTestSuite struct {
	suite.Suite

	ctx                    sdk.Context
	app                    *app.Evmos
	consAddress            sdk.ConsAddress
	expiredOsmoClient      *tmclient.ClientState
	activeOsmoClient       *tmclient.ClientState
	expiredCosmosHubClient *tmclient.ClientState
	activeCosmosHubClient  *tmclient.ClientState
	osmoConsState          *tmclient.ConsensusState
	cosmosHubConsState     *tmclient.ConsensusState
}

func (suite *UpgradeTestSuite) SetupTest() {
	checkTx := false


	priv, err := ethsecp256k1.GenerateKey()
	suite.Require().NoError(err)
	suite.consAddress = sdk.ConsAddress(priv.PubKey().Address())

	suite.app = app.Setup(checkTx, feemarkettypes.DefaultGenesisState())
	suite.ctx = suite.app.BaseApp.NewContext(checkTx, tmproto.Header{
		Height:          1,
		ChainID:         "evmos_9001-1",
		Time:            time.Date(2022, 5, 9, 8, 0, 0, 0, time.UTC),
		ProposerAddress: suite.consAddress.Bytes(),

		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		LastBlockId: tmproto.BlockID{
			Hash: tmhash.Sum([]byte("block_id")),
			PartSetHeader: tmproto.PartSetHeader{
				Total: 11,
				Hash:  tmhash.Sum([]byte("partset_header")),
			},
		},
		AppHash:            tmhash.Sum([]byte("app")),
		DataHash:           tmhash.Sum([]byte("data")),
		EvidenceHash:       tmhash.Sum([]byte("evidence")),
		ValidatorsHash:     tmhash.Sum([]byte("validators")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators")),
		ConsensusHash:      tmhash.Sum([]byte("consensus")),
		LastResultsHash:    tmhash.Sum([]byte("last_result")),
	})


	suite.expiredOsmoClient = &tmclient.ClientState{
		ChainId:                      "osmosis-1",
		TrustLevel:                   tmclient.DefaultTrustLevel,
		TrustingPeriod:               10 * 24 * time.Hour,
		UnbondingPeriod:              14 * 24 * time.Hour,
		MaxClockDrift:                25 * time.Second,
		FrozenHeight:                 clienttypes.NewHeight(0, 0),
		LatestHeight:                 clienttypes.NewHeight(1, 3484087),
		UpgradePath:                  []string{"upgrade", "upgradedIBCState"},
		AllowUpdateAfterExpiry:       true,
		AllowUpdateAfterMisbehaviour: true,
	}


	suite.activeOsmoClient = &tmclient.ClientState{
		ChainId:                      "osmosis-1",
		TrustLevel:                   tmclient.DefaultTrustLevel,
		TrustingPeriod:               10 * 24 * time.Hour,
		UnbondingPeriod:              14 * 24 * time.Hour,
		MaxClockDrift:                25 * time.Second,
		FrozenHeight:                 clienttypes.NewHeight(0, 0),
		LatestHeight:                 clienttypes.NewHeight(1, 4264373),
		UpgradePath:                  []string{"upgrade", "upgradedIBCState"},
		AllowUpdateAfterExpiry:       true,
		AllowUpdateAfterMisbehaviour: true,
	}

	suite.expiredCosmosHubClient = &tmclient.ClientState{
		ChainId:                      "cosmoshub-4",
		TrustLevel:                   tmclient.DefaultTrustLevel,
		TrustingPeriod:               10 * 24 * time.Hour,
		UnbondingPeriod:              21 * 24 * time.Hour,
		MaxClockDrift:                20 * time.Second,
		FrozenHeight:                 clienttypes.NewHeight(0, 0),
		LatestHeight:                 clienttypes.NewHeight(4, 9659547),
		UpgradePath:                  []string{"upgrade", "upgradedIBCState"},
		AllowUpdateAfterExpiry:       true,
		AllowUpdateAfterMisbehaviour: true,
	}

	suite.activeCosmosHubClient = &tmclient.ClientState{
		ChainId:                      "cosmoshub-4",
		TrustLevel:                   tmclient.DefaultTrustLevel,
		TrustingPeriod:               10 * 24 * time.Hour,
		UnbondingPeriod:              21 * 24 * time.Hour,
		MaxClockDrift:                20 * time.Second,
		FrozenHeight:                 clienttypes.NewHeight(0, 0),
		LatestHeight:                 clienttypes.NewHeight(4, 10409568),
		UpgradePath:                  []string{"upgrade", "upgradedIBCState"},
		AllowUpdateAfterExpiry:       true,
		AllowUpdateAfterMisbehaviour: true,
	}

	suite.osmoConsState = &tmclient.ConsensusState{
		Timestamp: time.Date(2022, 5, 4, 23, 41, 9, 152600097, time.UTC),


	}

	suite.cosmosHubConsState = &tmclient.ConsensusState{
		Timestamp: time.Date(2022, 4, 29, 11, 9, 59, 595932461, time.UTC),


	}








}

func TestUpgradeTestSuite(t *testing.T) {
	s := new(UpgradeTestSuite)
	suite.Run(t, s)
}

func (suite *UpgradeTestSuite) TestUpdateIBCClients() {
	testCases := []struct {
		name     string
		malleate func()
		expError bool
	}{
		{
			"IBC clients updated successfully",
			func() {

				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ExpiredOsmosisClient, suite.expiredOsmoClient)
				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ExpiredCosmosHubClient, suite.expiredCosmosHubClient)


				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ActiveOsmosisClient, suite.activeOsmoClient)
				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ActiveCosmosHubClient, suite.activeCosmosHubClient)


				suite.app.IBCKeeper.ClientKeeper.SetClientConsensusState(suite.ctx, v4.ActiveOsmosisClient, suite.activeOsmoClient.GetLatestHeight(), suite.osmoConsState)
				suite.app.IBCKeeper.ClientKeeper.SetClientConsensusState(suite.ctx, v4.ActiveCosmosHubClient, suite.activeCosmosHubClient.GetLatestHeight(), suite.cosmosHubConsState)

				activeOsmoStore := suite.app.IBCKeeper.ClientKeeper.ClientStore(suite.ctx, v4.ActiveOsmosisClient)
				activeHubStore := suite.app.IBCKeeper.ClientKeeper.ClientStore(suite.ctx, v4.ActiveCosmosHubClient)


				tmclient.SetProcessedHeight(activeOsmoStore, suite.activeOsmoClient.LatestHeight, suite.activeOsmoClient.LatestHeight)
				tmclient.SetProcessedHeight(activeHubStore, suite.activeCosmosHubClient.LatestHeight, suite.activeCosmosHubClient.LatestHeight)

				tmclient.SetProcessedTime(activeOsmoStore, suite.activeOsmoClient.LatestHeight, suite.osmoConsState.GetTimestamp())
				tmclient.SetProcessedTime(activeHubStore, suite.activeCosmosHubClient.LatestHeight, suite.cosmosHubConsState.GetTimestamp())
			},
			false,
		},
		{
			"Osmosis IBC client update failed",
			func() {

				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ExpiredOsmosisClient, suite.expiredOsmoClient)


				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ActiveCosmosHubClient, suite.activeCosmosHubClient)
			},
			true,
		},
		{
			"Cosmos Hub IBC client update failed",
			func() {

				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ExpiredOsmosisClient, suite.expiredOsmoClient)
				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ExpiredCosmosHubClient, suite.expiredCosmosHubClient)


				suite.app.IBCKeeper.ClientKeeper.SetClientState(suite.ctx, v4.ActiveOsmosisClient, suite.activeOsmoClient)


				suite.app.IBCKeeper.ClientKeeper.SetClientConsensusState(suite.ctx, v4.ActiveOsmosisClient, suite.activeOsmoClient.GetLatestHeight(), suite.osmoConsState)

				activeOsmoStore := suite.app.IBCKeeper.ClientKeeper.ClientStore(suite.ctx, v4.ActiveOsmosisClient)


				tmclient.SetProcessedHeight(activeOsmoStore, suite.activeOsmoClient.LatestHeight, suite.activeOsmoClient.LatestHeight)
				tmclient.SetProcessedTime(activeOsmoStore, suite.activeOsmoClient.LatestHeight, suite.osmoConsState.GetTimestamp())
			},
			true,
		},
	}

	for _, tc := range testCases {
		suite.Run(fmt.Sprintf("Case %s", tc.name), func() {
			suite.SetupTest()

			tc.malleate()


			canUpdateClient := tmclient.IsMatchingClientState(*suite.expiredOsmoClient, *suite.activeOsmoClient)
			suite.Require().True(canUpdateClient, "cannot update non-matching osmosis client")

			canUpdateClient = tmclient.IsMatchingClientState(*suite.expiredCosmosHubClient, *suite.activeCosmosHubClient)
			suite.Require().True(canUpdateClient, "cannot update non-matching cosmos hub client")

			err := v4.UpdateIBCClients(suite.ctx, suite.app.IBCKeeper.ClientKeeper)
			if tc.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
