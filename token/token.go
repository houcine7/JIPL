package token

/*
* We are going to use int as the token type to provide more performance
* Using string's as token type would be less performant
* And also the int allows Us to define many values
* String's will provide less debugging headache '(that's what's fun to more challenging)
 */
type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

/*
* Create a new token helper function
*/

func NewToken (tType TokenType, value string) Token{
	return Token{Type: tType , Value: value};
}

/*
*  The keywords map of our tokenizer
*  Contains keywords of the tokenizer and their corresponding literals 
*/
var keywords = map[string]TokenType{
	"function": FUNCTION,
	"def": DEF,
	"if": IF,
	"else": ELSE,
	"break": BREAK,
	"continue": CONTINUE,
	"return": RETURN,
}

func GetIdentifierTokenType(identifier string) TokenType{
	if tokenType, ok := keywords[identifier]; ok{
		return tokenType
	}
	return IDENTIFIER
}