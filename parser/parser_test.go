package parser

import (
	"fmt"
	"testing"

	"github.com/eyanshu1997/yacgo/ast"
	"github.com/eyanshu1997/yacgo/common/log"
	"github.com/eyanshu1997/yacgo/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},

		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integ.TokenLiteral())
		return false
	}
	return true
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"i = -a * b;",
			"i = ((-a) * b);",
		},
		{
			"i = !-a;",
			"i = (!(-a));",
		},
		{
			"i = a + b + c;",
			"i = ((a + b) + c);",
		},
		{
			"i = a + b - c;",
			"i = ((a + b) - c);",
		},
		{
			"i = a * b * c;",
			"i = ((a * b) * c);",
		},
		{
			"i = a * b / c;",
			"i = ((a * b) / c);",
		},
		{
			"i = a + b / c;",
			"i = (a + (b / c));",
		},
		{
			"i = a + b * c + d / e - f;",
			"i = (((a + (b * c)) + (d / e)) - f);",
		},
		{
			"i = (3 + 4) * (-5 * 5);",
			"i = ((3 + 4) * ((-5) * 5));",
		},
		{
			"i = 5 > 4 == 3 < 4;",
			"i = ((5 > 4) == (3 < 4));",
		},
		{
			"i = 5 < 4 != 3 > 4;",
			"i = ((5 < 4) != (3 > 4));",
		},
		{
			"i = 3 + 4 * 5 == 3 * 1 + 4 * 5;",
			"i = ((3 + (4 * 5)) == ((3 * 1) + (4 * 5)));",
		},
		{
			"i = 3 + 4 * 5 == 3 * 1 + 4 * 5;",
			"i = ((3 + (4 * 5)) == ((3 * 1) + (4 * 5)));",
		},
		{
			"i = true;",
			"i = true;",
		},
		{
			"i = false;",
			"i = false;",
		},
		{
			"i = 3 > 5 == false;",
			"i = ((3 > 5) == false);",
		},
		{
			"i = 3 < 5 == true;",
			"i = ((3 < 5) == true);",
		},
		{
			"i =1 + (2 + 3) + 4;",
			"i = ((1 + (2 + 3)) + 4);",
		},
		{
			"i =(5 + 5) * 2;",
			"i = ((5 + 5) * 2);",
		},
		{
			"i =2 / (5 + 5);",
			"i = (2 / (5 + 5));",
		},
		{
			"i =-(5 + 5);",
			"i = (-(5 + 5));",
		},
		{
			"i =!(true == true);",
			"i = (!(true == true));",
		},
	}
	for i, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		log.Printf("For test %d test[%s]", i, tt)
		checkParserErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("test [%d] expected=%q, got=%q", i, tt.expected, actual)
		}
	}
}
func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"i=5 + 5;", 5, "+", 5},
		{"i=5 - 5;", 5, "-", 5},
		{"i=5 * 5;", 5, "*", 5},
		{"i=5 / 5;", 5, "/", 5},
		{"i=5 > 5;", 5, ">", 5},
		{"i=5 < 5;", 5, "<", 5},
		{"i=5 == 5;", 5, "==", 5},
		{"i=5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.AssignmentStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Value.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Value)
		}
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestCallExpressionParsing(t *testing.T) {
	input := "a=add(1, 2 * 3, 4 + 5);"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
			1, len(program.Statements))

	}
	stmt, ok := program.Statements[0].(*ast.AssignmentStatement)
	if !ok {
		t.Fatalf("stmt is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	exp, ok := stmt.Value.(*ast.CallExpression)
	if !ok {
		t.Fatalf("stmt.Expression is not ast.CallExpression. got=%T",
			stmt.Value)
	}
	if !testIdentifier(t, exp.Function, "add") {
		return
	}
	if len(exp.Arguments) != 3 {
		t.Fatalf("wrong length of arguments. got=%d", len(exp.Arguments))
	}
	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {

		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value,
			ident.TokenLiteral())
		return false
	}
	return true
}

func testLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
	operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.OperatorExpression. got=%T(%s)", exp, exp)
		return false
	}
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}
	if !testLiteralExpression(t, opExp.Right, right) {

		return false
	}
	return true
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "i=5;"
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d",
			len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.AssignmentStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}
	literal, ok := stmt.Value.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Value)
	}
	if literal.Value != 5 {
		t.Errorf("literal.Value not %d. got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "5",
			literal.TokenLiteral())
	}
}

func TestIfExpression(t *testing.T) {
	input := `if (x < y) { x=1; }`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])

	}
	if !testInfixExpression(t, stmt.Condition, "x", "<", "y") {
		return
	}
	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n",
			len(stmt.Consequence.Statements))
	}
	consequence, ok := stmt.Consequence.Statements[0].(*ast.AssignmentStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			stmt.Consequence.Statements[0])
	}
	if consequence.String() != "x = 1;" {
		t.Fatalf("not matched expression [%s] [%s]",
			consequence.Value.String(), "x = 1;")
		return
	}
	if stmt.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil. got=%+v", stmt.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := `if (x < y) 
	{ 
		x=1; 
	}
	else {
		x=10;
	}
	`
	l := lexer.NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 1 {
		t.Fatalf("program.Body does not contain %d statements. got=%d\n",
			1, len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.IfStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])

	}
	if !testInfixExpression(t, stmt.Condition, "x", "<", "y") {
		return
	}
	if len(stmt.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n",
			len(stmt.Consequence.Statements))
	}
	consequence, ok := stmt.Consequence.Statements[0].(*ast.AssignmentStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			stmt.Consequence.Statements[0])
	}
	if consequence.String() != "x = 1;" {
		t.Fatalf("not matched expression [%s] [%s]",
			consequence.Value.String(), "x = 1;")
		return
	}
	alternate, ok := stmt.Alternative.Statements[0].(*ast.AssignmentStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T",
			stmt.Alternative.Statements[0])
	}
	if alternate.String() != "x = 10;" {
		t.Fatalf("not matched expression [%s] [%s]",
			consequence.Value.String(), "x = 10;")
		return
	}
}
