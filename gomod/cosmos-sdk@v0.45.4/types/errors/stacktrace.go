package errors

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

func matchesFunc(f errors.Frame, prefixes ...string) bool {
	fn := funcName(f)
	for _, prefix := range prefixes {
		if strings.HasPrefix(fn, prefix) {
			return true
		}
	}
	return false
}


func funcName(f errors.Frame) string {


	pc := uintptr(f) - 1
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}

func fileLine(f errors.Frame) (string, int) {



	pc := uintptr(f) - 1
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "unknown", 0
	}
	return fn.FileLine(pc)
}

func trimInternal(st errors.StackTrace) errors.StackTrace {


	for matchesFunc(st[0],

		"github.com/cosmos/cosmos-sdk/types/errors.Wrap",
		"github.com/cosmos/cosmos-sdk/types/errors.Wrapf",
		"github.com/cosmos/cosmos-sdk/types/errors.WithType",

		"runtime.",


	) {
		st = st[1:]
	}

	for l := len(st) - 1; l > 0 && matchesFunc(st[l], "runtime.", "testing."); l-- {
		st = st[:l]
	}
	return st
}

func writeSimpleFrame(s io.Writer, f errors.Frame) {
	file, line := fileLine(f)


	chunks := strings.SplitN(file, "github.com/", 2)
	if len(chunks) == 2 {
		file = chunks[1]
	}
	fmt.Fprintf(s, " [%s:%d]", file, line)
}






//

func (e *wrappedError) Format(s fmt.State, verb rune) {

	if verb != 'v' {
		fmt.Fprint(s, e.Error())
		return
	}

	stack := trimInternal(stackTrace(e))
	if s.Flag('+') {
		fmt.Fprintf(s, "%+v\n", stack)
		fmt.Fprint(s, e.Error())
	} else {
		fmt.Fprint(s, e.Error())
		writeSimpleFrame(s, stack[0])
	}
}



func stackTrace(err error) errors.StackTrace {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	for {
		if st, ok := err.(stackTracer); ok {
			return st.StackTrace()
		}

		if c, ok := err.(causer); ok {
			err = c.Cause()
		} else {
			return nil
		}
	}
}
