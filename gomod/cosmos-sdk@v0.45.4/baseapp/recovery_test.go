package baseapp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestRecoveryChain(t *testing.T) {
	createError := func(id int) error {
		return fmt.Errorf("error from id: %d", id)
	}

	createHandler := func(id int, handle bool) RecoveryHandler {
		return func(_ interface{}) error {
			if handle {
				return createError(id)
			}
			return nil
		}
	}


	{
		mw := newRecoveryMiddleware(createHandler(3, false), nil)
		mw = newRecoveryMiddleware(createHandler(2, false), mw)
		mw = newRecoveryMiddleware(createHandler(1, true), mw)
		receivedErr := processRecovery(nil, mw)

		require.Equal(t, createError(1), receivedErr)
	}


	{
		mw := newRecoveryMiddleware(createHandler(3, false), nil)
		mw = newRecoveryMiddleware(createHandler(2, true), mw)
		mw = newRecoveryMiddleware(createHandler(1, false), mw)
		receivedErr := processRecovery(nil, mw)

		require.Equal(t, createError(2), receivedErr)
	}


	{
		mw := newRecoveryMiddleware(createHandler(3, true), nil)
		mw = newRecoveryMiddleware(createHandler(2, false), mw)
		mw = newRecoveryMiddleware(createHandler(1, false), mw)
		receivedErr := processRecovery(nil, mw)

		require.Equal(t, createError(3), receivedErr)
	}


	{
		mw := newRecoveryMiddleware(createHandler(3, false), nil)
		mw = newRecoveryMiddleware(createHandler(2, false), mw)
		mw = newRecoveryMiddleware(createHandler(1, false), mw)
		receivedErr := processRecovery(nil, mw)

		require.Nil(t, receivedErr)
	}
}
