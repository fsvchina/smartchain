package simulation


type Config struct {
	GenesisFile string
	ParamsFile  string

	ExportParamsPath   string
	ExportParamsHeight int
	ExportStatePath    string
	ExportStatsPath    string

	Seed               int64
	InitialBlockHeight int
	NumBlocks          int
	BlockSize          int
	ChainID            string

	Lean   bool
	Commit bool

	OnOperation   bool
	AllInvariants bool
}
