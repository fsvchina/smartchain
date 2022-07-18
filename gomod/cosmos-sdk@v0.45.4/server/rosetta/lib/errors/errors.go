package errors




import (
	"fmt"

	grpccodes "google.golang.org/grpc/codes"
	grpcstatus "google.golang.org/grpc/status"

	"github.com/coinbase/rosetta-sdk-go/types"
)


func ListErrors() []*types.Error {
	return registry.list()
}


func SealAndListErrors() []*types.Error {
	registry.seal()
	return registry.list()
}


type Error struct {
	rosErr *types.Error
}

func (e *Error) Error() string {
	if e.rosErr == nil {
		return ErrUnknown.Error()
	}
	return fmt.Sprintf("rosetta: (%d) %s", e.rosErr.Code, e.rosErr.Message)
}



func (e *Error) Is(err error) bool {

	rosErr, ok := err.(*Error)
	if rosErr == nil || !ok {
		return false
	}

	if rosErr.rosErr == nil || e.rosErr == nil {
		return false
	}

	return rosErr.rosErr.Code == e.rosErr.Code
}


func WrapError(err *Error, msg string) *Error {
	return &Error{rosErr: &types.Error{
		Code:        err.rosErr.Code,
		Message:     err.rosErr.Message,
		Description: err.rosErr.Description,
		Retriable:   err.rosErr.Retriable,
		Details: map[string]interface{}{
			"info": msg,
		},
	}}
}



func ToRosetta(err error) *types.Error {

	rosErr, ok := err.(*Error)
	if rosErr == nil || !ok {
		return ToRosetta(WrapError(ErrUnknown, ErrUnknown.Error()))
	}
	return rosErr.rosErr
}


func FromGRPCToRosettaError(err error) *Error {
	status, ok := grpcstatus.FromError(err)
	if !ok {
		return WrapError(ErrUnknown, err.Error())
	}
	switch status.Code() {
	case grpccodes.NotFound:
		return WrapError(ErrNotFound, status.Message())
	case grpccodes.FailedPrecondition:
		return WrapError(ErrBadArgument, status.Message())
	case grpccodes.InvalidArgument:
		return WrapError(ErrBadArgument, status.Message())
	case grpccodes.Internal:
		return WrapError(ErrInternal, status.Message())
	default:
		return WrapError(ErrUnknown, status.Message())
	}
}

func RegisterError(code int32, message string, retryable bool, description string) *Error {
	e := &Error{rosErr: &types.Error{
		Code:        code,
		Message:     message,
		Description: &description,
		Retriable:   retryable,
		Details:     nil,
	}}
	registry.add(e)
	return e
}


var (


	ErrUnknown = RegisterError(0, "unknown", false, "unknown error")

	ErrOffline = RegisterError(1, "cannot query endpoint in offline mode", false, "returned when querying an online endpoint in offline mode")

	ErrNetworkNotSupported = RegisterError(2, "network is not supported", false, "returned when querying a non supported network")

	ErrCodec = RegisterError(3, "encode/decode error", true, "returned when there are errors encoding or decoding information to and from the node")

	ErrInvalidOperation = RegisterError(4, "invalid operation", false, "returned when the operation is not valid")

	ErrInvalidTransaction = RegisterError(5, "invalid transaction", false, "returned when the transaction is invalid")

	ErrInvalidAddress = RegisterError(7, "invalid address", false, "returned when the address is malformed")

	ErrInvalidPubkey = RegisterError(8, "invalid pubkey", false, "returned when the public key is invalid")

	ErrInterpreting = RegisterError(9, "error interpreting data from node", false, "returned when there are issues interpreting requests or response from node")
	ErrInvalidMemo  = RegisterError(11, "invalid memo", false, "returned when the memo is invalid")

	ErrBadArgument = RegisterError(400, "bad argument", false, "request is malformed")



	ErrNotFound = RegisterError(404, "not found", true, "returned when the node does not find what the client is asking for")

	ErrInternal = RegisterError(500, "internal error", false, "returned when the node experiences internal errors")

	ErrBadGateway = RegisterError(502, "bad gateway", true, "return when the node is unreachable")

	ErrNotImplemented = RegisterError(14, "not implemented", false, "returned when querying an endpoint which is not implemented")

	ErrUnsupportedCurve = RegisterError(15, "unsupported curve, expected secp256k1", false, "returned when using an unsupported crypto curve")
)
