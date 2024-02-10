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
