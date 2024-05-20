package tokens

const (
	// Operators
	TokenTypeAssign  TokenType = "="
	TokenTypeExclaim TokenType = "!"
	TokenTypeLT      TokenType = "<"
	TokenTypeGT      TokenType = ">"

	TokenTypePlus     TokenType = "+"
	TokenTypeSubtract TokenType = "-"
	TokenTypeAstrisk  TokenType = "*"
	TokenTypeDivide   TokenType = "/"

	//MultiToken operators
	TokenTypeEQ    TokenType = "=="
	TokenTypeNotEQ TokenType = "!="
)
