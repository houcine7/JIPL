package lexer

import (
	"fmt"
	"unicode/utf8"

	"github.com/houcine7/JIPL/token"
	"github.com/houcine7/JIPL/utils"
)

type Lexer struct {
	input      string // the string to tokenize
	currentPos int    // points to the current position in the input
	readPos    int    // current read position after the current char
	char       rune   // the current char (byte as the binary representation of )
}

/*
* Init a Lexer
 */
func InitLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.ReadChar() // READ FIRST CHAR 
	return l
}


/*
* LEXER METHODS 
*/
func (l *Lexer) NextToken() token.Token {
	// var tokens []token.Token
	var test token.Token

	switch l.char{
		case '=':
			test = token.NewToken(token.ASSIGN,l.char);
		case '+':
			test = token.NewToken(token.PLUS, l.char)
		case ')':
			test = token.NewToken(token.RP,l.char)
		case '(':
			test = token.NewToken(token.LP,l.char)
		case '{':
			test = token.NewToken(token.LCB, l.char)
		case '}':
			test = token.NewToken(token.RCB, l.char)
		case ',':
			test = token.NewToken(token.CAMMA, l.char)
		case ';':
			test = token.NewToken(token.SCOLON, l.char)
		case 0:
			// program ends here 
			test = token.NewToken(token.FILEENDED, 0)
		default:
			
	}

	fmt.Print(test)

	l.ReadChar() // move to next char
	return test
}

// HELPER FUNCTIONS
/*
* This function give us the next character
*/
func (l *Lexer) ReadChar(){
	
	if l.readPos >= utf8.RuneCount([]byte(l.input)){ // the number of runes in the string
		l.char = 0 // SET THE CURRENT CHAR TO NUL CHARACTER (TO INDICATE THE TERMINATION OF THE STRING)
	}else{
		r, size :=utf8.DecodeRuneInString(l.input[l.readPos:])
		l.char = r
		l.currentPos = l.readPos;
		l.readPos += size
	} 
}



/*
* 	this functions reads the identifiers and keywords
*   starting with a letter and stops when it finds a non letter character
*/
func (l *Lexer) ReadIdentifier() string{
	currPosition := l.currentPos;
	
	//identifiers can't start with numbers 
	if currPosition ==0 && utils.IsDigit(l.char){
		return "" // this identifier is invalid 
	}

	for utils.IsDigit(l.char) || utils.IsLetter(l.char)  {
		l.ReadChar()
	}
	return l.input[currPosition:l.readPos]
}


