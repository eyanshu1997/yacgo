package ast

import "github.com/eyanshu1997/yacgo/tokens"

type LetStatement struct {
	Token tokens.Token // let
	Name  *Identifier  // identifier
	Value Expression   // expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type ReturnStatement struct {
	Token tokens.Token // return
	Value Expression   // expression
}

func (ls *ReturnStatement) statementNode()       {}
func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }
