package testutil

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)


type BufferReader interface {
	io.Reader
	Reset(string)
}


type BufferWriter interface {
	io.Writer
	Reset()
	Bytes() []byte
	String() string
}



func ApplyMockIO(c *cobra.Command) (BufferReader, BufferWriter) {
	mockIn := strings.NewReader("")
	mockOut := bytes.NewBufferString("")

	c.SetIn(mockIn)
	c.SetOut(mockOut)
	c.SetErr(mockOut)

	return mockIn, mockOut
}



func ApplyMockIODiscardOutErr(c *cobra.Command) BufferReader {
	mockIn := strings.NewReader("")

	c.SetIn(mockIn)
	c.SetOut(ioutil.Discard)
	c.SetErr(ioutil.Discard)

	return mockIn
}



func WriteToNewTempFile(t testing.TB, s string) *os.File {
	t.Helper()

	fp := TempFile(t)
	_, err := fp.WriteString(s)

	require.Nil(t, err)

	return fp
}


func TempFile(t testing.TB) *os.File {
	t.Helper()

	fp, err := ioutil.TempFile(t.TempDir(), "")
	require.NoError(t, err)

	return fp
}
