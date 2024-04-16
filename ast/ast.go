package ast

import "interpreter/token"

// Type definitions
type Node interface {
	TokenLiteral() string
}

// A statement is a piece of code that does NOT produce a value
type Statement interface {
	Node
	statementNode()
}

// An expression is a piece of code that produces a value
// Produce value means that it takes some input or call and returns something different
type Expression interface {
	Node
	expressionNode()
}

// AST Implementation
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token // the IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type LetStatement struct {
	Token token.Token // the LET token
	Value Expression
	Name  *Identifier
}

func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // the RETURN token
	ReturnValue Expression
}

func (l *ReturnStatement) statementNode()       {}
func (l *ReturnStatement) TokenLiteral() string { return l.Token.Literal }
