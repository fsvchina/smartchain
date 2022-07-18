package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)


type ParamSubspace interface {
	Get(ctx sdk.Context, key []byte, ptr interface{})
	Set(ctx sdk.Context, key []byte, param interface{})
}


type StakingKeeper interface {

	IterateBondedValidatorsByPower(
		sdk.Context, func(index int64, validator stakingtypes.ValidatorI) (stop bool),
	)

	TotalBondedTokens(sdk.Context) sdk.Int
	IterateDelegations(
		ctx sdk.Context, delegator sdk.AccAddress,
		fn func(index int64, delegation stakingtypes.DelegationI) (stop bool),
	)
}


type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI

	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) types.ModuleAccountI


	SetModuleAccount(sdk.Context, types.ModuleAccountI)
}


type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) error
}






type GovHooks interface {
	AfterProposalSubmission(ctx sdk.Context, proposalID uint64)
	AfterProposalDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress)
	AfterProposalVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress)
	AfterProposalFailedMinDeposit(ctx sdk.Context, proposalID uint64)
	AfterProposalVotingPeriodEnded(ctx sdk.Context, proposalID uint64)
}
