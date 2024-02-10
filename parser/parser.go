package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken   token.Token
	peekTooken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.NextToken()
	p.NextToken()

	return p
}

func (p *Parser) NextToken() {
	p.curToken = p.peekTooken
	p.peekTooken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parserLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parserLetStatement() ast.Statement {
	letStmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Name assingation after checking that next token was a token.IDENT
	letStmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// ignore the expression until we find a semicolon
	for !p.curTokenIs(token.SEMICOLON) {
		p.NextToken()
	}

	return letStmt
}

// Instead of moving and check if the current token is the expected token
// Here we check the peekToken and the move if is the correct one
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.NextToken()
		return true
	}

	return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekTooken.Type == t
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}
