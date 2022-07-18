




package ledger

import (
	"github.com/pkg/errors"
)




func init() {
	discoverLedger = func() (SECP256K1, error) {
		return nil, errors.New("support for ledger devices is not available in this executable")
	}
}
