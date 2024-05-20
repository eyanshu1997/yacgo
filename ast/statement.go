package ast

import "github.com/eyanshu1997/yacgo/tokens"

type LetStatement struct {
	Token tokens.Token // let
	Name  *Identifier  // identifier
	Value Expression   // expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
