package token

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

/*
* Create a new token helper function
 */
func CreateToken(tType TokenType, value string) Token {
	return Token{Type: tType, Value: value}
}

/*
*  The keywords map of our tokenizer
 */
var keywords = map[string]TokenType{
	"break":       BREAK,
	"continue":    CONTINUE,
	"return":      RETURN,
	"true":        TRUE,
	"false":       FALSE,
	"for":         FOR,
	"function":    FUNCTION,
	"def":         DEF,
	"if":          IF,
	"else":        ELSE,
	"class":       CLASS,
	"constructor": CONSTRUCTOR,
}

func GetIdentifierTokenType(identifier string) TokenType {
	if tokenType, ok := keywords[identifier]; ok {
		return tokenType
	}
	return IDENTIFIER
}
