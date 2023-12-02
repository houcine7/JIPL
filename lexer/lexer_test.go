package lexer

import (
	"log"
	"testing"

	"github.com/houcine7/JIPL/token"
)

//MOCK1_TEST := "=+(){},;"

//basic test

func TestNextToken(t *testing.T) {
	MOCK1_TEST := "=+(){},;"

	var tests = []struct {
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
		{expectedTokenType: token.SCOLON, expectedValue: ";"},
	}

	myLexer := InitLexer(MOCK1_TEST)

	for i, et := range tests {

		calculatedToken := myLexer.NextToken()
		
		log.Print(calculatedToken)
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


func TestNextToken2(t *testing.T) {
	MOCK1_TEST := `def val1 = 30;
	def val2 = 3;
	def add = function(x, y) {
		x + y;
	};
	def result = add(val1, val2);`

	var tests = []struct {
		expectedTokenType token.TokenType
		expectedValue string
	}{
		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "30"},
		{expectedTokenType: token.SCOLON, expectedValue: ";"},

		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "3"},
		{expectedTokenType: token.SCOLON, expectedValue: ";"},

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
		{expectedTokenType: token.SCOLON, expectedValue: ";"},
		{expectedTokenType: token.RCB, expectedValue: "}"},
		{expectedTokenType: token.SCOLON, expectedValue: ";"},
		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "result"},
		{expectedTokenType: token.ASSIGN , expectedValue: "="},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "add"},
		{expectedTokenType: token.LP,expectedValue: "("},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.COMMA, expectedValue: ","},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.RP,expectedValue: ")"},
		{expectedTokenType: token.SCOLON,expectedValue: ";"},
		
	}
	myLexer := InitLexer(MOCK1_TEST)


	for i,et :=range(tests) {
		calculatedToken:= myLexer.NextToken()
		log.Println("------- TOKEN ---------")
		log.Println(calculatedToken,calculatedToken.Type,calculatedToken.Value)
		log.Println("----------------")
		// test the token type

		if et.expectedTokenType !=  calculatedToken.Type {
			log.Fatalf("tests index %d -> tokenType wrong, expected:[%d] and got:[%d]",
			i,et.expectedTokenType,calculatedToken.Type)
		}

		// test the token literal value
		if et.expectedValue != calculatedToken.Value {
			log.Fatalf("tests index %d -> token value is wrong, expected:[%q] and got:[%q]",
			i,et.expectedValue,calculatedToken.Value)
		}
	}
}



func TestNextToken3(t *testing.T){
	MOCK1_TEST := `def val1 = 30;
	
	def val2 = 3;
	def add = function(x, y) {
		x + y;
	};
	
	def result = add(val1, val2);
	
	!-*7/;

	5 < 10;
	`

	var tests = []struct {
		expectedTokenType token.TokenType
		expectedValue string
	}{
		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "30"},
		{expectedTokenType: token.SCOLON, expectedValue: ";"},

		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.ASSIGN, expectedValue: "="},
		{expectedTokenType: token.INT, expectedValue: "3"},
		{expectedTokenType: token.SCOLON, expectedValue: ";"},

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
		{expectedTokenType: token.SCOLON, expectedValue: ";"},
		{expectedTokenType: token.RCB, expectedValue: "}"},
		{expectedTokenType: token.SCOLON, expectedValue: ";"},

		{expectedTokenType: token.DEF, expectedValue: "def"},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "result"},
		{expectedTokenType: token.ASSIGN , expectedValue: "="},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "add"},
		{expectedTokenType: token.LP,expectedValue: "("},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val1"},
		{expectedTokenType: token.COMMA, expectedValue: ","},
		{expectedTokenType: token.IDENTIFIER, expectedValue: "val2"},
		{expectedTokenType: token.RP,expectedValue: ")"},
		{expectedTokenType: token.SCOLON,expectedValue: ";"},

		{expectedTokenType: token.EXMARK, expectedValue: "!"},
		{expectedTokenType: token.MINUS, expectedValue: "-"},
		{expectedTokenType: token.STAR,expectedValue: "*"},
		{expectedTokenType: token.INT, expectedValue: "7"},
		{expectedTokenType: token.SLASH,expectedValue: "/"},
		{expectedTokenType: token.SCOLON,expectedValue: ";"},

		{expectedTokenType:  token.INT, expectedValue: "5"},
		{expectedTokenType: token.LT, expectedValue: "<"},
		{expectedTokenType: token.INT, expectedValue: "10"},
		{expectedTokenType: token.SCOLON,expectedValue: ";"},
	}

	myLexer := InitLexer(MOCK1_TEST)


	for i,et :=range(tests) {
		calculatedToken:= myLexer.NextToken()
		log.Println("------- TOKEN ---------")
		log.Println(calculatedToken,calculatedToken.Type,calculatedToken.Value)
		log.Println("----------------")
		// test the token type

		if et.expectedTokenType !=  calculatedToken.Type {
			log.Fatalf("tests index %d -> tokenType wrong, expected:[%d] and got:[%d]",
			i,et.expectedTokenType,calculatedToken.Type)
		}

		// test the token literal value
		if et.expectedValue != calculatedToken.Value {
			log.Fatalf("tests index %d -> token value is wrong, expected:[%q] and got:[%q]",
			i,et.expectedValue,calculatedToken.Value)
		}
	}
}





// func TestReadIdentifier(t *testing.T){
	
// 	MOCKS := []string{"def","1ref","test1","test4t"}
// 	EXPECTED :=[]string{"def","","test1","test4t"}

// 	for i :=range MOCKS {
// 		lexer := InitLexer(MOCKS[i])
// 		res :=lexer.ReadIdentifier()
// 		log.Print(res)
// 		if res != EXPECTED[i] {
// 			log.Fatalf("Wrong result expected: %q and got : %q", EXPECTED[i], res)
// 		}
// 	}
// }

