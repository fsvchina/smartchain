package v040

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	v039auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v039"
	v040auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v040"
	v036supply "github.com/cosmos/cosmos-sdk/x/bank/legacy/v036"
	v038bank "github.com/cosmos/cosmos-sdk/x/bank/legacy/v038"
	v040bank "github.com/cosmos/cosmos-sdk/x/bank/legacy/v040"
	v039crisis "github.com/cosmos/cosmos-sdk/x/crisis/legacy/v039"
	v040crisis "github.com/cosmos/cosmos-sdk/x/crisis/legacy/v040"
	v036distr "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v036"
	v038distr "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v038"
	v040distr "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v040"
	v038evidence "github.com/cosmos/cosmos-sdk/x/evidence/legacy/v038"
	v040evidence "github.com/cosmos/cosmos-sdk/x/evidence/legacy/v040"
	v039genutil "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v039"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
	v036gov "github.com/cosmos/cosmos-sdk/x/gov/legacy/v036"
	v040gov "github.com/cosmos/cosmos-sdk/x/gov/legacy/v040"
	v039mint "github.com/cosmos/cosmos-sdk/x/mint/legacy/v039"
	v040mint "github.com/cosmos/cosmos-sdk/x/mint/legacy/v040"
	v036params "github.com/cosmos/cosmos-sdk/x/params/legacy/v036"
	v039slashing "github.com/cosmos/cosmos-sdk/x/slashing/legacy/v039"
	v040slashing "github.com/cosmos/cosmos-sdk/x/slashing/legacy/v040"
	v038staking "github.com/cosmos/cosmos-sdk/x/staking/legacy/v038"
	v040staking "github.com/cosmos/cosmos-sdk/x/staking/legacy/v040"
	v038upgrade "github.com/cosmos/cosmos-sdk/x/upgrade/legacy/v038"
)

func migrateGenutil(oldGenState v039genutil.GenesisState) *types.GenesisState {
	return &types.GenesisState{
		GenTxs: oldGenState.GenTxs,
	}
}


func Migrate(appState types.AppMap, clientCtx client.Context) types.AppMap {
	v039Codec := codec.NewLegacyAmino()
	v039auth.RegisterLegacyAminoCodec(v039Codec)
	v036gov.RegisterLegacyAminoCodec(v039Codec)
	v036distr.RegisterLegacyAminoCodec(v039Codec)
	v036params.RegisterLegacyAminoCodec(v039Codec)
	v038upgrade.RegisterLegacyAminoCodec(v039Codec)

	v040Codec := clientCtx.Codec

	if appState[v038bank.ModuleName] != nil {

		var bankGenState v038bank.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v038bank.ModuleName], &bankGenState)


		var authGenState v039auth.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039auth.ModuleName], &authGenState)


		var supplyGenState v036supply.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v036supply.ModuleName], &supplyGenState)


		delete(appState, v038bank.ModuleName)


		delete(appState, v036supply.ModuleName)



		appState[v040bank.ModuleName] = v040Codec.MustMarshalJSON(v040bank.Migrate(bankGenState, authGenState, supplyGenState))
	}


	if appState[v039auth.ModuleName] != nil {

		var authGenState v039auth.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039auth.ModuleName], &authGenState)


		delete(appState, v039auth.ModuleName)



		appState[v040auth.ModuleName] = v040Codec.MustMarshalJSON(v040auth.Migrate(authGenState))
	}


	if appState[v039crisis.ModuleName] != nil {

		var crisisGenState v039crisis.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039crisis.ModuleName], &crisisGenState)


		delete(appState, v039crisis.ModuleName)



		appState[v040crisis.ModuleName] = v040Codec.MustMarshalJSON(v040crisis.Migrate(crisisGenState))
	}


	if appState[v038distr.ModuleName] != nil {

		var distributionGenState v038distr.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v038distr.ModuleName], &distributionGenState)


		delete(appState, v038distr.ModuleName)



		appState[v040distr.ModuleName] = v040Codec.MustMarshalJSON(v040distr.Migrate(distributionGenState))
	}


	if appState[v038evidence.ModuleName] != nil {

		var evidenceGenState v038evidence.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v038bank.ModuleName], &evidenceGenState)


		delete(appState, v038evidence.ModuleName)



		appState[v040evidence.ModuleName] = v040Codec.MustMarshalJSON(v040evidence.Migrate(evidenceGenState))
	}


	if appState[v036gov.ModuleName] != nil {

		var govGenState v036gov.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v036gov.ModuleName], &govGenState)


		delete(appState, v036gov.ModuleName)



		appState[v040gov.ModuleName] = v040Codec.MustMarshalJSON(v040gov.Migrate(govGenState))
	}


	if appState[v039mint.ModuleName] != nil {

		var mintGenState v039mint.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039mint.ModuleName], &mintGenState)


		delete(appState, v039mint.ModuleName)



		appState[v040mint.ModuleName] = v040Codec.MustMarshalJSON(v040mint.Migrate(mintGenState))
	}


	if appState[v039slashing.ModuleName] != nil {

		var slashingGenState v039slashing.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039slashing.ModuleName], &slashingGenState)


		delete(appState, v039slashing.ModuleName)



		appState[v040slashing.ModuleName] = v040Codec.MustMarshalJSON(v040slashing.Migrate(slashingGenState))
	}


	if appState[v038staking.ModuleName] != nil {

		var stakingGenState v038staking.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v038staking.ModuleName], &stakingGenState)


		delete(appState, v038staking.ModuleName)



		appState[v040staking.ModuleName] = v040Codec.MustMarshalJSON(v040staking.Migrate(stakingGenState))
	}


	if appState[v039genutil.ModuleName] != nil {

		var genutilGenState v039genutil.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v039genutil.ModuleName], &genutilGenState)


		delete(appState, v039genutil.ModuleName)



		appState[ModuleName] = v040Codec.MustMarshalJSON(migrateGenutil(genutilGenState))
	}

	return appState
}
