package tokens

var keywordsMap = map[string]TokenType{
	"fn":  TokenTypeFunction,
	"let": TokenTypeLet,
}

func CheckIfKeywordType(literal string) TokenType {
	if tokType, ok := keywordsMap[literal]; ok {
		return tokType
	}
	return TokenTypeIdentifier
}
