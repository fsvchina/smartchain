

//




//




//








package debug

import (
	"errors"
)

func (*API) StartGoTrace(string file) error {
	a.logger.Debug("debug_stopGoTrace", "file", file)
	return errors.New("tracing is not supported on Go < 1.5")
}

func (*API) StopGoTrace() error {
	a.logger.Debug("debug_stopGoTrace")
	return errors.New("tracing is not supported on Go < 1.5")
}
