package simulation

import (
	"fmt"
	"os"
	"path"
	"time"
)


type LogWriter interface {
	AddEntry(OperationEntry)
	PrintLogs()
}


func NewLogWriter(testingmode bool) LogWriter {
	if !testingmode {
		return &DummyLogWriter{}
	}

	return &StandardLogWriter{}
}


type StandardLogWriter struct {
	OpEntries []OperationEntry `json:"op_entries" yaml:"op_entries"`
}


func (lw *StandardLogWriter) AddEntry(opEntry OperationEntry) {
	lw.OpEntries = append(lw.OpEntries, opEntry)
}


func (lw *StandardLogWriter) PrintLogs() {
	f := createLogFile()
	defer f.Close()

	for i := 0; i < len(lw.OpEntries); i++ {
		writeEntry := fmt.Sprintf("%s\n", (lw.OpEntries[i]).MustMarshal())
		_, err := f.WriteString(writeEntry)

		if err != nil {
			panic("Failed to write logs to file")
		}
	}
}

func createLogFile() *os.File {
	var f *os.File

	fileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02_15:04:05"))
	folderPath := path.Join(os.ExpandEnv("$HOME"), ".simapp", "simulations")
	filePath := path.Join(folderPath, fileName)

	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, err = os.Create(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Logs to writing to %s\n", filePath)

	return f
}


type DummyLogWriter struct{}


func (lw *DummyLogWriter) AddEntry(_ OperationEntry) {}


func (lw *DummyLogWriter) PrintLogs() {}
