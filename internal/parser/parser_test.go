package parser

import (
	"fmt"
	"testing"

	ast "github.com/houcine7/JIPL/internal/AST"
	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/parser/data"
)

/*TEST functions*/

// test def statement
func TestDefStatement(t *testing.T) {

	input := data.InputDefStm
	program, parser := getProg(input)

	// check parser errors
	checkParserErrors(parser, t)

	//check is length of the program statement slice is 3
	checkIsProgramStmLengthValid(program, t, 3)

	tests := data.DefStmExpected

	for i, t1 := range tests {

		stm := program.Statements[i]

		if !testDefStatement(t, stm, t1.ExpectedIdentifier, t1.ExpectedValue) {
			return
		}

	}
}

// test return statement
func TestReturnStatement(t *testing.T) {

	input := data.ReturnStm
	expectReturnValue := data.ExpectedReturnValue
	pr, parser := getProg(input)

	checkParserErrors(parser, t)
	checkIsProgramStmLengthValid(pr, t, 3)

	for i, stm := range pr.Statements {
		returnStm, ok := stm.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement not *ast.ReturnStatement type got:%T", stm)
			continue
		}

		if returnStm.TokenLiteral() != "return" {
			t.Errorf("returnStatement token literal is not 'return' instead got: %s",
				returnStm.TokenLiteral())
		}
		if !testLiteralExpression(t, returnStm.ReturnValue, expectReturnValue[i]) {
			return
		}

	}
}

// Expression tests
func TestIdentifiers(t *testing.T) {
	input := data.Identifier

	program, parser := getProg(input)

	checkParserErrors(parser, t)
	// check the length of the program
	checkIsProgramStmLengthValid(program, t, 1)

	stm, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement .got=%T",
			program.Statements[0])
	}

	if testIdentifier(t, stm.Expression, "varName") {
		return
	}

}

// Integer literals test
func TestIntegerLiteral(t *testing.T) {
	input := data.IntegerLit

	pr, parser := getProg(input)

	checkParserErrors(parser, t)
	checkIsProgramStmLengthValid(pr, t, 1)

	stm, ok := pr.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.statement[0] is not of type expressionStatement instead got=%T",
			pr.Statements[0])
	}

	intLiteral, ok := stm.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("stm.Expression is not of type *ast.IntegerLiteral instead got=%T",
			stm.Expression)
	}

	if intLiteral.Value != int(81) {
		t.Errorf("the Integer literal value is not correct, expected=%d instead got=%d",
			81, intLiteral.Value)
	}

	if intLiteral.TokenLiteral() != "81" {
		t.Errorf("the TokenLiteral value is not correct, expected=%s instead got=%s",
			"81", intLiteral.TokenLiteral())
	}
}

// prefix operators
func TestParsePrefixExp(t *testing.T) {
	tests := data.PrefixExpression

	for _, test := range tests {

		pr, parser := getProg(test.Input)

		checkParserErrors(parser, t)

		checkIsProgramStmLengthValid(pr, t, 1)

		stm, ok := pr.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.statement[0] is not of type expressionStatement instead got=%T",
				pr.Statements[0])
		}

		exp, ok := stm.Expression.(*ast.PrefixExpression)

		if !ok {
			t.Fatalf("stm.Expression is not of type PrefixStatement instead got=%T",
				stm.Expression)
		}

		if exp.Operator != test.Operator {
			t.Fatalf("exp Operator is not as expected %s, got=%s",
				test.Operator, exp.Operator)
		}
		if !testLiteralExpression(t, exp.Right, test.IntOperand) {
			return
		}
	}
}

