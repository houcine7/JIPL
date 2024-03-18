package parser

import (
	"github.com/houcine7/JIPL/internal/token"
)

// this ctes will be used to handle operator precedence
const (
	_ int = iota
	LOWEST
	EQUALS //==

	LESS_OR_GREATER // > <
	SUM             // +
	PRODUCT         // *
	PREFIX          // -a or !a
	CALL            // hello(a)

	INDEX // [ array indexing

	INCREMENT // -- ++
)

// precedence map to map tokens with their precedence
var precedences = map[token.TokenType]int{
	token.EQUAL:     EQUALS,
	token.NOT_EQUAL: EQUALS,

	token.LT:       LESS_OR_GREATER,
	token.GT:       LESS_OR_GREATER,
	token.LT_OR_EQ: LESS_OR_GREATER,
	token.GT_OR_EQ: LESS_OR_GREATER,

	token.PLUS:   SUM,
	token.MINUS:  SUM,
	token.SLASH:  PRODUCT,
	token.STAR:   PRODUCT,
	token.MODULO: PRODUCT,

	token.AND: EQUALS,
	token.OR:  EQUALS,

	token.LP: CALL,

	token.INCREMENT: INCREMENT,
	token.DECREMENT: INCREMENT,

	token.LB: INDEX,
}
