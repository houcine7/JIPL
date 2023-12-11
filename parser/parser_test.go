package parser

import (
	"fmt"
	"testing"

	ast "github.com/houcine7/JIPL/AST"
	"github.com/houcine7/JIPL/lexer"
)

/*TEST functions*/

// test def statement
func TestDefStatement(t *testing.T) {

	input := `
def num1 = 13;
def  total= 0;
def a= 5321;
`
	l := lexer.InitLexer(input)
	
	t.Log("-----------------------")
	t.Log("lexer", l)
	parser := InitParser(l)
	t.Log("-----------------------")
	t.Log("parser",parser)

	program := parser.Parse()

	// check parser errors 
	checkParserErrors(parser,t)

	//check is length of the program statement slice is 3 
	checkIsProgramStmLengthValid(program,t,3)

	tests := []struct{
		expectedIdentifier string
	}{
		{"num1"},
		{"total"},
		{"a"},
	}

	for i,t1 := range tests{
		
		stm := program.Statements[i]
		
		if !testDefStatement(t,stm,t1.expectedIdentifier){
			return
		}
	}
}

// test return statement 
func TestReturnStatement(t *testing.T){
	input :=`
	return 545;
	return 101232;
	return 0;
	`

	l :=lexer.InitLexer(input)
	parser := InitParser(l)

	pr := parser.Parse()
	
	//check is length of the program statement slice is 3 
	checkIsProgramStmLengthValid(pr, t,3)

	for _,stm := range pr.Statements {
		returnStm,ok := stm.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("statement not *ast.ReturnStatement type got:%T",stm)
			continue
		}

		if returnStm.TokenLiteral() !="return" {
			t.Errorf("returnStatement token literal is not 'return' instead got: %s",
			returnStm.TokenLiteral())
		}

	}
}

// Expression tests
func TestIdentifier(t *testing.T){
	input :=`varName;`

	lexer :=lexer.InitLexer(input)
	parser := InitParser(lexer)
	program := parser.Parse()

	checkParserErrors(parser,t)
	// check the length of the program
	checkIsProgramStmLengthValid(program,t,1)

	stm,ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement .got=%T",
	program.Statements[0])
	}

	ident,ok := stm.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("Expression of type *ast.Identifier instead, got=%T",stm.Expression)
	}

	if ident.Value !="varName" {
		t.Errorf("ident.Value expected=%s, and got=%s","varName",ident.Value)
	}

	if ident.TokenLiteral() !="varName"{
		t.Errorf("ident.TokenLiteral is not %s. instead got=%s","foobar",
		ident.TokenLiteral())
	}
}


// Integer literals test
func TestIntegerLiteral(t *testing.T){
	input :="81;"

	lexer := lexer.InitLexer(input)
	parser := InitParser(lexer)
	
	pr :=parser.Parse()
	checkParserErrors(parser,t)
	checkIsProgramStmLengthValid(pr,t,1)

	stm,ok := pr.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.statement[0] is not of type expressionStatement instead got=%T",
	pr.Statements[0])
	}

	intLiteral,ok :=stm.Expression.(*ast.IntegerLiteral)
	if !ok{
		t.Fatalf("stm.Expression is not of type *ast.IntegerLiteral instead got=%T",
	stm.Expression)
	}

	if intLiteral.Value != int(81){
		t.Errorf("the Integer literal value is not correct, expected=%d instead got=%d",
	81,intLiteral.Value)
	}

	if intLiteral.TokenLiteral() !="81" {
		t.Errorf("the TokenLiteral value is not correct, expected=%s instead got=%s",
	"81",intLiteral.TokenLiteral())
	}
}

// prefix operators 
func TestParsePrefixExp(t *testing.T){
	tests := []struct{
		input string
		operator string
		intOperand int
	}{
		{input :"!7;", operator :"!", intOperand: 7},
		{input :"-42;", operator :"-", intOperand: 42},
	}

	for _,test :=range tests {
		l :=lexer.InitLexer(test.input)
		parser := InitParser(l)
		fmt.Println("----------- Parser ---------",l)
		pr := parser.Parse()

		checkParserErrors(parser,t)

		fmt.Println("Length of statements ", len(pr.Statements))

		checkIsProgramStmLengthValid(pr,t,1)

		stm,ok := pr.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.statement[0] is not of type expressionStatement instead got=%T",
			pr.Statements[0])
		}

		exp,ok := stm.Expression.(*ast.PrefixExpression) 

		if !ok{
			t.Fatalf("stm.Expression is not of type PrefixStatement instead got=%T",
			stm.Expression)
		}

		if exp.Operator != test.operator {
			t.Fatalf("exp Operator is not as expected %s, got=%s",
			test.operator,exp.Operator)
		}
		if !testIntegerLiteral(t,exp.Right,test.intOperand){
			return 
		}
	}
}


// Tests helper functions 
func checkIsProgramStmLengthValid(program *ast.Program,t *testing.T,length int){
	if len(program.Statements) !=length {
		t.Fatalf("the program.Statements doesn't not contain %d  statements, instead we got %d",
		length,len(program.Statements))
	}
}

func checkParserErrors(p *Parser,t *testing.T){
	errors := p.Errors()

	if len(errors) ==0 {
		t.Log("INFO: no ERRORS OCCURRED")
		return
	}

	//
	t.Log("-------PARSING ERRORS: --------")
	t.Errorf("%d error found on the parser",len(errors))

	for i,msg := range errors  {
		t.Errorf("Parser index:%d has message %s",i,msg)
	}

	t.FailNow() // mark tests as failed and stop execution 
} 

func testDefStatement(t *testing.T, stm ast.Statement, name string) bool {
	if stm.TokenLiteral() !="def" {
		t.Errorf("s.tokenLiteral is not 'def'. got instead:%q",stm.TokenLiteral())
		return false
	}

	defStm, ok := stm.(*ast.DefStatement) // type assertion (casting if stm is of type *ast.Statement)

	if !ok {
		t.Errorf("s not *ast.DefStatement. got=%T", stm)
		return false
	}

	if defStm.Name.Value !=name{
		t.Errorf("def Statement Name.Value not '%s'. got=%s",name,defStm.Name.Value)
		return false
	}
	
	if defStm.Name.Value !=name{
		t.Errorf("def Statement Name.Value not '%s'. got=%s",name,defStm.Name.Value)
		return false;
	}
	
	return true;
}

func testIntegerLiteral(t *testing.T, il ast.Expression , value int) bool{
	intVal,ok :=il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il type is not as expected *ast.IntegerLiteral. got=%T",il)
		return false
	}

	if intVal.Value !=value {
		t.Errorf("intVal.val is not as expected %d. instead got %d",value,intVal.Value)
		return false
	}
	
	if intVal.TokenLiteral() != fmt.Sprintf("%d",value){
		t.Errorf("intVal.TokenLiteral not equal to %d instead got %s",value,intVal.TokenLiteral())
		return false
	}
	
	return true
}

