package token

// types of Token
type TokenType int

const (
	TokenError	TokenType = iota + 1
	TokenEOF
)

func (typ TokenType) String() string {
	switch typ {
	case TokenError:
		return "ERROR"
	case TokenEOF:
		return "EOF"
	default:
		return "UNDEFINED"
	}
}
