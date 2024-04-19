package parser

import (
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
