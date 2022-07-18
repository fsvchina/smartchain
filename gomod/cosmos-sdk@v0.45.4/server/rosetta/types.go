package rosetta

import (
	"crypto/sha256"
)


const (
	StatusTxSuccess   = "Success"
	StatusTxReverted  = "Reverted"
	StatusPeerSynced  = "synced"
	StatusPeerSyncing = "syncing"
)





const (
	DeliverTxSize       = sha256.Size
	BeginEndBlockTxSize = DeliverTxSize + 1
	EndBlockHashStart   = 0x0
	BeginBlockHashStart = 0x1
)

const (




	BurnerAddressIdentifier = "burner"
)



type TransactionType int

const (
	UnrecognizedTx TransactionType = iota
	BeginBlockTx
	EndBlockTx
	DeliverTxTx
)




const (
	Log = "log"
)



type ConstructionPreprocessMetadata struct {
	Memo     string `json:"memo"`
	GasLimit uint64 `json:"gas_limit"`
	GasPrice string `json:"gas_price"`
}

func (c *ConstructionPreprocessMetadata) FromMetadata(meta map[string]interface{}) error {
	return unmarshalMetadata(meta, c)
}


type PreprocessOperationsOptionsResponse struct {
	ExpectedSigners []string `json:"expected_signers"`
	Memo            string   `json:"memo"`
	GasLimit        uint64   `json:"gas_limit"`
	GasPrice        string   `json:"gas_price"`
}

func (c PreprocessOperationsOptionsResponse) ToMetadata() (map[string]interface{}, error) {
	return marshalMetadata(c)
}

func (c *PreprocessOperationsOptionsResponse) FromMetadata(meta map[string]interface{}) error {
	return unmarshalMetadata(meta, c)
}



type SignerData struct {
	AccountNumber uint64 `json:"account_number"`
	Sequence      uint64 `json:"sequence"`
}




type ConstructionMetadata struct {
	ChainID     string        `json:"chain_id"`
	SignersData []*SignerData `json:"signer_data"`
	GasLimit    uint64        `json:"gas_limit"`
	GasPrice    string        `json:"gas_price"`
	Memo        string        `json:"memo"`
}

func (c ConstructionMetadata) ToMetadata() (map[string]interface{}, error) {
	return marshalMetadata(c)
}

func (c *ConstructionMetadata) FromMetadata(meta map[string]interface{}) error {
	return unmarshalMetadata(meta, c)
}
