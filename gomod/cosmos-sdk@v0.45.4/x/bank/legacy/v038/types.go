


package v038



const (
	ModuleName = "bank"
)

type (
	GenesisState struct {
		SendEnabled bool `json:"send_enabled" yaml:"send_enabled"`
	}
)
