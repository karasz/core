package lexer

import (
	"asciigoat.org/core/ebnf/token"
	"testing"
)

func cmp(t *testing.T, l *Lexer, tokens []*token.Token) {
	i := 0

	for {
		var x, y *token.Token

		x = l.NextToken()
		if i < len(tokens) {
			y = tokens[i]
		}
		i++

		if x == nil || y == nil || *x != *y {
			t.Errorf("token[%v] failed: %v != %v", i, x, y)
		} else {
			t.Logf("token[%v]", i, x)
		}

		// Lexer stopped too soon
		if x == nil || x.Typ == token.TokenEOF {
			x = nil

			// print all the rest of the expected tokens
			for i < len(tokens) {
				y = tokens[i]
				i++
				t.Errorf("token[%v] failed: %v != %v", i, x, y)
			}
		}
	}
}

// Really Empty
func TestEmpty1(t *testing.T) {
	s := ""
	l := newLexer(t, s)

	cmp(t, l, []*token.Token{
		l.Token(token.TokenEOF, "").Loc(1, 1),
	})
}
