package scanner

import (
	"unicode/utf8"
)

const (
	// EOF is a dummy rune representing End-Of-File
	EOF = -1
)

// A Position in the input string and in the line-based document
type Position struct {
	Offset       uint
	Line, Column uint
}

// An Scanner represent the low level layer for text parsers
type Scanner struct {
	name  string
	input string

	base   Position
	cursor Position
	runes  uint
}

// NewScannerFromString instantiates a new Scanner to
// parse a given string
func NewScannerFromString(name, input string) *Scanner {
	return &Scanner{
		name:   name,
		input:  input,
		base:   Position{0, 1, 1},
		cursor: Position{0, 1, 1},
		runes:  0,
	}
}

// Length returns the number of bytes and runes in the Terminal that is been detected
func (l *Scanner) Length() (uint, uint) {
	return l.cursor.Offset - l.base.Offset, l.runes
}

// Empty tells if there are no runes accounted for the next Terminal yet
func (l *Scanner) Empty() bool {
	return l.runes == 0
}

// StepForth moves the cursor forward
func (l *Scanner) StepForth(runes, bytes uint) {
	l.cursor.Offset += bytes
	l.cursor.Column += runes
	l.runes += runes
}

// StepBack moves the cursor backward
func (l *Scanner) StepBack(runes, bytes uint) {
	l.cursor.Offset -= bytes
	// FIXME: what if column goes < 1?
	l.cursor.Column -= runes
	l.runes -= runes
}

// Reset moves the cursor back to the base
func (l *Scanner) Reset() {
	l.cursor = l.base
	l.runes = 0
}

// Skip trashes everything up to the cursor
func (l *Scanner) Skip() {
	l.base = l.cursor
	l.runes = 0
}

// NewLine accounts a line break in the position of the cursor
func (l *Scanner) NewLine() {
	l.cursor.Line++
	l.cursor.Column = 1
}

// Peek returns the next rune but not moving the cursor
func (l *Scanner) Peek() (rune, uint) {
	if l.cursor.Offset == uint(len(l.input)) {
		return EOF, 0
	}
	r, bytes := utf8.DecodeRuneInString(l.input[l.cursor.Offset:])
	return r, uint(bytes)
}

// Next returns the next rune but moving the cursor
func (l *Scanner) Next() (rune, uint) {
	r, bytes := l.Peek()
	if bytes > 0 {
		l.StepForth(1, bytes)
	}
	return r, bytes
}
