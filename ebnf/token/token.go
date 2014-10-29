package token

import (
	"asciigoat.org/core/scanner"
)

// A Token is a Terminal with known type
type Token struct {
	scanner.Terminal

	Typ TokenType
}
