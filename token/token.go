package token

/*
* We are going to use int as the token type to povide more performance
* Using string's as token type would be less performant
* And also the int allows Us to define many values
* String's will provide less debuging headache '(that's what's fun to more challanging)
*
 */

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

/*
*As the first iteration OUR JIPL lang
* will define limited tokens so we decalre em
* as constats
 */

const (
	// SPECIALS
	ILLEGAL   = -1 // unknown token
	FILEENDED = 0  //file ended

	//Identifiers and literals [1,20]
	IDENTIFIER = 1 // varaibles
	INT        = 2 // int values

	//OPERATORS  values [40,80]
	ASSIGN = 40
	PLUS   = 41

	//DELIMITERS [20,39]
	CAMMA  = 20 // ,
	SCOLON = 21 // ;

	LP = 22 // (
	RP = 23 // )

	LCB = 24 // {
	RCB = 25 // }

	// KEYWORDS [100,110]
	FUNCTION = 100
	DEF      = 101 // a variable definition
)


/*
* Create a new token helper function
*/

func NewToken (tType TokenType, value rune) Token{
	return Token{Type: tType , Value: string(value)};
}