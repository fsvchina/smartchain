package snapshots

import (
	"io"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)



type ChunkWriter struct {
	ch        chan<- io.ReadCloser
	pipe      *io.PipeWriter
	chunkSize uint64
	written   uint64
	closed    bool
}


func NewChunkWriter(ch chan<- io.ReadCloser, chunkSize uint64) *ChunkWriter {
	return &ChunkWriter{
		ch:        ch,
		chunkSize: chunkSize,
	}
}


func (w *ChunkWriter) chunk() error {
	if w.pipe != nil {
		err := w.pipe.Close()
		if err != nil {
			return err
		}
	}
	pr, pw := io.Pipe()
	w.ch <- pr
	w.pipe = pw
	w.written = 0
	return nil
}


func (w *ChunkWriter) Close() error {
	if !w.closed {
		w.closed = true
		close(w.ch)
		var err error
		if w.pipe != nil {
			err = w.pipe.Close()
		}
		return err
	}
	return nil
}


func (w *ChunkWriter) CloseWithError(err error) {
	if !w.closed {
		w.closed = true
		close(w.ch)
		if w.pipe != nil {
			w.pipe.CloseWithError(err)
		}
	}
}


func (w *ChunkWriter) Write(data []byte) (int, error) {
	if w.closed {
		return 0, sdkerrors.Wrap(sdkerrors.ErrLogic, "cannot write to closed ChunkWriter")
	}
	nTotal := 0
	for len(data) > 0 {
		if w.pipe == nil || (w.written >= w.chunkSize && w.chunkSize > 0) {
			err := w.chunk()
			if err != nil {
				return nTotal, err
			}
		}

		var writeSize uint64
		if w.chunkSize == 0 {
			writeSize = uint64(len(data))
		} else {
			writeSize = w.chunkSize - w.written
		}
		if writeSize > uint64(len(data)) {
			writeSize = uint64(len(data))
		}

		n, err := w.pipe.Write(data[:writeSize])
		w.written += uint64(n)
		nTotal += n
		if err != nil {
			return nTotal, err
		}
		data = data[writeSize:]
	}
	return nTotal, nil
}


type ChunkReader struct {
	ch     <-chan io.ReadCloser
	reader io.ReadCloser
}


func NewChunkReader(ch <-chan io.ReadCloser) *ChunkReader {
	return &ChunkReader{ch: ch}
}


func (r *ChunkReader) next() error {
	reader, ok := <-r.ch
	if !ok {
		return io.EOF
	}
	r.reader = reader
	return nil
}


func (r *ChunkReader) Close() error {
	var err error
	if r.reader != nil {
		err = r.reader.Close()
		r.reader = nil
	}
	for reader := range r.ch {
		if e := reader.Close(); e != nil && err == nil {
			err = e
		}
	}
	return err
}


func (r *ChunkReader) Read(p []byte) (int, error) {
	if r.reader == nil {
		err := r.next()
		if err != nil {
			return 0, err
		}
	}
	n, err := r.reader.Read(p)
	if err == io.EOF {
		err = r.reader.Close()
		r.reader = nil
		if err != nil {
			return 0, err
		}
		return r.Read(p)
	}
	return n, err
}


func DrainChunks(chunks <-chan io.ReadCloser) {
	for chunk := range chunks {
		_ = chunk.Close()
	}
}
