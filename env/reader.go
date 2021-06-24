package env

import (
	"bytes"
	"io"
	"os"
)

const (
	MinimumBufferSize = 32
	DefaultBufferSize = 4096
)

type Reader struct {
	in  io.Reader
	buf *bytes.Buffer
	out *bytes.Buffer
}

func (rd *Reader) Read(b []byte) (int, error) {
	return rd.in.Read(b)
}

func (rd *Reader) Close() error {
	if f, ok := rd.in.(io.Closer); ok {
		return f.Close()
	}
	return nil
}

//
// Constructors
//
func NewReaderSize(in io.Reader, size int) *Reader {
	if size < MinimumBufferSize {
		size = DefaultBufferSize
	}

	return &Reader{
		in:  in,
		buf: bytes.NewBuffer(make([]byte, 0, size)),
		out: bytes.NewBuffer(make([]byte, 0, size)),
	}
}

func NewReader(in io.Reader) *Reader {
	return NewReaderSize(in, DefaultBufferSize)
}

func NewReaderFileSize(filename string, size int) (*Reader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	rd := NewReaderSize(f, size)
	return rd, nil
}

func NewReaderFile(filename string) (*Reader, error) {
	return NewReaderFileSize(filename, DefaultBufferSize)
}