// Test infix Expression
func TestInfixExpression(t *testing.T) {

	testsData := data.InfixExpression

	for _, test := range testsData {
		//
		pr, parser := getProg(test.Input)

		checkParserErrors(parser, t)
		checkIsProgramStmLengthValid(pr, t, 1)

		stm, ok := pr.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("pr.Statements[0] type is not as expected: *ast.Expression, got=%T", pr.Statements[0])
		}

		exp, ok := stm.Expression.(*ast.InfixExpression)

		if !ok {
			t.Fatalf("stm.Expression type is not as expected instead got= %T", stm.Expression)
		}

		if !testLiteralExpression(t, exp.Left, test.Left) ||
			!testLiteralExpression(t, exp.Right, test.Right) {
			return
		}

		if exp.Operator != test.Operator {
			t.Fatalf("Operator is not as expected:%s, but instead got %s", test.Operator,
				exp.Operator,
			)
		}

	}
}

// tests for preceedence
func TestPrecedenceOrderParsing(t *testing.T) {
	tests := data.PrecedenceOrder

	for _, test := range tests {
		prog, parser := getProg(test.Input)
		//checkIsProgramStmLengthValid(pr, t, 1)
		checkParserErrors(parser, t)
		ans := prog.ToString()
		if ans != test.Expected {
			t.Fatalf("wrong restult occurred expected=%s and got=%s", test.Expected, ans)
		}
	}
}

func TestIfExpression(t *testing.T) {
	input := data.IfExpression
	prog, parser := getProg(input)

	checkParserErrors(parser, t)
	checkIsProgramStmLengthValid(prog, t, 1)

	stm, ok := prog.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("pr.Statements[0] type is not as expected: *ast.Expression, got=%T", prog.Statements[0])
	}

	ifExpr, ok := stm.Expression.(*ast.IfExpression)

	if !ok {
		t.Fatalf("stm.Expression type is not as expected: *ast.Expression, got=%T", ifExpr)
	}

	if !testInfixExpression(t, ifExpr.Condition, "m", "n", ">=") {
		return
	}

	if len(ifExpr.Body.Statements) != 1 {
		t.Fatalf("statements of if Expression body is more than expected %d, instead got %d", 1,
			len(ifExpr.Body.Statements))
	}

	body, ok := ifExpr.Body.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("ifExpr.Body.Statements[0] is not ast.ExpressionStatement. got=%T",
			ifExpr.Body.Statements[0])
	}

	if !testInfixExpression(t, body.Expression, "m", 1, "+") {
		return
	}
	elsBody, ok := ifExpr.ElseBody.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("IfStm.ElseBody is not ast.ExpressionStatement. got=%T",
			ifExpr.Body.Statements[0])
	}

	if !testInfixExpression(t, elsBody.Expression, "n", 1, "+") {
		return
	}
}

func TestForLoopFunctions(t *testing.T) {
	input := data.ForLoopTestSimple

	pr, parser := getProg(input)

	checkParserErrors(parser, t)

	checkIsProgramStmLengthValid(pr, t, 1)

	stm, ok := pr.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("pr.Statements[0] type is not *ast.ExpressionStatement instead got %T",
			pr.Statements[0],
		)
	}

	exp, ok := stm.Expression.(*ast.ForLoopExpression)

	if !ok {
		t.Fatalf("*ast.ForLoopExpression type is not *ast.ForLoopExpression instead got %T",
			stm.Expression,
		)
	}

	if !testDefStatement(t, exp.InitStm, "i", 0) {
		return
	}

	if !testInfixExpression(t, exp.Condition, "i", 10, "<=") {
		return
	}

	postFix, ok := exp.PostIteration.(*ast.PostfixExpression)

	if !ok {
		t.Fatalf("exp.PostIteration type is not *ast.PostfixExpression instead got %T",
			exp.PostIteration,
		)
	}

	if postFix.Operator != "++" {
		t.Fatal("the operator of the postIteration is invalid")
	}

	if !testLiteralExpression(t, postFix.Left, "i") {
		return
	}
}

