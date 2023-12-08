package parser

import (
	"fmt"
	"testing"

	ast "github.com/houcine7/JIPL/AST"
	"github.com/houcine7/JIPL/lexer"
)

func TestDefStatement(t *testing.T) {

	input := `
def num1 = 5;
def num2 = 10;
def foobar = 838383;
`
	l := lexer.InitLexer(input)
	fmt.Println("-----------------------")
	fmt.Println("lexer", l)
	parser := InitParser(l)
	fmt.Println("-----------------------")
	fmt.Println("parser",parser)

	program := parser.Parse()

	fmt.Println(program)

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

