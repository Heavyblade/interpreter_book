package ast

import (
	"bytes"
	"interpreter/token"
)

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (infx *InfixExpression) expressionNode()      {}
func (infx *InfixExpression) TokenLiteral() string { return infx.Token.Literal }
func (infx *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infx.Left.String())
	out.WriteString(" " + infx.Operator + " ")
	out.WriteString(infx.Right.String())
	out.WriteString(")")

	return out.String()
}
