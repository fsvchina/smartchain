package errors

import (
	"fmt"
	"reflect"

	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	
	
	SuccessABCICode = 0

	
	
	
	internalABCICodespace        = UndefinedCodespace
	internalABCICode      uint32 = 1
)








func ABCIInfo(err error, debug bool) (codespace string, code uint32, log string) {
	if errIsNil(err) {
		return "", SuccessABCICode, ""
	}

	encode := defaultErrEncoder
	if debug {
		encode = debugErrEncoder
	}

	return abciCodespace(err), abciCode(err), encode(err)
}



func ResponseCheckTx(err error, gw, gu uint64, debug bool) abci.ResponseCheckTx {
	space, code, log := ABCIInfo(err, debug)
	return abci.ResponseCheckTx{
		Codespace: space,
		Code:      code,
		Log:       log,
		GasWanted: int64(gw),
		GasUsed:   int64(gu),
	}
}



func ResponseCheckTxWithEvents(err error, gw, gu uint64, events []abci.Event, debug bool) abci.ResponseCheckTx {
	space, code, log := ABCIInfo(err, debug)
	return abci.ResponseCheckTx{
		Codespace: space,
		Code:      code,
		Log:       log,
		GasWanted: int64(gw),
		GasUsed:   int64(gu),
		Events:    events,
	}
}



func ResponseDeliverTx(err error, gw, gu uint64, debug bool) abci.ResponseDeliverTx {
	space, code, log := ABCIInfo(err, debug)
	return abci.ResponseDeliverTx{
		Codespace: space,
		Code:      code,
		Log:       log,
		GasWanted: int64(gw),
		GasUsed:   int64(gu),
	}
}



func ResponseDeliverTxWithEvents(err error, gw, gu uint64, events []abci.Event, debug bool) abci.ResponseDeliverTx {
	space, code, log := ABCIInfo(err, debug)
	return abci.ResponseDeliverTx{
		Codespace: space,
		Code:      code,
		Log:       log,
		GasWanted: int64(gw),
		GasUsed:   int64(gu),
		Events:    events,
	}
}



func QueryResult(err error) abci.ResponseQuery {
	space, code, log := ABCIInfo(err, false)
	return abci.ResponseQuery{
		Codespace: space,
		Code:      code,
		Log:       log,
	}
}




func QueryResultWithDebug(err error, debug bool) abci.ResponseQuery {
	space, code, log := ABCIInfo(err, debug)
	return abci.ResponseQuery{
		Codespace: space,
		Code:      code,
		Log:       log,
	}
}


func debugErrEncoder(err error) string {
	return fmt.Sprintf("%+v", err)
}


func defaultErrEncoder(err error) string {
	return Redact(err).Error()
}

type coder interface {
	ABCICode() uint32
}




func abciCode(err error) uint32 {
	if errIsNil(err) {
		return SuccessABCICode
	}

	for {
		if c, ok := err.(coder); ok {
			return c.ABCICode()
		}

		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return internalABCICode
		}
	}
}

type codespacer interface {
	Codespace() string
}




func abciCodespace(err error) string {
	if errIsNil(err) {
		return ""
	}

	for {
		if c, ok := err.(codespacer); ok {
			return c.Codespace()
		}

		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return internalABCICodespace
		}
	}
}


//



func errIsNil(err error) bool {
	if err == nil {
		return true
	}
	if val := reflect.ValueOf(err); val.Kind() == reflect.Ptr {
		return val.IsNil()
	}
	return false
}

var errPanicWithMsg = Wrapf(ErrPanic, "panic message redacted to hide potentially sensitive system info")




func Redact(err error) error {
	if ErrPanic.Is(err) {
		return errPanicWithMsg
	}
	if abciCode(err) == internalABCICode {
		return errInternal
	}

	return err
}
