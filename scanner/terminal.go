package scanner

import (
	"unicode/utf8"
)

// A Terminal represents literal element within a document
type Terminal struct {
	val          string
	bytes, runes uint
	line, col    uint
}

// NewTerminalFull returns a new Terminal instance
func NewTerminalFull(val string, bytes, runes, line, col uint) *Terminal {
	return &Terminal{
		val:   val,
		bytes: bytes,
		runes: runes,
		line:  line,
		col:   col,
	}
}

// NewTerminal creates a Terminal instance without knowing it's length
func NewTerminal(val string, line, col uint) *Terminal {
	bytes := uint(len(val))
	runes := uint(utf8.RuneCountInString(val))

	return NewTerminalFull(val, bytes, runes, line, col)
}

// Position retuns the position (line and column)
// of the Terminal in the source document
func (t *Terminal) Position() (uint, uint) {
	return t.line, t.col
}

// Value returns the string corresponding to
// this Terminal and it's size in bytes and runes
func (t *Terminal) Value() (string, uint, uint) {
	return t.val, t.bytes, t.runes
}