func TestParseFunctions(t *testing.T) {
	input := data.FunctionExp2
	pr, parser := getProg(input)
	checkParserErrors(parser, t)
	checkIsProgramStmLengthValid(pr, t, 1)

	stm, ok := pr.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("pr.Statements[0] type is not *ast.ExpressionStatement instead got %T",
			pr.Statements[0],
		)
	}

	fnExp, ok := stm.Expression.(*ast.FunctionExp)
	if !ok {
		t.Fatalf("stm.Expression type is not *ast.FunctionExp instead got %T",
			stm.Expression,
		)
	}

	testLiteralExpression(t, fnExp.Name, "test")
	testLiteralExpression(t, fnExp.Parameters[0], "pr1")
	testLiteralExpression(t, fnExp.Parameters[1], "pr2")
}

func TestFnCallExpression(t *testing.T) {

	input := data.FunctionCall
	pr, parser := getProg(input)

	checkParserErrors(parser, t)
	//check is length of the program statement slice is 3
	checkIsProgramStmLengthValid(pr, t, 1)

	stm, ok := pr.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("the pr.Stataments[0] is not of type *ast.ExpressionStatement. instead got %T",
			pr.Statements[0])
	}

	fnCall, ok := stm.Expression.(*ast.FunctionCall)

	if !ok {
		t.Fatalf("the stm.Expression is not of type FunctionCall instead got %T",
			stm.Expression)
	}

	testLiteralExpression(t, fnCall.Function, "functionName")
	for i, argExp := range fnCall.Arguments {
		testLiteralExpression(t, argExp, fmt.Sprintf("arg%d", i+1))
	}

}

func TestAssignExpr(t *testing.T) {
	testData := data.AssignExp
	pr, parser := getProg(testData.Input)

	// check for parsing errors
	checkParserErrors(parser, t)

	// check for program length
	checkIsProgramStmLengthValid(pr, t, 1)

	stm, ok := pr.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("the pr.Stataments[0] is not of type *ast.ExpressionStatement. instead got %T",
			pr.Statements[0])
	}

	assignExp, ok := stm.Expression.(*ast.AssignmentExpression)

	if !ok {
		t.Fatalf("the stm.Expression is not not of type *ast.AssignmentExpression. instead got %T",
			stm.Expression)
	}

	testIdentifier(t, assignExp.Left, testData.ExpectedIden)
	testIntegerLiteral(t, assignExp.AssignmentValue, testData.ExpectedVal)

	fmt.Println("------------ assign expression String ------------")
	fmt.Println(assignExp.ToString())

}

func TestParseArraysLit(t *testing.T) {
	input := data.Arrays
	pr, parser := getProg(input)

	checkParserErrors(parser, t)
	checkIsProgramStmLengthValid(pr, t, 1)

	stm, ok := pr.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("the pr.stamtemtns[0] is not of type *ast.Expression instead go %T", stm)
	}

	arrLit, ok := stm.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("the stm.Expression is not of type *ast.ArrayList, instead got %T", arrLit)
	}

	testIntegerLiteral(t, arrLit.Values[0], 1)
	testInfixExpression(t, arrLit.Values[1], 12, 8, "-")
	testIntegerLiteral(t, arrLit.Values[2], 7)

}

func TestArrayIndexExp(t *testing.T) {
	input := data.ArrayIndex

	pr, parser := getProg(input)

	checkParserErrors(parser, t)
	checkIsProgramStmLengthValid(pr, t, 1)

	stm, ok := pr.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("the pr.stamtemtns[0] is not of type *ast.Expression instead go %T", stm)
	}

	arrIdx, ok := stm.Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("the stm.Expression is not of type *ast.IndexExpression instead go %T", arrIdx)
	}

	if !testIdentifier(t, arrIdx.Left, "nums") {
		return
	}

	if !testInfixExpression(t, arrIdx.Index, 7, 4, "-") {
		return
	}
}

