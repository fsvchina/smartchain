package types


const (
	QueryParams = "params"
)



type QuerySubspaceParams struct {
	Subspace string
	Key      string
}


type SubspaceParamsResponse struct {
	Subspace string
	Key      string
	Value    string
}

func NewQuerySubspaceParams(ss, key string) QuerySubspaceParams {
	return QuerySubspaceParams{
		Subspace: ss,
		Key:      key,
	}
}

func NewSubspaceParamsResponse(ss, key, value string) SubspaceParamsResponse {
	return SubspaceParamsResponse{
		Subspace: ss,
		Key:      key,
		Value:    value,
	}
}
