package runtime

import (
	"testing"

	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/parser"
	"github.com/houcine7/JIPL/internal/types"
)

func TestIntegerEval(t *testing.T) {
	testData := intTestData
	for _, test := range testData {
		evaluated := getEvaluated(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}


func TestBooleanEval(t *testing.T) {
	testData := boolInputData

	for _, test := range testData {
		evaluated := getEvaluated(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}


func TestReturnEval(t *testing.T) {
	input := "return 10;5454447;"
	evaluated := getEvaluated(input)
	intObj, ok := evaluated.(*types.Integer)
	if !ok {
		t.Fatalf("the obj is not of type *types.Integer, instead got %T",
			evaluated,
		)
	}
	if intObj.Val != 10 {
		t.Fatalf("the value of the integer object is not valid expected :%d and got %d", 10, intObj.Val)
	}
}


func TestIfElseEval(t *testing.T) {
	input := `
	if (10 > 5) {
		return 10;
	} else {
		return 5;
	}
	`
	evaluated := getEvaluated(input)
	intObj, ok := evaluated.(*types.Integer)
	if !ok {
		t.Fatalf("the obj is not of type types.Integer, instead got %T",
			evaluated,
		)
	}
	if intObj.Val != 10 {
		t.Fatalf("the value of the integer object is not valid expected :%d and got %d", 10, intObj.Val)
	}
}



///
func testBooleanObject(t *testing.T, evaluated types.ObjectJIPL, expected bool) {
	boolObj, ok := evaluated.(*types.Boolean)
	if !ok {
		t.Fatalf("the obj is not of type types.Boolean, instead got %T",
			evaluated,
		)
	}
	if boolObj.Val != expected {
		t.Fatalf("the value of the boolean object is not valid expected :%t and got %t", expected, boolObj.Val)
	}
}
func getEvaluated(input string) types.ObjectJIPL {
	l := lexer.InitLexer(input)
	p := parser.InitParser(l)
	program := p.Parse()
	ev,_:=Eval(program)
	return ev 
}

func testIntegerObject(t *testing.T, obj types.ObjectJIPL, expected int) {
	intObj, ok := obj.(*types.Integer)
	if !ok {
		t.Fatalf("the obj is not of type types.Integer, instead got %T",
			obj,
		)
	}
	if intObj.Val != expected {
		t.Fatalf("the value of the integer object is not valid expected :%d and got %d", expected, intObj.Val)
	}
}




// data
var (
	boolInputData = []struct {
		input    string
		expected bool
	}{
		{"true;", true},
		{"false;", false},
		{"1 < 2;", true},
		{"1 > 2;", false},
		{"1 == 2;", false},
		{"1 != 2;", true},
		{"true == true;", true},
		{"true != true;", false},
		{"true == false;", false},
		{"!true;", false},
		{"!false;", true},
	}
	intTestData = []struct {
		input    string
		expected int
	}{
		{"4545;", 4545},
		{"7;", 7},
	}

)