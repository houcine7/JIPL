package lexer

import (
	"unicode/utf8"

	"github.com/houcine7/JIPL/internal/token"
	"github.com/houcine7/JIPL/pkg/utils"
)

type Lexer struct {
	input      string // the string to tokenize
	currentPos int    // points to the current position in the input
	readPos    int    // current read position after the current char
	char       rune   // the current char (byte as the binary representation of )
}

func InitLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // READ FIRST CHAR
	return l
}

/*
* LEXER METHODS
 */
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.ignoreWhiteSpace()

	switch l.char {
	case '=':
		if l.peek() == '=' {
			prev := l.char
			l.readChar()
			tok = token.CreateToken(token.EQUAL, string(prev)+string(l.char))
		} else {
			tok = token.CreateToken(token.ASSIGN, string(l.char))
		}
	case '&':
		if l.peek() == '&' {
			prev := l.char
			l.readChar()
			tok = token.CreateToken(token.AND, string(prev)+string(l.char))
		} else {
			tok = token.CreateToken(token.ILLEGAL, string(l.char))
		}

	case '|':
		if l.peek() == '|' {
			prev := l.char
			l.readChar()
			tok = token.CreateToken(token.OR, string(prev)+string(l.char))
		} else {
			tok = token.CreateToken(token.ILLEGAL, string(l.char))
		}
	case '+':
		if l.peek() == '+' {
			prev := l.char
			l.readChar()
			tok = token.CreateToken(token.INCREMENT, string(prev)+string(l.char))
		} else {
			tok = token.CreateToken(token.PLUS, string(l.char))
		}
	case '-':
		if l.peek() == '-' {
			prev := l.char
			l.readChar()
			tok = token.CreateToken(token.DECREMENT, string(l.char)+string(prev))
		} else {
			tok = token.CreateToken(token.MINUS, string(l.char))
		}
	case '/':
		tok = token.CreateToken(token.SLASH, string(l.char))
	case '%':
		tok = token.CreateToken(token.MODULO, string(l.char))
	case '*':
		tok = token.CreateToken(token.STAR, string(l.char))
	case '!':
		if l.peek() == '=' {
			prev := l.char
			l.readChar()
			tok = token.CreateToken(token.NOT_EQUAL, string(prev)+string(l.char))
		} else {
			tok = token.CreateToken(token.EX_MARK, string(l.char))
		}
	case '<':
		if l.peek() == '=' {
			prev := l.char
			l.readChar()
			tok = token.CreateToken(token.LT_OR_EQ, string(prev)+string(l.char))
		} else {
			tok = token.CreateToken(token.LT, string(l.char))
		}
	case '>':
		if l.peek() == '=' {
			prev := l.char
			l.readChar()
			tok = token.CreateToken(token.GT_OR_EQ, string(prev)+string(l.char))
		} else {
			tok = token.CreateToken(token.GT, string(l.char))
		}
	case ')':
		tok = token.CreateToken(token.RP, string(l.char))
	case '(':
		tok = token.CreateToken(token.LP, string(l.char))
	case '{':
		tok = token.CreateToken(token.LCB, string(l.char))
	case '}':
		tok = token.CreateToken(token.RCB, string(l.char))
	case ',':
		tok = token.CreateToken(token.COMMA, string(l.char))
	case ';':
		tok = token.CreateToken(token.S_COLON, string(l.char))
	case '"':
		tok = token.CreateToken(token.STRING, l.ReadString())
	case 0:
		// program ends here
		tok = token.CreateToken(token.FILE_ENDED, string(rune(0)))
	default:
		if utils.IsLetter(l.char) {
			ident := l.ReadIdentifier()
			tok = token.CreateToken(token.GetIdentifierTokenType(ident), ident)
			return tok
		} else if utils.IsDigit(l.char) {
			num := l.ReadNumber()
			tok = token.CreateToken(token.INT, num)
			return tok // this prevents calling read char which is already done with the method ReadNumber()
		} else {
			tok = token.CreateToken(token.ILLEGAL, string(l.char))
		}
	}

	l.readChar()
	return tok
}

// HELPER FUNCTIONS
/*
* This function give us the next character
 */
func (l *Lexer) readChar() {

	if l.readPos >= utf8.RuneCount([]byte(l.input)) {
		l.char = 0 // SET THE CURRENT CHAR TO NUL CHARACTER (TO INDICATE THE TERMINATION OF THE STRING)
	} else {
		r, size := utf8.DecodeRuneInString(l.input[l.readPos:])
		l.char = r
		l.currentPos = l.readPos
		l.readPos += size
	}
}

// read string literals
func (l *Lexer) ReadString() string {
	currPosition := l.currentPos + 1
	for {
		l.readChar()
		if l.char == '"' || l.char == 0 {
			break
		}
	}
	return l.input[currPosition:l.currentPos]
}

/*
* This function peeks the next character
 */

func (l *Lexer) peek() rune {
	// if we still in the input length range
	if l.readPos >= utf8.RuneCount([]byte(l.input)) {
		return 0
	}
	return rune(l.input[l.readPos])
}

/*
* 	this functions reads the identifiers and keywords
 */
func (l *Lexer) ReadIdentifier() string {
	currPosition := l.currentPos

	//identifiers can't start with numbers
	if currPosition == 0 && utils.IsDigit(l.char) {
		return ""
	}

	for utils.IsDigit(l.char) || utils.IsLetter(l.char) {
		l.readChar()
	}
	return l.input[currPosition:l.currentPos]
}

/*
* this function reads the numbers
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
func (l *Lexer) ignoreWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}
