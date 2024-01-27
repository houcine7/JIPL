package ast

import (
	"bytes"

	"github.com/houcine7/JIPL/internal/token"
)

/*
* This node is the root of each AST tree
* A program is a set of statements
 */
type Program struct {
	Statements []Statement
}

type DefStatement struct {
	Token token.Token // toke.DEF token
	Name  *Identifier
	Value Expression
}

type Identifier struct { //
	Token token.Token // token.IDENTIFIER token
	Value string
}

type ReturnStatement struct {
	Token       token.Token // the token is "return"
	ReturnValue Expression
}

type BooleanExp struct {
	Token token.Token
	Value bool // the boolean value corresponds to bool
}

type FunctionExp struct {
	Token      token.Token   // the function token used to represent functions
	Name       *Identifier   // the name of the functoin
	Parameters []*Identifier // function parmas
	FnBody     *BlockStm     // function body
}

type FunctionCall struct {
	Token     token.Token  // token '(' LP AST node constructs in infix pos fun()
	Function  Expression   // identifier
	Arguments []Expression // function call args
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type IntegerLiteral struct {
	Token token.Token
	Value int
}

type PrefixExpression struct {
	Token    token.Token // the prefix token
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token // the operator token
	Right    Expression
	Left     Expression
	Operator string
}

type ForLoopExpression struct {
	Token         token.Token // the 'for' token idencate for loop starting point
	InitStm       Statement   // the initializaiton stm
	Condition     Expression  // loop condition
	PostIteration Expression  // the post iteration expression
	Body          *BlockStm   // loop body that would be executed
}

type PostfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
}

type IfExpression struct {
	Token     token.Token // the if token (token.IF)
	Condition Expression
	Body      *BlockStm
	ElseBody  *BlockStm
}

type BlockStm struct {
	Token      token.Token // the { token the starting of if block
	Statements []Statement
}

type AssignmentExpression struct {
	Token            token.Token
	Left             *Identifier
	AssignmentValue Expression
}

type ArrayLiteral struct {
	Token  token.Token //  the [ token starting the arrayLiteral
	Values []Expression
}

type IndexExpression struct {
	Token token.Token // [ token
	Left  Expression
	Index Expression
}

// Node implementation
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
func (defStm *DefStatement) TokenLiteral() string {
	return defStm.Token.Value
}

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

func (postfixExp *PostfixExpression) TokenLiteral() string {
	return postfixExp.Token.Value
}

func (postfixExp *PostfixExpression) ToString() string {
	var bf bytes.Buffer
	bf.WriteRune('(')
	bf.WriteString(postfixExp.Left.ToString())
	bf.WriteString(postfixExp.Operator)
	bf.WriteRune(')')
	return bf.String()
}
func (forExp *ForLoopExpression) TokenLiteral() string {
	return forExp.Token.Value
}
func (forExp *ForLoopExpression) ToString() string {

	var bf bytes.Buffer
	bf.WriteString(forExp.TokenLiteral())
	bf.WriteString(" (")
	bf.WriteString(forExp.InitStm.ToString())
	bf.WriteString("; ")
	bf.WriteString(forExp.Condition.ToString())
	bf.WriteString("; ")
	bf.WriteString(forExp.PostIteration.ToString())
	bf.WriteString(" )")

	bf.WriteString(forExp.Body.ToString())
	return bf.String()
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

func (strLit *StringLiteral) TokenLiteral() string {
	return strLit.Token.Value
}
func (strLit *StringLiteral) ToString() string {
	return strLit.Token.Value
}
func (intLiteral *IntegerLiteral) TokenLiteral() string {
	return intLiteral.Token.Value
}

func (intLiteral *IntegerLiteral) ToString() string {
	return intLiteral.Token.Value
}

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
func (fnExp *FunctionExp) TokenLiteral() string {
	return fnExp.Token.Value
}

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
func (b *BooleanExp) TokenLiteral() string {
	return b.Token.Value
}

func (b *BooleanExp) ToString() string {
	return b.Token.Value
}
func (reStm *ReturnStatement) TokenLiteral() string {
	return reStm.Token.Value
}

func (resStm *ReturnStatement) ToString() string {
	var bf bytes.Buffer
	bf.WriteString(resStm.TokenLiteral())
	if resStm.ReturnValue != nil {
		bf.WriteRune(' ')
		bf.WriteString(resStm.ReturnValue.ToString())
	}
	bf.WriteString(";")
	return bf.String()
}
func (ident *Identifier) TokenLiteral() string {
	return ident.Token.Value
}

func (ident *Identifier) ToString() string {
	return ident.Value
}

func (assignExpr *AssignmentExpression) TokenLiteral() string {
	return assignExpr.Token.Value
}

func (assignExpr *AssignmentExpression) ToString() string {

	var bf bytes.Buffer
	bf.WriteString(assignExpr.Left.ToString())
	bf.WriteString(" = ")
	bf.WriteString(assignExpr.AssignmentValue.ToString())

	return bf.String()
}

func (arr *ArrayLiteral) TokenLiteral() string {
	return arr.Token.Value
}

func (arr *ArrayLiteral) ToString() string {
	var bf bytes.Buffer
	bf.WriteRune('[')
	for i, exp := range arr.Values {
		bf.WriteString(exp.ToString())
		if i != len(arr.Values)-1 {
			bf.WriteRune(',')
		}
	}
	return bf.String()
}

func (indexExp *IndexExpression) ToString() string {
	var bf bytes.Buffer
	bf.WriteString(indexExp.Left.ToString())
	bf.WriteRune('[')
	bf.WriteString(indexExp.Index.ToString())
	bf.WriteRune(']')
	return bf.String()
}
func (indexExp *IndexExpression) TokenLiteral() string {
	return indexExp.Token.Value
}

// expression implementaions
func (postfixExp *PostfixExpression) expressionNode()     {}
func (forExp *ForLoopExpression) expressionNode()         {}
func (infixExp *InfixExpression) expressionNode()         {}
func (prefixExp *PrefixExpression) expressionNode()       {}
func (strLit *StringLiteral) expressionNode()             {}
func (intLiteral *IntegerLiteral) expressionNode()        {}
func (fnCall *FunctionCall) expressionNode()              {}
func (fnExp *FunctionExp) expressionNode()                {}
func (b *BooleanExp) expressionNode()                     {}
func (ident *Identifier) expressionNode()                 {}
func (assignExpr *AssignmentExpression) expressionNode() {}
func (arr *ArrayLiteral) expressionNode()                 {}
func (indexExp *IndexExpression) expressionNode()         {}

// statemetns implmentations
func (b *BlockStm) statementNode()                {}
func (defStm *DefStatement) statementNode()       {}
func (exStm *ExpressionStatement) statementNode() {}
func (reStm *ReturnStatement) statementNode()     {}
