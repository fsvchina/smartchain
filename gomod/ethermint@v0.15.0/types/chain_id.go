package types

import (
	"fmt"
	"math/big"
	"regexp"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	regexChainID         = `[a-z]{1,}`
	regexEIP155Separator = `_{1}`
	regexEIP155          = `[1-9][0-9]*`
	regexEpochSeparator  = `-{1}`
	regexEpoch           = `[1-9][0-9]*`
	ethermintChainID     = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)%s(%s)$`, regexChainID, regexEIP155Separator, regexEIP155, regexEpochSeparator, regexEpoch))
)


func IsValidChainID(chainID string) bool {
	if len(chainID) > 48 {
		return false
	}

	return ethermintChainID.MatchString(chainID)
}



func ParseChainID(chainID string) (*big.Int, error) {
	chainID = strings.TrimSpace(chainID)
	if len(chainID) > 48 {
		return nil, sdkerrors.Wrapf(ErrInvalidChainID, "chain-id '%s' cannot exceed 48 chars", chainID)
	}

	matches := ethermintChainID.FindStringSubmatch(chainID)
	if matches == nil || len(matches) != 4 || matches[1] == "" {
		return nil, sdkerrors.Wrapf(ErrInvalidChainID, "%s: %v", chainID, matches)
	}


	chainIDInt, ok := new(big.Int).SetString(matches[2], 10)
	if !ok {
		return nil, sdkerrors.Wrapf(ErrInvalidChainID, "epoch %s must be base-10 integer format", matches[2])
	}

	return chainIDInt, nil
}
