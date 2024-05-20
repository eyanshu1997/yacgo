package tokens

const (
	TokenTypeIllegal    TokenType = "illegal"
	TokenTypeEOF        TokenType = "EOF"
	TokenTypeIdentifier TokenType = "IDENTIFIER"
	TokenTypeInt        TokenType = "INT"

	// Operators
	TokenTypeAssign  TokenType = "="
	TokenTypeExclaim TokenType = "!"
	TokenTypeLT      TokenType = "<"
	TokenTypeGT      TokenType = ">"

	TokenTypePlus     TokenType = "+"
	TokenTypeSubtract TokenType = "-"
	TokenTypeAstrisk  TokenType = "*"
	TokenTypeDivide   TokenType = "/"

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
	TokenTypeTrue     TokenType = "TRUE"
	TokenTypeFalse    TokenType = "FALSE"
	TokenTypeIf       TokenType = "IF"
	TokenTypeElse     TokenType = "ELSE"
	TokenTypeReturn   TokenType = "RETURN"

	//MultiToken
	TokenTypeEQ    TokenType = "=="
	TokenTypeNotEQ TokenType = "!="
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
