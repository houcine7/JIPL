package token

/*
*As the first iteration OUR JIPL lang
* will define limited tokens so we declare em
* as constants
 */
 const (
	// SPECIALS
	ILLEGAL   = -1 // unknown token
	FILE_ENDED = 0  //file ended
	
	
	// references  
	IDENTIFIER = 1
	
	// literals [2,20]
	INT        = 2 // int values
	

	//OPERATORS  values [40,80]
	ASSIGN = 40 // =
	PLUS   = 41 // +
	MINUS  = 42 // -
	STAR   = 43 // *
	SLASH  = 44 // /
	EX_MARK = 45 // !
	EQUAL   = 46 // ==
	NOT_EQUAL = 47 // !=
    INCREMENT = 50 // ++
 	DECREMENT = 51 // --

	/*Comparators operators*/
	LT = 46 // <
	GT = 47 // >
	LT_OR_EQ = 48 // <=
	GT_OR_EQ = 49 // >=


	//DELIMITERS [20,39]
	COMMA  = 20 // ,
	S_COLON = 21 // ;

	LP = 22 // (
	RP = 23 // )

	LCB = 24 // {
	RCB = 25 // }

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
)
