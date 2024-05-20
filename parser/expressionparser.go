package parser

import (
	"fmt"
	"strconv"

	"github.com/eyanshu1997/yacgo/ast"
	"github.com/eyanshu1997/yacgo/common/log"
	"github.com/eyanshu1997/yacgo/tokens"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var precedences = map[tokens.TokenType]int{
	tokens.TokenTypeEQ:       EQUALS,
	tokens.TokenTypeNotEQ:    EQUALS,
	tokens.TokenTypeLT:       LESSGREATER,
	tokens.TokenTypeGT:       LESSGREATER,
	tokens.TokenTypePlus:     SUM,
	tokens.TokenTypeSubtract: SUM,
	tokens.TokenTypeDivide:   PRODUCT,
	tokens.TokenTypeAstrisk:  PRODUCT,
	tokens.TokenTypeLParen:   CALL,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) registerPrefix(tokenType tokens.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType tokens.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}
	p.nextToken()
	for !p.curTokenIs(tokens.TokenTypeRBrace) && !p.curTokenIs(tokens.TokenTypeEOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}
	return block
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(tokens.TokenTypeTrue)}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeek(tokens.TokenTypeRParen) {
		return nil
	}
	return exp
}

func (p *Parser) noPrefixParseFnError(t tokens.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	log.Printf("parseExpression current token %s", p.curToken)
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	log.Printf("parseExpression after prefixParseFns leftExp [%s] currtoen[%s] peektokenType[%s][%d]", leftExp, p.curToken, p.peekToken, p.peekPrecedence())
	for !p.peekTokenIs(tokens.TokenTypeSemiColon) && precedence < p.peekPrecedence() {
		log.Printf("inside loop for parseExpression")
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
		log.Printf("inside loop for parseExpression after infix [%s] curtoken [%s]", leftExp, p.curToken)
	}
	return leftExp
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	log.Printf("Call expresion called %s token:[%s]", function, p.curToken)
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}
func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}
	if p.peekTokenIs(tokens.TokenTypeRParen) {
		p.nextToken()
		return args
	}
	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))
	for p.peekTokenIs(tokens.TokenTypeComma) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(tokens.TokenTypeRParen) {
		return nil
	}
	return args
}
