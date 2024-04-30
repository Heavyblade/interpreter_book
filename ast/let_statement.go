package ast

import (
	"bytes"
	"interpreter/token"
)

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
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
