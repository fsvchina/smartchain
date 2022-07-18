package types

import (
	ethaccounts "github.com/ethereum/go-ethereum/accounts"
)

var (

	Bip44CoinType uint32 = 60


	BIP44HDPath = ethaccounts.DefaultBaseDerivationPath.String()
)

type (
	HDPathIterator func() ethaccounts.DerivationPath
)



func NewHDPathIterator(basePath string, ledgerIter bool) (HDPathIterator, error) {
	hdPath, err := ethaccounts.ParseDerivationPath(basePath)
	if err != nil {
		return nil, err
	}

	if ledgerIter {
		return ethaccounts.LedgerLiveIterator(hdPath), nil
	}

	return ethaccounts.DefaultIterator(hdPath), nil
}
