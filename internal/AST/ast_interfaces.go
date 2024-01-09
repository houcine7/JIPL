package ast

// the nodes of the	AST tree
type Node interface {
	TokenLiteral() string
	ToString() string
}

/*
* Statements nodes: doesn't return any value
* Expression nodes : does return a value
 */
type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}
