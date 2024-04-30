package ast

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
