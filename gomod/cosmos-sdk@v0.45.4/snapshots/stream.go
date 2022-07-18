package snapshots

import (
	"bufio"
	"compress/zlib"
	"io"

	protoio "github.com/gogo/protobuf/io"
	"github.com/gogo/protobuf/proto"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (

	snapshotChunkSize  = uint64(10e6)
	snapshotBufferSize = int(snapshotChunkSize)

	snapshotCompressionLevel = 7
)



type StreamWriter struct {
	chunkWriter *ChunkWriter
	bufWriter   *bufio.Writer
	zWriter     *zlib.Writer
	protoWriter protoio.WriteCloser
}


func NewStreamWriter(ch chan<- io.ReadCloser) *StreamWriter {
	chunkWriter := NewChunkWriter(ch, snapshotChunkSize)
	bufWriter := bufio.NewWriterSize(chunkWriter, snapshotBufferSize)
	zWriter, err := zlib.NewWriterLevel(bufWriter, snapshotCompressionLevel)
	if err != nil {
		chunkWriter.CloseWithError(sdkerrors.Wrap(err, "zlib failure"))
		return nil
	}
	protoWriter := protoio.NewDelimitedWriter(zWriter)
	return &StreamWriter{
		chunkWriter: chunkWriter,
		bufWriter:   bufWriter,
		zWriter:     zWriter,
		protoWriter: protoWriter,
	}
}


func (sw *StreamWriter) WriteMsg(msg proto.Message) error {
	return sw.protoWriter.WriteMsg(msg)
}


func (sw *StreamWriter) Close() error {
	if err := sw.protoWriter.Close(); err != nil {
		sw.chunkWriter.CloseWithError(err)
		return err
	}
	if err := sw.zWriter.Close(); err != nil {
		sw.chunkWriter.CloseWithError(err)
		return err
	}
	if err := sw.bufWriter.Flush(); err != nil {
		sw.chunkWriter.CloseWithError(err)
		return err
	}
	return sw.chunkWriter.Close()
}


func (sw *StreamWriter) CloseWithError(err error) {
	sw.chunkWriter.CloseWithError(err)
}



type StreamReader struct {
	chunkReader *ChunkReader
	zReader     io.ReadCloser
	protoReader protoio.ReadCloser
}


func NewStreamReader(chunks <-chan io.ReadCloser) (*StreamReader, error) {
	chunkReader := NewChunkReader(chunks)
	zReader, err := zlib.NewReader(chunkReader)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "zlib failure")
	}
	protoReader := protoio.NewDelimitedReader(zReader, snapshotMaxItemSize)
	return &StreamReader{
		chunkReader: chunkReader,
		zReader:     zReader,
		protoReader: protoReader,
	}, nil
}


func (sr *StreamReader) ReadMsg(msg proto.Message) error {
	return sr.protoReader.ReadMsg(msg)
}


func (sr *StreamReader) Close() error {
	sr.protoReader.Close()
	sr.zReader.Close()
	return sr.chunkReader.Close()
}
