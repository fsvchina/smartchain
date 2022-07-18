package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)




const (
	QueryParams    = "params"
	QueryProposals = "proposals"
	QueryProposal  = "proposal"
	QueryDeposits  = "deposits"
	QueryDeposit   = "deposit"
	QueryVotes     = "votes"
	QueryVote      = "vote"
	QueryTally     = "tally"

	ParamDeposit  = "deposit"
	ParamVoting   = "voting"
	ParamTallying = "tallying"
)





type QueryProposalParams struct {
	ProposalID uint64
}


func NewQueryProposalParams(proposalID uint64) QueryProposalParams {
	return QueryProposalParams{
		ProposalID: proposalID,
	}
}


type QueryProposalVotesParams struct {
	ProposalID uint64
	Page       int
	Limit      int
}


func NewQueryProposalVotesParams(proposalID uint64, page, limit int) QueryProposalVotesParams {
	return QueryProposalVotesParams{
		ProposalID: proposalID,
		Page:       page,
		Limit:      limit,
	}
}


type QueryDepositParams struct {
	ProposalID uint64
	Depositor  sdk.AccAddress
}


func NewQueryDepositParams(proposalID uint64, depositor sdk.AccAddress) QueryDepositParams {
	return QueryDepositParams{
		ProposalID: proposalID,
		Depositor:  depositor,
	}
}


type QueryVoteParams struct {
	ProposalID uint64
	Voter      sdk.AccAddress
}


func NewQueryVoteParams(proposalID uint64, voter sdk.AccAddress) QueryVoteParams {
	return QueryVoteParams{
		ProposalID: proposalID,
		Voter:      voter,
	}
}


type QueryProposalsParams struct {
	Page           int
	Limit          int
	Voter          sdk.AccAddress
	Depositor      sdk.AccAddress
	ProposalStatus ProposalStatus
}


func NewQueryProposalsParams(page, limit int, status ProposalStatus, voter, depositor sdk.AccAddress) QueryProposalsParams {
	return QueryProposalsParams{
		Page:           page,
		Limit:          limit,
		Voter:          voter,
		Depositor:      depositor,
		ProposalStatus: status,
	}
}
