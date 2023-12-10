package parser

import (
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

func checkIsProgramStmLengthValid(program *ast.Program,t *testing.T,length int){
	if len(program.Statements) !=length {
		t.Fatalf("the program.Statements doesn't not contain 3 statements, instead we got %d",
		len(program.Statements))
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


