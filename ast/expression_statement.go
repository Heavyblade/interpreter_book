package ast

import "interpreter/token"

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (l *ExpressionStatement) statementNode()       {}
func (l *ExpressionStatement) TokenLiteral() string { return l.Token.Literal }
func (l *ExpressionStatement) String() string {
	return l.Expression.String()
}
