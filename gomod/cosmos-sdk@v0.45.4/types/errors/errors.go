package errors

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)


const RootCodespace = "sdk"


const UndefinedCodespace = "undefined"

var (

	errInternal = Register(UndefinedCodespace, 1, "internal")


	ErrTxDecode = Register(RootCodespace, 2, "tx parse error")



	ErrInvalidSequence = Register(RootCodespace, 3, "invalid sequence")



	ErrUnauthorized = Register(RootCodespace, 4, "unauthorized")


	ErrInsufficientFunds = Register(RootCodespace, 5, "insufficient funds")


	ErrUnknownRequest = Register(RootCodespace, 6, "unknown request")


	ErrInvalidAddress = Register(RootCodespace, 7, "invalid address")


	ErrInvalidPubKey = Register(RootCodespace, 8, "invalid pubkey")


	ErrUnknownAddress = Register(RootCodespace, 9, "unknown address")


	ErrInvalidCoins = Register(RootCodespace, 10, "invalid coins")


	ErrOutOfGas = Register(RootCodespace, 11, "out of gas")


	ErrMemoTooLarge = Register(RootCodespace, 12, "memo too large")


	ErrInsufficientFee = Register(RootCodespace, 13, "insufficient fee")


	ErrTooManySignatures = Register(RootCodespace, 14, "maximum number of signatures exceeded")


	ErrNoSignatures = Register(RootCodespace, 15, "no signatures supplied")


	ErrJSONMarshal = Register(RootCodespace, 16, "failed to marshal JSON bytes")


	ErrJSONUnmarshal = Register(RootCodespace, 17, "failed to unmarshal JSON bytes")



	ErrInvalidRequest = Register(RootCodespace, 18, "invalid request")



	ErrTxInMempoolCache = Register(RootCodespace, 19, "tx already in mempool")


	ErrMempoolIsFull = Register(RootCodespace, 20, "mempool is full")


	ErrTxTooLarge = Register(RootCodespace, 21, "tx too large")


	ErrKeyNotFound = Register(RootCodespace, 22, "key not found")


	ErrWrongPassword = Register(RootCodespace, 23, "invalid account password")


	ErrorInvalidSigner = Register(RootCodespace, 24, "tx intended signer does not match the given signer")


	ErrorInvalidGasAdjustment = Register(RootCodespace, 25, "invalid gas adjustment")


	ErrInvalidHeight = Register(RootCodespace, 26, "invalid height")


	ErrInvalidVersion = Register(RootCodespace, 27, "invalid version")


	ErrInvalidChainID = Register(RootCodespace, 28, "invalid chain-id")


	ErrInvalidType = Register(RootCodespace, 29, "invalid type")



	ErrTxTimeoutHeight = Register(RootCodespace, 30, "tx timeout height")


	ErrUnknownExtensionOptions = Register(RootCodespace, 31, "unknown extension options")



	ErrWrongSequence = Register(RootCodespace, 32, "incorrect account sequence")


	ErrPackAny = Register(RootCodespace, 33, "failed packing protobuf message to Any")


	ErrUnpackAny = Register(RootCodespace, 34, "failed unpacking protobuf message from Any")



	ErrLogic = Register(RootCodespace, 35, "internal logic error")



	ErrConflict = Register(RootCodespace, 36, "conflict")



	ErrNotSupported = Register(RootCodespace, 37, "feature not supported")


	ErrNotFound = Register(RootCodespace, 38, "not found")



	ErrIO = Register(RootCodespace, 39, "Internal IO error")


	ErrAppConfig = Register(RootCodespace, 40, "error in app.toml")



	ErrPanic = Register(UndefinedCodespace, 111222, "panic")
)



//



//

func Register(codespace string, code uint32, description string) *Error {

	if e := getUsed(codespace, code); e != nil {
		panic(fmt.Sprintf("error with code %d is already registered: %q", code, e.desc))
	}

	err := New(codespace, code, description)
	setUsed(err)

	return err
}



var usedCodes = map[string]*Error{}

func errorID(codespace string, code uint32) string {
	return fmt.Sprintf("%s:%d", codespace, code)
}

func getUsed(codespace string, code uint32) *Error {
	return usedCodes[errorID(codespace, code)]
}

func setUsed(err *Error) {
	usedCodes[errorID(err.codespace, err.code)] = err
}





//


func ABCIError(codespace string, code uint32, log string) error {
	if e := getUsed(codespace, code); e != nil {
		return Wrap(e, log)
	}


	return Wrap(New(codespace, code, "unknown"), log)
}


//



//



type Error struct {
	codespace string
	code      uint32
	desc      string
}

func New(codespace string, code uint32, desc string) *Error {
	return &Error{codespace: codespace, code: code, desc: desc}
}

func (e Error) Error() string {
	return e.desc
}

func (e Error) ABCICode() uint32 {
	return e.code
}

func (e Error) Codespace() string {
	return e.codespace
}



func (e *Error) Is(err error) bool {


	if e == nil {
		return isNilErr(err)
	}

	for {
		if err == e {
			return true
		}



		if u, ok := err.(unpacker); ok {
			for _, er := range u.Unpack() {
				if e.Is(er) {
					return true
				}
			}
		}

		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return false
		}
	}
}



func (e *Error) Wrap(desc string) error { return Wrap(e, desc) }



func (e *Error) Wrapf(desc string, args ...interface{}) error { return Wrapf(e, desc, args...) }

func isNilErr(err error) bool {


	if err == nil {
		return true
	}
	if reflect.ValueOf(err).Kind() == reflect.Struct {
		return false
	}
	return reflect.ValueOf(err).IsNil()
}


//


//


func Wrap(err error, description string) error {
	if err == nil {
		return nil
	}




	if stackTrace(err) == nil {
		err = errors.WithStack(err)
	}

	return &wrappedError{
		parent: err,
		msg:    description,
	}
}


//


func Wrapf(err error, format string, args ...interface{}) error {
	desc := fmt.Sprintf(format, args...)
	return Wrap(err, desc)
}

type wrappedError struct {

	msg string

	parent error
}

func (e *wrappedError) Error() string {
	return fmt.Sprintf("%s: %s", e.msg, e.parent.Error())
}

func (e *wrappedError) Cause() error {
	return e.parent
}


func (e *wrappedError) Is(target error) bool {
	if e == target {
		return true
	}

	w := e.Cause()
	for {
		if w == target {
			return true
		}

		x, ok := w.(causer)
		if ok {
			w = x.Cause()
		}
		if x == nil {
			return false
		}
	}
}


func (e *wrappedError) Unwrap() error {
	return e.parent
}




func Recover(err *error) {
	if r := recover(); r != nil {
		*err = Wrapf(ErrPanic, "%v", r)
	}
}


func WithType(err error, obj interface{}) error {
	return Wrap(err, fmt.Sprintf("%T", obj))
}



func IsOf(received error, targets ...error) bool {
	for _, t := range targets {
		if errors.Is(received, t) {
			return true
		}
	}
	return false
}



type causer interface {
	Cause() error
}

type unpacker interface {
	Unpack() []error
}
