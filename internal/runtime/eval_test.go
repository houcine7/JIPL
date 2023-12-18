package runtime

import (
	"testing"

	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/parser"
	"github.com/houcine7/JIPL/internal/types"
)

func TestIntegerEval(t *testing.T) {
	testData := []struct {
		input    string
		expected int
	}{
		{"4545;", 4545},
		{"7;", 7},
	}

	for _, test := range testData {
		evaluated := getEvaluated(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func getEvaluated(input string) types.ObjectJIPL {
	l := lexer.InitLexer(input)
	p := parser.InitParser(l)
	program := p.Parse()
	return Eval(program)
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
