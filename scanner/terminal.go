package scanner

import (
	"unicode/utf8"
)

// A literal element within a document
type Terminal struct {
	val          string
	bytes, runes uint
	line, col    uint
}

// Creates a new Terminal
func NewTerminalFull(val string, bytes, runes, line, col uint) *Terminal {
	return &Terminal{
		val:   val,
		bytes: bytes,
		runes: runes,
		line:  line,
		col:   col,
	}
}

// Creates a new Terminal without knowing it's length
func NewTerminal(val string, line, col uint) *Terminal {
	bytes := uint(len(val))
	runes := uint(utf8.RuneCountInString(val))

	return NewTerminalFull(val, bytes, runes, line, col)
}

// Position of the terminal in the document
func (t *Terminal) Position() (uint, uint) {
	return t.line, t.col
}

// Value and size of the Terminal
func (t *Terminal) Value() (string, uint, uint) {
	return t.val, t.bytes, t.runes
}
