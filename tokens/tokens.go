package tokens

const (
	TokenTypeIllegal TokenType = "illegal"
	TokenTypeEOF     TokenType = "EOF"
)

var (
	multiToken = []byte{'='}
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal byte) *Token {
	return &Token{Type: tokenType, Literal: string(literal)}
}

func CanHaveNextToken(ch byte) bool {
	for _, b := range multiToken {
		if b == ch {
			return true
		}
	}
	return false
}
