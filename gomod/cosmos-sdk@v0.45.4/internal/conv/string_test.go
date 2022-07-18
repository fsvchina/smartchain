package conv

import (
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestStringSuite(t *testing.T) {
	suite.Run(t, new(StringSuite))
}

type StringSuite struct{ suite.Suite }

func unsafeConvertStr() []byte {
	return UnsafeStrToBytes("abc")
}

func (s *StringSuite) TestUnsafeStrToBytes() {


	for i := 0; i < 5; i++ {
		b := unsafeConvertStr()
		runtime.GC()
		<-time.NewTimer(2 * time.Millisecond).C
		b2 := append(b, 'd')
		s.Equal("abc", string(b))
		s.Equal("abcd", string(b2))
	}
}

func unsafeConvertBytes() string {
	return UnsafeBytesToStr([]byte("abc"))
}

func (s *StringSuite) TestUnsafeBytesToStr() {


	for i := 0; i < 5; i++ {
		str := unsafeConvertBytes()
		runtime.GC()
		<-time.NewTimer(2 * time.Millisecond).C
		s.Equal("abc", str)
	}
}

func BenchmarkUnsafeStrToBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UnsafeStrToBytes(strconv.Itoa(i))
	}
}
