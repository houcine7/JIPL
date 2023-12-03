package lexer

import (
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
	l.readChar() // READ FIRST CHAR 
	return l
}


/*
* LEXER METHODS 
*/
func (l *Lexer) NextToken() token.Token {
	// var tokens []token.Token
	var test token.Token
	l.ignoreWhiteSpace()

	switch l.char{
		case '=':
			if l.peek() =='='{
				prev :=l.char
				l.readChar()
				test = token.NewToken(token.EQUAL, string(prev)+ string(l.char));
			} else {
				test = token.NewToken(token.ASSIGN,string(l.char));
			}
		case '+':
			test = token.NewToken(token.PLUS, string(l.char))
		case '-':
			test = token.NewToken(token.MINUS,string(l.char))
		case '/':
			test = token.NewToken(token.SLASH,string(l.char))
		case '*':
			test = token.NewToken(token.STAR,string(l.char))
		case '!':
			if l.peek() =='=' {
				prev := l.char
				l.readChar()
				test = token.NewToken(token.NOT_EQUAL,string(prev)+ string(l.char))
			}else {
				test = token.NewToken(token.EX_MARK,string(l.char))
			}
		case '<':
			if l.peek() =='='{
				prev :=l.char
				l.readChar()
				test = token.NewToken(token.LT_OR_EQ, string(prev) + string(l.char))
			}else{
				test = token.NewToken(token.LT,string(l.char))
			}
		case '>':
			if l.peek() == '='{
				prev :=l.char
				l.readChar()
				test = token.NewToken(token.GT_OR_EQ,string(prev) + string(l.char))
			}else{
				test = token.NewToken(token.GT,string(l.char))
			}
		case ')':
			test = token.NewToken(token.RP,string(l.char))
		case '(':
			test = token.NewToken(token.LP,string(l.char))
		case '{':
			test = token.NewToken(token.LCB, string(l.char))
		case '}':
			test = token.NewToken(token.RCB, string(l.char))
		case ',':
			test = token.NewToken(token.COMMA, string(l.char))
		case ';':
			test = token.NewToken(token.S_COLON, string(l.char))
		case 0:
			// program ends here 
			test = token.NewToken(token.FILE_ENDED, string(rune(0)))
		default:
			if utils.IsLetter(l.char)  {
				ident := l.ReadIdentifier()
				test = token.NewToken(token.GetIdentifierTokenType(ident),ident)
				return test
			}else if utils.IsDigit(l.char){
				num := l.ReadNumber()
				test = token.NewToken(token.INT,num)
				return test // this prevents calling read char which is already done with the method ReadNumber()
			}else {
				test = token.NewToken(token.ILLEGAL,string(l.char))
			}
	}

	l.readChar() // move to next char
	return test
}

// HELPER FUNCTIONS
/*
* This function give us the next character
*/
func (l *Lexer) readChar(){
	
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
* This function peeks the next character 
* used in case of tokens that are compose of more than 2 tokens ( like "==" and "<=" ">=" and "!=")
*/

func (l *Lexer) peek() rune {
	// if we still in the input length range
	if l.readPos >= utf8.RuneCount([]byte(l.input)) {
		return 0;
	}
	return rune(l.input[l.readPos]);
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
		l.readChar()
	}
	return l.input[currPosition:l.currentPos]
}

/*
* this function reads the numbers 
* starting with a digit and stops when it reaches a non digit value
*/
func (l *Lexer) ReadNumber() string {
	currentPos := l.currentPos
	for utils.IsDigit(l.char) {
		l.readChar()
	}
	return l.input[currentPos:l.currentPos]
}


/*
*  function to skip white space and break lines
*/
func (l *Lexer) ignoreWhiteSpace(){
	for l.char ==' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}