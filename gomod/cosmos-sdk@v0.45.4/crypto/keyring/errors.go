package keyring

import "github.com/pkg/errors"

var (
	
	
	ErrUnsupportedSigningAlgo = errors.New("unsupported signing algo")

	
	
	ErrUnsupportedLanguage = errors.New("unsupported language: only english is supported")
)
