package lexer

import (
	"testing"

	"github.com/eyanshu1997/yacgo/tokens"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.TokenTypeAssign, "="},
		{tokens.TokenTypePlus, "+"},
		{tokens.TokenTypeLParen, "("},
		{tokens.TokenTypeRParen, ")"},
		{tokens.TokenTypeLBrace, "{"},
		{tokens.TokenTypeRBrace, "}"},
		{tokens.TokenTypeComma, ","},
		{tokens.TokenTypeSemiColon, ";"},
		{tokens.TokenTypeEOF, ""},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.ReadNextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLotsOfTokens(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y) {
	x + y;
	};
	let result = add(five, ten);`
	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.TokenTypeLet, "let"},
		{tokens.TokenTypeIdentifier, "five"},
		{tokens.TokenTypeAssign, "="},
		{tokens.TokenTypeInt, "5"},
		{tokens.TokenTypeSemiColon, ";"},
		{tokens.TokenTypeLet, "let"},
		{tokens.TokenTypeIdentifier, "ten"},
		{tokens.TokenTypeAssign, "="},
		{tokens.TokenTypeInt, "10"},
		{tokens.TokenTypeSemiColon, ";"},
		{tokens.TokenTypeLet, "let"},
		{tokens.TokenTypeIdentifier, "add"},
		{tokens.TokenTypeAssign, "="},
		{tokens.TokenTypeFunction, "fn"},
		{tokens.TokenTypeLParen, "("},
		{tokens.TokenTypeIdentifier, "x"},
		{tokens.TokenTypeComma, ","},
		{tokens.TokenTypeIdentifier, "y"},
		{tokens.TokenTypeRParen, ")"},
		{tokens.TokenTypeLBrace, "{"},
		{tokens.TokenTypeIdentifier, "x"},
		{tokens.TokenTypePlus, "+"},
		{tokens.TokenTypeIdentifier, "y"},
		{tokens.TokenTypeSemiColon, ";"},
		{tokens.TokenTypeRBrace, "}"},
		{tokens.TokenTypeSemiColon, ";"},
		{tokens.TokenTypeLet, "let"},
		{tokens.TokenTypeIdentifier, "result"},
		{tokens.TokenTypeAssign, "="},
		{tokens.TokenTypeIdentifier, "add"},
		{tokens.TokenTypeLParen, "("},
		{tokens.TokenTypeIdentifier, "five"},
		{tokens.TokenTypeComma, ","},
		{tokens.TokenTypeIdentifier, "ten"},
		{tokens.TokenTypeRParen, ")"},
		{tokens.TokenTypeSemiColon, ";"},
		{tokens.TokenTypeEOF, ""},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.ReadNextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestMultiToken(t *testing.T) {
	input := `==5!=`
	tests := []struct {
		expectedType    tokens.TokenType
		expectedLiteral string
	}{
		{tokens.TokenTypeEQ, "=="},
		{tokens.TokenTypeInt, "5"},
		{tokens.TokenTypeNotEQ, "!="},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.ReadNextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
