package ast

import (
	"testing"

	"github.com/houcine7/JIPL/internal/token"
)

// test toString method
/*
	for this test we are going to test this statement
	def var1 = var2;
*/
func TestToString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&DefStatement{
				Token: token.NewToken(token.DEF, "def"),
				Name: &Identifier{
					Token: token.NewToken(token.IDENTIFIER, "var1"),
					Value: "var1",
				},
				Value: &Identifier{
					Token: token.NewToken(token.IDENTIFIER, "var2"),
					Value: "var2",
				},
			},
		},
	}

	t.Log(program.ToString())

	if program.ToString() != "def var1 = var2;" {
		t.Errorf("program.String() wrong got=%q", program.ToString())
	}
}
