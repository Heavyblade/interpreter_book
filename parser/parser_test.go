package parser

import (
	"fmt"
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	  let x = 5;
	  let y = 10;
	  let foobar = 838383;
	`

	lex := lexer.New(input)
	p := New(lex)

	program := p.ParseProgram()

	if program == nil {
		t.Fatal("ParseProgram return nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Statements number is wrong it should have 3 but has (%d) statements", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, v := range tests {
		if testLetStatement(t, program.Statements[i], v.expectedIdentifier) == false {
			return
		}
	}
}

func TestParserErrors(t *testing.T) {
	input := `
	  let x 5;
	  let = 10;
	  let 838383;
	`

	lex := lexer.New(input)
	parser := New(lex)

	parser.ParseProgram()

	if len(parser.Errors()) == 0 {
		t.Fatalf("There are no errors present, expected 3")
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	  return 5;
	  return 10;
	  return 838383;
	`

	lex := lexer.New(input)
	p := New(lex)

	program := p.ParseProgram()

	if len(program.Statements) != 3 {
		t.Fatalf("Statements number is wrong it should have 3 but has (%d) statements", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt is not a *ReturnStatement, got =%T", returnStmt)
			continue
		}

		if returnStmt.Token.Literal != "return" {
			t.Errorf("returnStmt.TokenLiteral is not 'return', got %s", returnStmt.Token.Literal)
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	lex := lexer.New(input)
	p := New(lex)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Statements number is wrong it should have 1 but has (%d) statements", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an *ast.ExpressionStatement, got: %T", program.Statements[0])

	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("stmt.Expression is not an *ast.Identifier, got: %T", program.Statements[0])
	}

	if ident.Value != "foobar" {
		t.Fatalf("ident.Value is != from 'foobar', got: %s", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.Value is != from 'foobar', got: %s", ident.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	lex := lexer.New(input)
	p := New(lex)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Statements number is wrong it should have 1 but has (%d) statements", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statemens[0] is not an *ast.ExpressionStatement, got: %T", program.Statements[0])
	}

	intLt, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stmt is not an *ast.IntegerLiteral, got: %T", stmt)
	}

	if intLt.Value != 5 {
		t.Fatalf("Expected value for IntegerLiteral should be 5, got: %d", intLt.Value)
	}

	if intLt.TokenLiteral() != "5" {
		t.Fatalf("Expected value for IntegerLiteral should be '5', got: %s", intLt.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTest := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTest {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("Statements number is wrong it should have 1 but has (%d) statements", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statemens[0] is not an *ast.ExpressionStatement, got: %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("smt is not a *ast.PrefixExpression, got: %T", exp)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("the expected operator was '%s', got: '%s'", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}

	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTest := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 / 5", 5, "/", 5},
		{"5 > 5", 5, ">", 5},
		{"5 < 5", 5, "<", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
	}

	for _, tt := range infixTest {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("Statements number is wrong it should have 1 but has (%d) statements", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statemens[0] is not an *ast.ExpressionStatement, got: %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt.Expression is not an *ast.InfixExpression, got: %T", stmt.Expression)
		}

		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s, got: %s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}

	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let, got: %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not an ast.LetStatement, got: %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value is not (%s) correct, got: %s", name, letStmt.Name.Value)
		return false
	}

	// prevent a variable called 'let'
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not %s, got: %s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, error := range errors {
		t.Errorf("parser error: %s", error)
	}

	t.FailNow()
}

func testIntegerLiteral(t *testing.T, right ast.Expression, value int64) bool {
	il, ok := right.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il is not an *ast.IntegerLiteral, got: %T", il)
		return false
	}

	if il.Value != value {
		t.Errorf("il.Value not %d, got: %d", value, il.Value)
		return false
	}

	if il.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("il.TokenLiteral not %d, got: %s", value, il.TokenLiteral())
		return false
	}

	return true
}
