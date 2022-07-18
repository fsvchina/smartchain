package keyring

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
)


type SignatureAlgo interface {
	Name() hd.PubKeyType
	Derive() hd.DeriveFn
	Generate() hd.GenerateFn
}


func NewSigningAlgoFromString(str string, algoList SigningAlgoList) (SignatureAlgo, error) {
	for _, algo := range algoList {
		if str == string(algo.Name()) {
			return algo, nil
		}
	}
	return nil, fmt.Errorf("provided algorithm %q is not supported", str)
}


type SigningAlgoList []SignatureAlgo


func (sal SigningAlgoList) Contains(algo SignatureAlgo) bool {
	for _, cAlgo := range sal {
		if cAlgo.Name() == algo.Name() {
			return true
		}
	}

	return false
}


func (sal SigningAlgoList) String() string {
	names := make([]string, len(sal))
	for i := range sal {
		names[i] = string(sal[i].Name())
	}

	return strings.Join(names, ",")
}
