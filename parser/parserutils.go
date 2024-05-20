package parser

import (
	"fmt"

	"github.com/eyanshu1997/yacgo/common/log"
	"github.com/eyanshu1997/yacgo/tokens"
)

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
