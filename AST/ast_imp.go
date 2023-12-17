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

func (prog *Program) TokenLiteral() string {
	if len(prog.Statements) > 0 {
		return prog.Statements[0].TokenLiteral()
	}
	return ""
}

func (prog *Program) ToString() string {
	var bf bytes.Buffer

	for _, stm := range prog.Statements {
		bf.WriteString(stm.ToString())
	}

	return bf.String()
}

/*
* Def statement Node (def x = add(1,2) - 1 + (10/5))
 */
type DefStatement struct {
	Token token.Token // toke.DEF token
	Name  *Identifier
	Value Expression
}

// method satisfies the Node interface
func (defStm *DefStatement) TokenLiteral() string {
	return defStm.Token.Value
}

// toString method from Node interface
func (defStm *DefStatement) ToString() string {
	var bf bytes.Buffer

	bf.WriteString(defStm.TokenLiteral() + " ")
	bf.WriteString(defStm.Name.ToString())
	bf.WriteString(" = ")
	if defStm.Value != nil {
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
type Identifier struct { //
	Token token.Token // token.IDENTIFIER token
	Value string
}

// to imp Nodes
func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Value
}

func (ident *Identifier) ToString() string {
	return ident.Value
}

// to imp Expression Node
// even do in let identifiers are not expressions but in other
// parts they does provide a value
func (ident *Identifier) expressionNode() {}

/*
* return statement node
* return 5; || return f1(1,2); ==> return <expression>;
 */
type ReturnStatement struct {
	Token       token.Token // the token is "return"
	ReturnValue Expression
}

// Node interface methods
func (reStm *ReturnStatement) TokenLiteral() string { // satisfies the node interface
	return reStm.Token.Value
}

func (resStm *ReturnStatement) ToString() string {
	var bf bytes.Buffer
	bf.WriteString(resStm.TokenLiteral())
	if resStm.ReturnValue != nil {
		bf.WriteString(resStm.ReturnValue.ToString())
	}
	bf.WriteString(";")
	return bf.String()
}

// statements imp
func (reStm *ReturnStatement) statementNode() {} // satisfies the statement interface

/*
the ast node of boolean
*/
type BooleanExp struct {
	Token token.Token
	Value bool // the boolean value corresponds to bool
}

func (b *BooleanExp) TokenLiteral() string {
	return b.Token.Value
}

func (b *BooleanExp) ToString() string {
	return b.Token.Value
}

func (b *BooleanExp) expressionNode() {}

/*
the ast node for function expressions
(functoin definition are expressions)
*/
type FunctionExp struct {
	Token      token.Token   // the function token used to represent functions
	Name       *Identifier   // the name of the functoin
	Parameters []*Identifier // function parmas
	FnBody     *BlockStm     // function body
}

// implments Node & expression interface
func (fnExp *FunctionExp) TokenLiteral() string {
	return fnExp.Token.Value
}

func (fnExp *FunctionExp) expressionNode() {}

func (fnExp *FunctionExp) ToString() string {
	var bf bytes.Buffer

	bf.WriteString(fnExp.TokenLiteral())
	bf.WriteRune(' ')
	bf.WriteString(fnExp.Name.ToString())
	bf.WriteRune('(')

	for idx, iden := range fnExp.Parameters {
		bf.WriteString(iden.ToString())
		if idx != len(fnExp.Parameters)-1 {
			bf.WriteRune(',')
		}
	}
	bf.WriteRune(')')
	bf.WriteString(fnExp.FnBody.ToString())
	return bf.String()
}

// function invocation
type FunctionCall struct {
	Token     token.Token  // token '(' LP AST node constructs in infix pos fun()
	Function  Expression   // identifier
	Arguments []Expression // function call args
}

func (fnCall *FunctionCall) expressionNode() {}

func (fnCall *FunctionCall) TokenLiteral() string {
	return fnCall.Token.Value
}

func (fnCall *FunctionCall) ToString() string {
	var bf bytes.Buffer
	bf.WriteString(fnCall.Function.ToString())
	bf.WriteRune('(')

	for idx, arg := range fnCall.Arguments {
		bf.WriteString(arg.ToString())
		if idx != len(fnCall.Arguments)-1 {
			bf.WriteRune(',')
		}
	}

	bf.WriteRune(')')

	return bf.String()
}

/*
* Expressions statement node
* they are wrappers that consists solely of one expression
 */
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// Node's interface methods
func (exStm *ExpressionStatement) TokenLiteral() string {
	return exStm.Token.Value
}

func (exStm *ExpressionStatement) ToString() string {
	var bf bytes.Buffer
	if exStm.Expression != nil {
		bf.WriteString(exStm.Expression.ToString())
	}
	return bf.String()
}

func (exStm *ExpressionStatement) statementNode() {}

/*
* Integer Literals Node
* they can Occur in different type of expression's
 */
type IntegerLiteral struct {
	Token token.Token
	Value int // we do not specify int size to make it platform independent (32,64)
}

func (intLiteral *IntegerLiteral) TokenLiteral() string {
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
	Token    token.Token // the prefix token
	Operator string
	Right    Expression
}

func (prefixExp *PrefixExpression) TokenLiteral() string {
	return prefixExp.Token.Value
}

func (prefixExp *PrefixExpression) ToString() string {
	var bf bytes.Buffer

	bf.WriteRune('(')
	bf.WriteString(prefixExp.Operator)
	bf.WriteString(prefixExp.Right.ToString())
	bf.WriteRune(')')

	return bf.String()
}

func (prefixExp *PrefixExpression) expressionNode() {}

/*
Infix Expression Nodes
left + right
*/
type InfixExpression struct {
	Token    token.Token // the operator token
	Right    Expression
	Left     Expression
	Operator string
}

func (infixExp *InfixExpression) TokenLiteral() string {
	return infixExp.Token.Value
}

func (infixExp *InfixExpression) ToString() string {
	var bf bytes.Buffer

	bf.WriteRune('(')
	bf.WriteString(infixExp.Left.ToString())
	bf.WriteString(infixExp.Operator)
	bf.WriteString(infixExp.Right.ToString())
	bf.WriteRune(')')

	return bf.String()
}

func (infixExp *InfixExpression) expressionNode() {}

/*
  - If expression Nodes
    implements node and expression Interfaces
*/
type IfExpression struct {
	Token     token.Token // the if token (token.IF)
	Condition Expression
	Body      *BlockStm
	ElseBody  *BlockStm
}

func (ifExp *IfExpression) expressionNode() {}
func (ifExp *IfExpression) TokenLiteral() string {
	return ifExp.Token.Value
}

func (ifExp *IfExpression) ToString() string {
	var bf bytes.Buffer

	bf.WriteString("if")
	bf.WriteString(ifExp.Condition.ToString())
	bf.WriteRune(' ')
	bf.WriteString(ifExp.Body.ToString())

	if ifExp.ElseBody != nil {
		bf.WriteString("else")
		bf.WriteString(ifExp.ElseBody.ToString())
	}

	return bf.String()
}

/*
Block statments Node
implements the node and statment Interfaces
*/
type BlockStm struct {
	Token      token.Token // the { token the starting of if block
	Statements []Statement
}

func (b *BlockStm) statementNode() {}
func (b *BlockStm) TokenLiteral() string {
	return b.Token.Value
}
func (b *BlockStm) ToString() string {

	var bf bytes.Buffer

	bf.WriteRune('{')

	for _, st := range b.Statements {
		bf.WriteString(st.ToString())
	}

	bf.WriteRune('}')

	return bf.String()
}
