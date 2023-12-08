package ast

import "github.com/houcine7/JIPL/token"

/*
* This node is the root of each AST tree
* A program is a set of statements
 */
type Program struct {
	Statements []Statement
}

func (prog *Program) TokenLiteral() string{
	
	if len(prog.Statements) > 0 { 
		return prog.Statements[0].TokenLiteral()
	}

	return ""
}


/*
* Def statement Node (def x = add(1,2) - 1 + (10/5))
*/
type Identifier struct{ // 
	Token token.Token // token.IDENTIFIER token
	Value string
}

// to imp Nodes 
func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Value;
}

// to imp Expression Node
// even do in let identifiers are not expressions but in other 
//parts they does provide a value
func (ident *Identifier) expressionNode() {}



type DefStatement struct {
	Token token.Token // toke.DEF token
	Name *Identifier
	Value Expression
}
// method satisfies the Node interface
func (defSt *DefStatement) TokenLiteral() string{
	return defSt.Token.Value;
}
// satisfies the statement interface 
func (defSt *DefStatement) statementNode() {}
