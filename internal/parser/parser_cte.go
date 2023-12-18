package parser

// this ctes will be used to handle operator precedence
const (
	_ int = iota
	LOWEST
	EQUALS          //==
	LESS_OR_GREATER // > <
	SUM             // +
	PRODUCT         // *
	PREFIX          // -a or !a
	CALL            // hello(a)
)
