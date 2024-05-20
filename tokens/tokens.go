package tokens

const (
	TokenTypeIllegal    TokenType = "illegal"
	TokenTypeEOF        TokenType = "EOF"
	TokenTypeIdentifier TokenType = "IDENTIFIER"
	TokenTypeInt        TokenType = "INT"

	// Operators
	TokenTypeAssign TokenType = "="
	TokenTypePlus   TokenType = "+"

	// Delimiters
	TokenTypeComma     TokenType = ","
	TokenTypeSemiColon TokenType = ";"
	TokenTypeLParen    TokenType = "("
	TokenTypeRParen    TokenType = ")"
	TokenTypeLBrace    TokenType = "{"
	TokenTypeRBrace    TokenType = "}"

	// Keywords
	TokenTypeFunction TokenType = "FUNCTION"
	TokenTypeLet      TokenType = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, literal byte) *Token {
	return &Token{Type: tokenType, Literal: string(literal)}
}
