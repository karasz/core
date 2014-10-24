package scanner

import (
	"unicode/utf8"
)

const (
	EOF = -1
)

type Position struct {
	Offset       uint
	Line, Column uint
}

type Scanner struct {
	name  string
	input string

	base   Position
	cursor Position
	runes  uint
}

// Creates new Scanner to parse a given string, or nil if empty
func NewScannerFromString(name, input string) *Scanner {
	if len(input) == 0 {
		return nil
	}

	return &Scanner{
		name:   name,
		input:  input,
		base:   Position{0, 1, 1},
		cursor: Position{0, 1, 1},
		runes:  0,
	}
}

// Current length of the upcomming Terminal
func (l *Scanner) Length() (uint, uint) {
	return l.cursor.Offset - l.base.Offset, l.runes
}

// Is the upcoming Terminal stil empty?
func (l *Scanner) Empty() bool {
	return l.runes == 0
}

// Move cursor forward
func (l *Scanner) StepForth(runes, bytes uint) {
	l.cursor.Offset += bytes
	l.cursor.Column += runes
	l.runes += runes
}

// Move cursor backward
func (l *Scanner) StepBack(runes, bytes uint) {
	l.cursor.Offset -= bytes
	// FIXME: what if column goes < 1?
	l.cursor.Column -= runes
	l.runes -= runes
}

// Moves the cursor back to the back
func (l *Scanner) Reset() {
	l.cursor = l.base
	l.runes = 0
}

// Trashes everything up to the cursor
func (l *Scanner) Skip() {
	l.base = l.cursor
	l.runes = 0
}

// Count NewLine
func (l *Scanner) NextLine() {
	l.cursor.Line++
	l.cursor.Column = 1
}

// Return the next rune but not moving the cursor
func (l *Scanner) Peek() (rune, uint) {
	if l.cursor.Offset == uint(len(l.input)) {
		return EOF, 0
	}
	r, bytes := utf8.DecodeRuneInString(l.input[l.cursor.Offset:])
	return r, uint(bytes)
}

// Return the next rune but moving the cursor
func (l *Scanner) Next() (rune, uint) {
	r, bytes := l.Peek()
	if bytes > 0 {
		l.StepForth(1, bytes)
	}
	return r, bytes
}
