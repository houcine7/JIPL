package lexer

import (
	"log"
	"testing"

	"github.com/houcine7/JIPL/internal/token"
)

func TestNext(t *testing.T) {
	MOCK1_TEST := Mock0

	var tests = NextData0
	myLexer := InitLexer(MOCK1_TEST)

	for i, et := range tests {

		calculatedToken := myLexer.NextToken()

		// test the token type
		if et.expectedTokenType != calculatedToken.Type {
			log.Fatalf("tests index %d -> tokenType wrong, expected:[%q] and got:[%q]",
				i, et.expectedTokenType, calculatedToken.Type)
		}

		// test the token literal value
		if et.expectedValue != calculatedToken.Value {
			log.Fatalf("tests index %d -> token value is wrong, expected:[%q] and got:[%q]",
				i, et.expectedValue, calculatedToken.Value)
		}
	}
}

func TestNext2(t *testing.T) {
	MOCK1_TEST := Mock1

	var tests = NextData1
	myLexer := InitLexer(MOCK1_TEST)

	for i, et := range tests {
		calculatedToken := myLexer.NextToken()
		// test the token type
		if et.expectedTokenType != calculatedToken.Type {
			log.Fatalf("tests index %d -> tokenType wrong, expected:[%d] and got:[%d]",
				i, et.expectedTokenType, calculatedToken.Type)
		}

		// test the token literal value
		if et.expectedValue != calculatedToken.Value {
			log.Fatalf("tests index %d -> token value is wrong, expected:[%q] and got:[%q]",
				i, et.expectedValue, calculatedToken.Value)
		}
	}
}

func TestNext3(t *testing.T) {
	MOCK1_TEST := Mock2

	var tests = NextTestData

	myLexer := InitLexer(MOCK1_TEST)

	for i, et := range tests {
		calculatedToken := myLexer.NextToken()
		// test the token type

		if et.expectedTokenType != calculatedToken.Type {
			log.Fatalf("tests index %d -> tokenType wrong, expected:[%d] and got:[%d]",
				i, et.expectedTokenType, calculatedToken.Type)
		}

		// test the token literal value
		if et.expectedValue != calculatedToken.Value {
			log.Fatalf("tests index %d -> token value is wrong, expected:[%q] and got:[%q]",
				i, et.expectedValue, calculatedToken.Value)
		}
	}
}

// Test data
var (
	NextTestData = []struct {
		expectedTokenType token.TokenType
		expectedValue     string
	}{
		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "30"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "3"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "add"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.FUNCTION, expectedValue: "function"},
		{expectedTokenType: token.LP, expectedValue: "("},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "x"},
		{expectedTokenType: token.COMMA, expectedValue: ","},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "y"},
		{expectedTokenType: token.RP, expectedValue: ")"},
		{expectedTokenType: token.LCB, expectedValue: "{"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "x"},
		{expectedTokenType: token.PLUS, expectedValue: "+"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "y"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
		{expectedTokenType: token.RETURN, expectedValue: "return"},
		{expectedTokenType: token.TRUE, expectedValue: "true"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
		{expectedTokenType: token.RCB, expectedValue: "}"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "result"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "add"},
		{expectedTokenType: token.LP, expectedValue: "("},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.COMMA, expectedValue: ","},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.RP, expectedValue: ")"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.EX_MARK, expectedValue: "!"},
		{expectedTokenType: token.MINUS, expectedValue: "-"},
		{expectedTokenType: token.STAR, expectedValue: "*"},
		{expectedTokenType: token.INT, expectedValue: "7"},
		{expectedTokenType: token.SLASH, expectedValue: "/"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.IF, expectedValue: "if"},
		{expectedTokenType: token.LP, expectedValue: "("},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.LT, expectedValue: "<"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.RP, expectedValue: ")"},
		{expectedTokenType: token.LCB, expectedValue: "{"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "7777"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
		{expectedTokenType: token.RCB, expectedValue: "}"},

		{expectedTokenType: token.ELSE, expectedValue: "else"},
		{expectedTokenType: token.LCB, expectedValue: "{"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "7777"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
		{expectedTokenType: token.RCB, expectedValue: "}"},

		{expectedTokenType: token.INT, expectedValue: "10"},
		{expectedTokenType: token.EQUAL, expectedValue: "=="},
		{expectedTokenType: token.INT, expectedValue: "10"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.INT, expectedValue: "10"},
		{expectedTokenType: token.NOT_EQUAL, expectedValue: "!="},
		{expectedTokenType: token.INT, expectedValue: "7"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.INT, expectedValue: "10"},
		{expectedTokenType: token.LT_OR_EQ, expectedValue: "<="},
		{expectedTokenType: token.INT, expectedValue: "20"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.INT, expectedValue: "10"},
		{expectedTokenType: token.GT_OR_EQ, expectedValue: ">="},
		{expectedTokenType: token.INT, expectedValue: "0"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.INT, expectedValue: "10"},
		{expectedTokenType: token.INCREMENT, expectedValue: "++"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
		{expectedTokenType: token.INT, expectedValue: "10"},
		{expectedTokenType: token.DECREMENT, expectedValue: "--"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.RB, expectedValue: "["},
		{expectedTokenType: token.INT, expectedValue: "2222"},
		{expectedTokenType: token.COMMA, expectedValue: ","},
		{expectedTokenType: token.INT, expectedValue: "7777"},
		{expectedTokenType: token.LB, expectedValue: "]"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
	}

	NextData0 = []struct {
		expectedTokenType token.TokenType
		expectedValue     string
	}{
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.PLUS, expectedValue: "+"},
		{expectedTokenType: token.LP, expectedValue: "("},
		{expectedTokenType: token.RP, expectedValue: ")"},
		{expectedTokenType: token.LCB, expectedValue: "{"},
		{expectedTokenType: token.RCB, expectedValue: "}"},
		{expectedTokenType: token.COMMA, expectedValue: ","},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
	}

	NextData1 = []struct {
		expectedTokenType token.TokenType
		expectedValue     string
	}{
		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "30"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "3"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},

		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "add"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.FUNCTION, expectedValue: "function"},
		{expectedTokenType: token.LP, expectedValue: "("},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "x"},
		{expectedTokenType: token.COMMA, expectedValue: ","},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "y"},
		{expectedTokenType: token.RP, expectedValue: ")"},
		{expectedTokenType: token.LCB, expectedValue: "{"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "x"},
		{expectedTokenType: token.PLUS, expectedValue: "+"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "y"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
		{expectedTokenType: token.RCB, expectedValue: "}"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "result"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "add"},
		{expectedTokenType: token.LP, expectedValue: "("},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.COMMA, expectedValue: ","},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.RP, expectedValue: ")"},
		{expectedTokenType: token.S_COLON, expectedValue: ";"},
	}

	Mock2 = `def val1 = 30;
	
	def val2 = 3;
	def add = function(x, y) {
		x + y;
		return true;
	};
	
	def result = add(val1, val2);
	
	!-*7/;

	if(val1 < val2) {
		val1 = 7777;
	} else {
		val2=7777;
	}
	10 == 10;
	10 != 7;
	10 <= 20;
	10 >= 0;
	10++;
	10--;
	[2222,7777];
	`

	Mock1 = `def val1 = 30;
	def val2 = 3;
	def add = function(x, y) {
		x + y;
	};
	def result = add(val1, val2);`

	Mock0 = "=+(){},;"
)
