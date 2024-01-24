package token

const (
	// SPECIALS
	ILLEGAL    = -1 // unknown token
	FILE_ENDED = 0  //file ended

	// references
	IDENTIFIER = 1

	// literals [2,20]
	INT    = 2 // int values
	STRING = 3 // string values

	//OPERATORS  values [40,80]
	ASSIGN    = 40 // =
	PLUS      = 41 // +
	MINUS     = 42 // -
	STAR      = 43 // *
	SLASH     = 44 // /
	EX_MARK   = 45 // !
	EQUAL     = 46 // ==
	NOT_EQUAL = 47 // !=
	INCREMENT = 50 // ++
	DECREMENT = 51 // --
	MODULO    = 56 // %
	AND       = 57 // &&
	OR        = 58 // ||

	/*Comparators operators*/
	LT       = 52 // <
	GT       = 53 // >
	LT_OR_EQ = 54 // <=
	GT_OR_EQ = 55 // >=

	//DELIMITERS [20,39]
	COMMA   = 20 // ,
	S_COLON = 21 // ;

	LP = 22 // (
	RP = 23 // )

	LCB = 24 // {
	RCB = 25 // }
	RB  = 26 // [
	LB  = 27 // ]

	// KEYWORDS [100,110]
	FUNCTION = 100 // function keyword
	DEF      = 101 // an identifier definition
	IF       = 102 // if token
	ELSE     = 103 // else token
	RETURN   = 104 // return statement toke
	BREAK    = 105 // break token
	CONTINUE = 106 // continue token
	TRUE     = 107
	FALSE    = 108
	FOR      = 109
)
