package v039

import "encoding/json"

const (
	ModuleName = "genutil"
)


type GenesisState struct {
	GenTxs []json.RawMessage `json:"gentxs" yaml:"gentxs"`
}
