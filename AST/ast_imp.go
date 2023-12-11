package ast

import (
	"bytes"

	"github.com/houcine7/JIPL/token"
)

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

func (prog *Program) ToString() string {
	var bf bytes.Buffer

	for _,stm := range prog.Statements {
		bf.WriteString(stm.ToString())
	}

	return bf.String()
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
func (defStm *DefStatement) TokenLiteral() string{
	return defStm.Token.Value;
}
// toString method from Node interface
func (defStm *DefStatement) ToString() string{
	var bf bytes.Buffer

	bf.WriteString(defStm.TokenLiteral()+" ")
	bf.WriteString(defStm.Name.ToString())
	bf.WriteString(" = ")
	if defStm.Value !=nil{
		bf.WriteString(defStm.Value.ToString())
	} 
	bf.WriteString(";")
	return bf.String()
}
// satisfies the statement interface 
func (defStm *DefStatement) statementNode() {}




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

func (ident *Identifier) ToString() string{
	return ident.Value
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

//Node interface methods 
func (reStm *ReturnStatement) TokenLiteral() string{  // satisfies the node interface 
	return reStm.Token.Value
}

func (resStm *ReturnStatement) ToString() string {
	var bf bytes.Buffer
	bf.WriteString(resStm.TokenLiteral())
	if resStm.ReturnValue !=nil{
		bf.WriteString(resStm.ReturnValue.ToString())
	}
	bf.WriteString(";")
	return bf.String()
}

// statements imp
func (reStm *ReturnStatement) statementNode() { } // satisfies the statement interface




/*
* Expressions statement node
* they are wrappers that consists solely of one expression
*/
type ExpressionStatement struct{
	Token token.Token
	Expression Expression
}
// Node's interface methods 
func (exStm *ExpressionStatement) TokenLiteral() string{
	return exStm.Token.Value
}

func (exStm *ExpressionStatement) ToString() string{
	var bf bytes.Buffer
	if exStm.Expression !=nil {
		bf.WriteString(exStm.Expression.ToString())
	}
	return bf.String()
}

func (exStm *ExpressionStatement) statementNode(){}




/*
* Integer Literals Node
* they can Occur in different type of expression's
*/
type IntegerLiteral struct {
	Token token.Token
	Value int // we do not specify int size to make it platform independent (32,64)
}

func (intLiteral *IntegerLiteral) TokenLiteral() string{
	return intLiteral.Token.Value
}

func (intLiteral *IntegerLiteral) expressionNode() {}

func (intLiteral *IntegerLiteral) ToString() string {
	return intLiteral.Token.Value
}


/*
	Prefix Expression Nodes
	-5; !8787;
*/
type PrefixExpression struct {
	Token token.Token
	Operator string
	Right Expression
}

func (prefixExp *PrefixExpression) TokenLiteral() string{
	return prefixExp.Token.Value
}

func (prefixExp *PrefixExpression) ToString() string{
	var bf bytes.Buffer

	bf.WriteRune('(')
	bf.WriteString(prefixExp.Operator)
	bf.WriteString(prefixExp.Right.ToString())
	bf.WriteRune(')')

	return bf.String()
}

func (prefixExp *PrefixExpression) expressionNode() {}