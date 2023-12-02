package token

/*
*As the first iteration OUR JIPL lang
* will define limited tokens so we declare em
* as constants
 */
 const (
	// SPECIALS
	ILLEGAL   = -1 // unknown token
	FILEENDED = 0  //file ended

	//Identifiers and literals [1,20]
	IDENTIFIER = 1 // variables
	INT        = 2 // int values

	//OPERATORS  values [40,80]
	ASSIGN = 40 // =
	PLUS   = 41 // +
	MINUS  = 42 // -
	STAR   = 43 // *
	SLASH  = 44 // /
	EXMARK = 45 // !

	LT = 46 // <
	GT = 47 // >


	//DELIMITERS [20,39]
	COMMA  = 20 // ,
	SCOLON = 21 // ;

	LP = 22 // (
	RP = 23 // )

	LCB = 24 // {
	RCB = 25 // }

	// KEYWORDS [100,110]
	FUNCTION = 100
	DEF      = 101 // a variable definition
)
