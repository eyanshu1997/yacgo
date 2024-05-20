package tokens

const (
	// Keywords
	TokenTypeFunction TokenType = "FUNCTION"
	TokenTypeLet      TokenType = "LET"
	TokenTypeTrue     TokenType = "TRUE"
	TokenTypeFalse    TokenType = "FALSE"
	TokenTypeIf       TokenType = "IF"
	TokenTypeElse     TokenType = "ELSE"
	TokenTypeReturn   TokenType = "RETURN"
)

var keywordsMap = map[string]TokenType{
	"fn":     TokenTypeFunction,
	"let":    TokenTypeLet,
	"true":   TokenTypeTrue,
	"false":  TokenTypeFalse,
	"if":     TokenTypeIf,
	"else":   TokenTypeElse,
	"return": TokenTypeReturn,
}

func CheckIfKeywordType(literal string) TokenType {
	if tokType, ok := keywordsMap[literal]; ok {
		return tokType
	}
	return TokenTypeIdentifier
}
