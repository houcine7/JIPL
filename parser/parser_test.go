package parser

import (
	"testing"

	ast "github.com/houcine7/JIPL/AST"
	"github.com/houcine7/JIPL/lexer"
)

func TestDefStatement(t *testing.T) {

	input := `
def num1 = 13;
def num2 = 0;
def foobar = 5321;
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

	if program==nil{
		t.Fatalf("parse returned a nil value")
	}

	if len(program.Statements) !=3 {
		t.Fatalf("the program.Statements doesn't not contain 3 statements, instead we got %d",
		len(program.Statements))
	}

	tests := []struct{
		expectedIdentifier string
	}{
		{"num1"},
		{"num2"},
		{"foobar"},
	}

	for i,t1 := range tests{
		
		stm := program.Statements[i]
		
		if !testDefStatement(t,stm,t1.expectedIdentifier){
			return
		}
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

	defStm, ok := stm.(*ast.DefStatement)

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


