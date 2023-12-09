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


/*
* Identifier node as an expression node
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

/*
* return statement node 
* return 5; || return f1(1,2); ==> return <expression>;
*/
type ReturnStatement struct{
	Token token.Token // the token is "return"
	ReturnValue Expression
}

func (reStm *ReturnStatement) TokenLiteral() string{  // satisfies the node interface 
	return reStm.Token.Value
}
func (reStm *ReturnStatement) statementNode() { // satisfies the statement interface
}
