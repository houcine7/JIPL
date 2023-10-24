package lexer

import (
	"github.com/houcine7/JIPL/token"
)

type Lexer struct {
	input      string // the string to tokenize
	currentPos int    // points to the current position in the input
	readPos    int    // current read position after the current char
	char       byte   // the current char (byte as the binary representation of )
}

/*
* Init a Lexer
 */

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l Lexer) NextToken() token.Token {
	return token.Token{Type: 1, Value: "="}
}
