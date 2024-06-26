package ast

import (
	"bytes"
	"strings"

	"github.com/eyanshu1997/yacgo/tokens"
)

type FunctionStatement struct {
	Token      tokens.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionStatement) statementNode()       {}
func (fl *FunctionStatement) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionStatement) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}
