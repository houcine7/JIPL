package token

const (
	// SPECIALS
	_          TokenType = iota
	ILLEGAL              // unknown token
	FILE_ENDED           //file ended

	// references
	IDENTIFIER

	// literals [2,]
	INT    // int values
	STRING // string values

	//OPERATORS  values [40,80]
	ASSIGN    // =
	PLUS      // +
	MINUS     // -
	STAR      // *
	SLASH     // /
	EX_MARK   // !
	EQUAL     // ==
	NOT_EQUAL // !=
	INCREMENT // ++
	DECREMENT // --
	MODULO    // %
	AND       // &&
	OR        // ||

	/*Comparators operators*/
	LT       // <
	GT       // >
	LT_OR_EQ // <=
	GT_OR_EQ // >=

	//DELIMITERS [20,39]
	COMMA   // ,
	S_COLON // ;

	LP // (
	RP // )

	LCB // {
	RCB // }
	RB  // ]
	LB  // [

	// KEYWORDS
	FUNCTION // function keyword
	DEF      // an identifier definition
	IF       // if token
	ELSE     // else token
	RETURN   // return statement toke
	BREAK    // break token
	CONTINUE // continue token
	TRUE
	FALSE
	FOR

	CLASS       // the class key word to create a class
	CONSTRUCTOR // constructor keyword
)
