package tokens

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
