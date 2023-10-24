package lexer

import (
	"log"
	"testing"

	"github.com/houcine7/JIPL/token"
)

// MOCK1_TEST := "=+(){},;"

func TestNextToken(t *testing.T) {
	MOCK1_TEST := "=+(){},;"

	var tests = []struct {
		expectedTokenType token.TokenType
		expectedValue string
	}{
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.PLUS, expectedValue: "+"},
		{expectedTokenType: token.LP, expectedValue: "("},
		{expectedTokenType: token.RP, expectedValue: ")"},
		{expectedTokenType: token.LCB, expectedValue: "{"},
		{expectedTokenType: token.RCB, expectedValue: "}"},
		{expectedTokenType: token.CAMMA, expectedValue: ","},
		{expectedTokenType: token.SCOLON, expectedValue: ";"},
	}

	myLexer := New(MOCK1_TEST)
	

	for i,et :=range(tests) {
		calculatedToken:= myLexer.NextToken()
		log.Print(calculatedToken)
		// test the token type 

		if et.expectedTokenType !=  calculatedToken.Type {
			log.Fatalf("tests index %d -> tokenType wrong, expected:[%q] and got:[%q]",
			i,et.expectedTokenType,calculatedToken.Type)
		}

		// test the token literal value
		if et.expectedValue != calculatedToken.Value {
			log.Fatalf("tests index %d -> token value is wrong, expected:[%q] and got:[%q]",
			i,et.expectedValue,calculatedToken.Value)
		}
	}
}
