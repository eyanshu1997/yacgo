package lexer

import (
	"github.com/eyanshu1997/yacgo/tokens"
	"github.com/eyanshu1997/yacgo/utils"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readNextChar()
	return l
}

func (l *Lexer) readNextChar() {
	// log.Printf("readNextChar Called readPosition %d position %d len input %d", l.readPosition, l.position, len(l.input))
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for utils.IsLetter(l.ch) {
		l.readNextChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readNextChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for utils.IsDigit(l.ch) {
		l.readNextChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekNextChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}

}

func (l *Lexer) getMultiToken() *tokens.Token {
	if l.ch == '=' && l.peekNextChar() == '=' {
		ch := l.ch
		l.readNextChar()
		return &tokens.Token{Type: tokens.TokenTypeEQ, Literal: string(ch) + string(l.ch)}
	}
	if l.ch == '=' && l.peekNextChar() == '!' {
		ch := l.ch
		l.readNextChar()
		return &tokens.Token{Type: tokens.TokenTypeNotEQ, Literal: string(ch) + string(l.ch)}
	}
	return nil
}

func (l *Lexer) ReadNextToken() *tokens.Token {
	var tok = &tokens.Token{}
	l.skipWhitespace()
	if tokens.CanHaveNextToken(l.ch) {
		tok = l.getMultiToken()
		if tok != nil {
			return tok
		}
	}
	switch l.ch {
	case ',':
		tok = tokens.NewToken(tokens.TokenTypeComma, l.ch)
	case ';':
		tok = tokens.NewToken(tokens.TokenTypeSemiColon, l.ch)
	case '(':
		tok = tokens.NewToken(tokens.TokenTypeLParen, l.ch)
	case ')':
		tok = tokens.NewToken(tokens.TokenTypeRParen, l.ch)
	case '{':
		tok = tokens.NewToken(tokens.TokenTypeLBrace, l.ch)
	case '}':
		tok = tokens.NewToken(tokens.TokenTypeRBrace, l.ch)

	case '=':
		tok = tokens.NewToken(tokens.TokenTypeAssign, l.ch)
	case '!':
		tok = tokens.NewToken(tokens.TokenTypeExclaim, l.ch)
	case '<':
		tok = tokens.NewToken(tokens.TokenTypeLT, l.ch)
	case '>':
		tok = tokens.NewToken(tokens.TokenTypeGT, l.ch)

	case '+':
		tok = tokens.NewToken(tokens.TokenTypePlus, l.ch)
	case '-':
		tok = tokens.NewToken(tokens.TokenTypeSubtract, l.ch)
	case '*':
		tok = tokens.NewToken(tokens.TokenTypeAstrisk, l.ch)
	case '/':
		tok = tokens.NewToken(tokens.TokenTypeDivide, l.ch)

	case 0:
		tok.Type = tokens.TokenTypeEOF
	default:
		if utils.IsLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = tokens.CheckIfKeywordType(tok.Literal)
			return tok
		} else if utils.IsDigit(l.ch) {
			tok.Type = tokens.TokenTypeInt
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = tokens.NewToken(tokens.TokenTypeIllegal, l.ch)
		}
	}
	l.readNextChar()
	return tok
}
