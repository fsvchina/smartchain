package types


var MinterKey = []byte{0x00}

const (

	ModuleName = "mint"


	StoreKey = ModuleName


	QuerierRoute = StoreKey


	QueryParameters       = "parameters"
	QueryInflation        = "inflation"
	QueryAnnualProvisions = "annual_provisions"
)
