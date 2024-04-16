package ast

import (
	"bytes"
	"interpreter/token"
)

// Type definitions
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Identifier struct {
	Token token.Token // the IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string {
	return i.Value
}

type LetStatement struct {
	Token token.Token // the LET token
	Value Expression
	Name  *Identifier
}

func (l *LetStatement) statementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

func (l *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(l.Token.Literal + " ")
	out.WriteString(l.Name.String() + " ")
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	return out.String()
}

type ReturnStatement struct {
	Token       token.Token // the RETURN token
	ReturnValue Expression
}

func (l *ReturnStatement) statementNode()       {}
func (l *ReturnStatement) TokenLiteral() string { return l.Token.Literal }

func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.Token.Literal + " ")

	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}

	out.WriteString(" ; ")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (l *ExpressionStatement) statementNode()       {}
func (l *ExpressionStatement) TokenLiteral() string { return l.Token.Literal }
