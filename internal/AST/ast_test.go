package ast

import (
	"testing"

	"github.com/houcine7/JIPL/internal/token"
)

/*
def var1 = var2;
*/
func TestToString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&DefStatement{
				Token: token.CreateToken(token.DEF, "def"),
				Name: &Identifier{
					Token: token.CreateToken(token.IDENTIFIER, "var1"),
					Value: "var1",
				},
				Value: &Identifier{
					Token: token.CreateToken(token.IDENTIFIER, "var2"),
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
