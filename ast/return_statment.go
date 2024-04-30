package ast

import (
	"bytes"
	"interpreter/token"
)

type ReturnStatement struct {
	Token       token.Token // the RETURN token
	ReturnValue Expression
}

func (i *ReturnStatement) expressionNode()      {}
func (l *ReturnStatement) statementNode()       {}
func (l *ReturnStatement) TokenLiteral() string { return l.Token.Literal }

func (r *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(r.Token.Literal + " ")

	if r.ReturnValue != nil {
		out.WriteString(r.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
