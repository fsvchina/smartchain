package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterError(t *testing.T) {
	var error *Error

	registeredErrorsCount := 16
	assert.Equal(t, len(registry.list()), registeredErrorsCount)
	assert.ElementsMatch(t, registry.list(), ListErrors())

	error = RegisterError(69, "nice!", false, "nice!")
	assert.NotNil(t, error)

	registeredErrorsCount++
	assert.Equal(t, len(ListErrors()), registeredErrorsCount)

	error = RegisterError(69, "nice!", false, "nice!")
	assert.Equal(t, len(ListErrors()), registeredErrorsCount)


	assert.Equal(t, registry.sealed, false)
	errors := SealAndListErrors()
	assert.Equal(t, registry.sealed, true)
	assert.Equal(t, len(errors), registeredErrorsCount)

	error = RegisterError(1024, "bytes", false, "bytes")
	assert.NotNil(t, error)

}

func TestError_Error(t *testing.T) {
	var error *Error

	assert.False(t, ErrOffline.Is(error))
	error = &Error{}
	assert.False(t, ErrOffline.Is(error))

	assert.False(t, ErrOffline.Is(&MyError{}))

	error = WrapError(ErrOffline, "offline")
	assert.True(t, ErrOffline.Is(error))


	assert.False(t, ErrOffline.Is(ErrBadGateway))
	assert.True(t, ErrBadGateway.Is(ErrBadGateway))
}

func TestToRosetta(t *testing.T) {
	var error *Error

	assert.NotNil(t, ToRosetta(error))

	assert.NotNil(t, ToRosetta(&MyError{}))
}

type MyError struct {
}

func (e *MyError) Error() string {
	return ""
}
func (e *MyError) Is(err error) bool {
	return true
}
