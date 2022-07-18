package types

import (
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	query "github.com/cosmos/cosmos-sdk/types/query"
)


const (
	QueryEvidence    = "evidence"
	QueryAllEvidence = "all_evidence"
)


func NewQueryEvidenceRequest(hash tmbytes.HexBytes) *QueryEvidenceRequest {
	return &QueryEvidenceRequest{EvidenceHash: hash}
}


func NewQueryAllEvidenceRequest(pageReq *query.PageRequest) *QueryAllEvidenceRequest {
	return &QueryAllEvidenceRequest{Pagination: pageReq}
}


type QueryAllEvidenceParams struct {
	Page  int `json:"page" yaml:"page"`
	Limit int `json:"limit" yaml:"limit"`
}

func NewQueryAllEvidenceParams(page, limit int) QueryAllEvidenceParams {
	return QueryAllEvidenceParams{Page: page, Limit: limit}
}
