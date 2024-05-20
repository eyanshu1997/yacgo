package parser

import (
	"fmt"
	"log"

	"github.com/eyanshu1997/yacgo/ast"
	"github.com/eyanshu1997/yacgo/lexer"
	"github.com/eyanshu1997/yacgo/tokens"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  tokens.Token
	peekToken tokens.Token
	errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = *p.l.ReadNextToken()
}

func (p *Parser) peekError(t tokens.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	log.Println(msg)
	p.errors = append(p.errors, msg)
}
func (p *Parser) Errors() []string {
	return p.errors
}
func (p *Parser) curTokenIs(t tokens.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t tokens.TokenType) bool {
	return p.peekToken.Type == t
}

// checks next token if it is the expected one increment,else return
func (p *Parser) expectPeek(t tokens.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) ParseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(tokens.TokenTypeIdentifier) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(tokens.TokenTypeAssign) {
		return nil
	}
	//TODO implement handling for expressions
	for !p.curTokenIs(tokens.TokenTypeSemiColon) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) ParseStatement() ast.Statement {
	log.Printf("Parse Statement called token : %s %s", p.curToken, p.peekToken)
	switch p.curToken.Type {
	case tokens.TokenTypeLet:
		return p.ParseLetStatement()
	default:
		return nil
	}
}

// always check for errors after running this
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != tokens.TokenTypeEOF {
		stmt := p.ParseStatement()
		if stmt != nil {
			log.Printf("Got statement %s", stmt)
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}
