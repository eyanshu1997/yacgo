package ast

import (
	"bytes"

	"github.com/eyanshu1997/yacgo/tokens"
)

type IfStatement struct {
	Token       tokens.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfStatement) statementNode()       {}
func (ie *IfStatement) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfStatement) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}