func TestClassExpression(t *testing.T) {
	input := data.ClassExp

	pr, parser := getProg(input)

	checkParserErrors(parser, t)
	checkIsProgramStmLengthValid(pr, t, 1)

	stm, ok := pr.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("pr.Statements[0] is not of type *ast.ExpressionStatements instead got %T ", stm)
	}

	fmt.Printf("program is %s", pr.ToString())
}

// Tests helper functions
func checkIsProgramStmLengthValid(program *ast.Program, t *testing.T, length int) {
	if len(program.Statements) != length {
		t.Fatalf("the program.Statements doesn't not contain %d  statements, instead we got %d",
			length, len(program.Statements))
	}
}

func checkParserErrors(p *Parser, t *testing.T) {
	errors := p.Errors()

	if len(errors) == 0 {
		t.Log("INFO: no ERRORS OCCURRED")
		return
	}

	//
	t.Log("-------PARSING ERRORS: --------")
	t.Errorf("%d error found on the parser", len(errors))

	for i, err := range errors {
		t.Errorf("Parser index:%d has message %s", i, err.Message)
	}

	t.FailNow() // mark tests as failed and stop execution
}

func testDefStatement(t *testing.T, stm ast.Statement, name string, val int) bool {
	if stm.TokenLiteral() != "def" {
		t.Errorf("s.tokenLiteral is not 'def'. got instead:%q", stm.TokenLiteral())
		return false
	}

	defStm, ok := stm.(*ast.DefStatement) // type assertion (casting if stm is of type *ast.Statement)

	if !ok {
		t.Errorf("s not *ast.DefStatement. got=%T", stm)
		return false
	}

	if defStm.Name.Value != name {
		t.Errorf("def Statement Name.Value not '%s'. got=%s", name, defStm.Name.Value)
		return false
	}

	if defStm.Name.Value != name {
		t.Errorf("def Statement Name.Value not '%s'. got=%s", name, defStm.Name.Value)
		return false
	}

	if !testLiteralExpression(t, defStm.Value, val) {
		return false
	}

	return true
}

func testIntegerLiteral(t *testing.T, intLit ast.Expression, value interface{}) bool {
	intVal, ok := intLit.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il type is not as expected *ast.IntegerLiteral. got=%T", intLit)
		return false
	}

	if intVal.Value != value {
		t.Errorf("intVal.val is not as expected %d. instead got %d", value, intVal.Value)
		return false
	}

	if intVal.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("intVal.TokenLiteral not equal to %d instead got %s", value, intVal.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {

	ident, ok := exp.(*ast.Identifier)

	if !ok {
		t.Fatalf("expression is not of type *ast.Identifier instead, got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value expected=%s, and got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral is not %s. instead got=%s", "foobar",
			ident.TokenLiteral())
		return false
	}
	return true
}

func testBoolean(t *testing.T, exp ast.Expression, val interface{}) bool {
	boolexp, ok := exp.(*ast.BooleanExp)

	if !ok {
		t.Fatalf("expression is not of type *ast.BooleanExp instead, got=%T", exp)
		return false
	}

	if boolexp.Value != val {
		t.Errorf("boolexp.Value expected=%t, and got=%t", val, boolexp.Value)
		return false
	}

	if boolexp.TokenLiteral() != fmt.Sprintf("%t", val) {
		t.Errorf("ident.TokenLiteral is not %s. instead got=%s", "foobar",
			boolexp.TokenLiteral())
		return false
	}

	return true
}

func getProg(input string) (*ast.Program, *Parser) {
	lexer := lexer.InitLexer(input)
	parser := InitParser(lexer)
	pr := parser.Parse()

	return pr, parser
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch val := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, val)
	case string:
		return testIdentifier(t, exp, val)
	case bool:
		return testBoolean(t, exp, val)
	default:
		t.Fatalf("type of expression %T not handled", val)
		return false
	}
}

func testInfixExpression(t *testing.T, exp ast.Expression, left interface{},
	right interface{}, operator string) bool {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not of type ast.Expression instead got:%T", opExp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp operator is not %s, instead got %s", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true

}
