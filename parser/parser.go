package parser

import (
	"github.com/eyanshu1997/yacgo/ast"
	"github.com/eyanshu1997/yacgo/lexer"
	"github.com/eyanshu1997/yacgo/tokens"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l              *lexer.Lexer
	curToken       tokens.Token
	peekToken      tokens.Token
	errors         []string
	prefixParseFns map[tokens.TokenType]prefixParseFn
	infixParseFns  map[tokens.TokenType]infixParseFn
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	// Read two tokens, so curToken and peekToken are both set
	p.prefixParseFns = make(map[tokens.TokenType]prefixParseFn)
	p.registerPrefix(tokens.TokenTypeIdentifier, p.parseIdentifier)
	p.registerPrefix(tokens.TokenTypeInt, p.parseIntegerLiteral)
	p.registerPrefix(tokens.TokenTypeExclaim, p.parsePrefixExpression)
	p.registerPrefix(tokens.TokenTypeSubtract, p.parsePrefixExpression)
	p.infixParseFns = make(map[tokens.TokenType]infixParseFn)
	p.registerInfix(tokens.TokenTypePlus, p.parseInfixExpression)
	p.registerInfix(tokens.TokenTypeSubtract, p.parseInfixExpression)
	p.registerInfix(tokens.TokenTypeDivide, p.parseInfixExpression)
	p.registerInfix(tokens.TokenTypeAstrisk, p.parseInfixExpression)
	p.registerInfix(tokens.TokenTypeEQ, p.parseInfixExpression)
	p.registerInfix(tokens.TokenTypeNotEQ, p.parseInfixExpression)
	p.registerInfix(tokens.TokenTypeLT, p.parseInfixExpression)
	p.registerInfix(tokens.TokenTypeGT, p.parseInfixExpression)
	p.registerPrefix(tokens.TokenTypeTrue, p.parseBoolean)
	p.registerPrefix(tokens.TokenTypeFalse, p.parseBoolean)
	p.registerPrefix(tokens.TokenTypeLParen, p.parseGroupedExpression)
	p.registerInfix(tokens.TokenTypeLParen, p.parseCallExpression)

	return p
}

func (p *Parser) parseIfStatement() ast.Statement {

	expression := &ast.IfStatement{Token: p.curToken}
	if !p.expectPeek(tokens.TokenTypeLParen) {
		return nil
	}
	p.nextToken()
	//log.Printf("ifstatement found  %s", p.curToken)
	expression.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(tokens.TokenTypeRParen) {
		return nil
	}
	if !p.expectPeek(tokens.TokenTypeLBrace) {
		return nil
	}
	expression.Consequence = p.parseBlockStatement()
	if p.peekTokenIs(tokens.TokenTypeElse) {

		p.nextToken()
		//log.Printf("Looking in else block curtoken [%s]", p.curToken)
		if !p.expectPeek(tokens.TokenTypeLBrace) {
			return nil
		}
		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseLetStatement() ast.Statement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(tokens.TokenTypeIdentifier) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(tokens.TokenTypeAssign) {
		return nil
	}
	p.nextToken()
	//log.Printf("Found let statement [%s]: [%s]", stmt, p.curToken)
	stmt.Value = p.parseExpression(LOWEST)
	if !p.expectPeek(tokens.TokenTypeSemiColon) {
		return nil
	}
	return stmt
}

func (p *Parser) parseReturnStatement() ast.Statement {
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	//log.Printf("Found return statement [%s]: [%s]", stmt, p.curToken)
	stmt.ReturnValue = p.parseExpression(LOWEST)
	if !p.expectPeek(tokens.TokenTypeSemiColon) {
		return nil
	}
	return stmt
}

func (p *Parser) parseFunctionStatement() ast.Statement {
	stmt := &ast.FunctionStatement{Token: p.curToken}
	//TODO implement this
	return stmt
}

func (p *Parser) parseIdentifierStatement() ast.Statement {
	stmt := &ast.AssignmentStatement{Token: p.curToken}
	if !p.expectPeek(tokens.TokenTypeAssign) {
		return nil
	}
	p.nextToken()
	//log.Printf("Found assignment statement [%s]: [%s]", stmt, p.curToken)
	stmt.Value = p.parseExpression(LOWEST)
	if !p.expectPeek(tokens.TokenTypeSemiColon) {
		return nil
	}
	return stmt
}

func (p *Parser) parseStatement() ast.Statement {
	//log.Printf("Parse Statement called token : %s %s", p.curToken, p.peekToken)
	switch p.curToken.Type {
	case tokens.TokenTypeLet:
		return p.parseLetStatement()
	case tokens.TokenTypeReturn:
		return p.parseReturnStatement()
	case tokens.TokenTypeFunction:
		return p.parseFunctionStatement()
	case tokens.TokenTypeIf:
		return p.parseIfStatement()
	case tokens.TokenTypeIdentifier:
		return p.parseIdentifierStatement()
	default:
		return nil
	}
}

// always check for errors after running this
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.curToken.Type != tokens.TokenTypeEOF {
		stmt := p.parseStatement()
		if stmt != nil {
			//log.Printf("Got statement %s", stmt)
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}
