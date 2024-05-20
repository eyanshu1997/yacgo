package ast

import "github.com/eyanshu1997/yacgo/tokens"

type Identifier struct {
	Token tokens.Token // token { identifier, literal}
	Value string       // literal
}

// These will be used in ast
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
