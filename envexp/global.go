package envexp

import (
	"io"
	"os"
)

var envExpander = &Expander{
	get: os.Getenv,
}

// Expand uses os.GetEnv() to expand a given string
func Expand(s string) string {
	return envExpander.Expand(s)
}

// NewReaderSize creates a new Reader wrapper with a given buffer size
// using os.GetEnv() for expanding
func NewReaderSize(in io.Reader, size int) *Reader {
	return envExpander.NewReaderSize(in, size)
}

// NewReader creates a new Reader wrapper
// using os.GetEnv() for expanding
func NewReader(in io.Reader) *Reader {
	return envExpander.NewReader(in)
}

// NewReaderFileSize creates a new Reader for a file
// with a given buffer size using os.GetEnv() for expanding
func NewReaderFileSize(filename string, size int) (*Reader, error) {
	return envExpander.NewReaderFileSize(filename, size)
}

// NewReaderFile creates a new Reader for a file
// using os.GetEnv() for expanding
func NewReaderFile(filename string) (*Reader, error) {
	return envExpander.NewReaderFile(filename)
}
