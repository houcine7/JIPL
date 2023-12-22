package ast

/*
* Abstract syntax tree is the data structure
* which our parser will return it's is an hierarchical DS
* that represents 'the flow of the tokens' with a top down approach
* It's called abstract because it doesn't include all elements such whitespace break lines ..
 */

/*
* The node type which be contained in the tree
 */
type Node interface{
	/* 
	* The token literal representation in the node
	* This will be used only for testing and debugging purposes 
	*/
	TokenLiteral() string  
	ToString() string
}

/*
* we wil be define for now 2 types of nodes 
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