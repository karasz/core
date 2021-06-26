package envexp

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
	exp *Expander
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
func (exp *Expander) NewReaderSize(in io.Reader, size int) *Reader {
	if size < MinimumBufferSize {
		size = DefaultBufferSize
	}

	return &Reader{
		exp: exp,
		in:  in,
		buf: bytes.NewBuffer(make([]byte, 0, size)),
		out: bytes.NewBuffer(make([]byte, 0, size)),
	}
}

func (exp *Expander) NewReader(in io.Reader) *Reader {
	return exp.NewReaderSize(in, DefaultBufferSize)
}

func (exp *Expander) NewReaderFileSize(filename string, size int) (*Reader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	rd := exp.NewReaderSize(f, size)
	return rd, nil
}

func (exp *Expander) NewReaderFile(filename string) (*Reader, error) {
	return exp.NewReaderFileSize(filename, DefaultBufferSize)
}
