package ast

import (
	"testing"

	"github.com/eyanshu1997/yacgo/tokens"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: tokens.Token{Type: tokens.TokenTypeLet, Literal: "let"},
				Name: &Identifier{
					Token: tokens.Token{Type: tokens.TokenTypeIdentifier, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: tokens.Token{Type: tokens.TokenTypeIdentifier, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
