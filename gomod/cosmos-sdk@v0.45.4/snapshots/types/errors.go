package types

import (
	"errors"
)

var (

	ErrUnknownFormat = errors.New("unknown snapshot format")


	ErrChunkHashMismatch = errors.New("chunk hash verification failed")


	ErrInvalidMetadata = errors.New("invalid snapshot metadata")
)
