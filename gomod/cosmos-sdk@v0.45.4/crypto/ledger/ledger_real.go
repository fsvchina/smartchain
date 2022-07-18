

package ledger

import ledger "github.com/cosmos/ledger-cosmos-go"




func init() {
	discoverLedger = func() (SECP256K1, error) {
		device, err := ledger.FindLedgerCosmosUserApp()
		if err != nil {
			return nil, err
		}

		return device, nil
	}
}
