package types


const (
	EventTypeSubmitProposal   = "submit_proposal"
	EventTypeProposalDeposit  = "proposal_deposit"
	EventTypeProposalVote     = "proposal_vote"
	EventTypeInactiveProposal = "inactive_proposal"
	EventTypeActiveProposal   = "active_proposal"

	AttributeKeyProposalResult     = "proposal_result"
	AttributeKeyOption             = "option"
	AttributeKeyProposalID         = "proposal_id"
	AttributeKeyVotingPeriodStart  = "voting_period_start"
	AttributeValueCategory         = "governance"
	AttributeValueProposalDropped  = "proposal_dropped"  
	AttributeValueProposalPassed   = "proposal_passed"   
	AttributeValueProposalRejected = "proposal_rejected" 
	AttributeValueProposalFailed   = "proposal_failed"   
	AttributeKeyProposalType       = "proposal_type"
)
