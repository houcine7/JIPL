package parser

import (
	ast "github.com/houcine7/JIPL/AST"
	"github.com/houcine7/JIPL/lexer"
	"github.com/houcine7/JIPL/token"
)

/*
* The parser struct defining the parser type
* which is used to parse the tokens
 */
type Parser struct {
	lexer *lexer.Lexer // points to the lexer 
	currToken token.Token // the current token in examination
	peekedToken token.Token // the next token after the current one
}

func InitParser(l *lexer.Lexer) *Parser{
	p := &Parser{lexer: l}

	p.NextToken() // to peek the first token
	p.NextToken() // first token in the currentToken

	return p;
}


/*
 Helper function to move the pointer of the token in the lexer
 reads the next token stores it on the peek but before store the previous 
 peekedToken on currentToken
*/
func (p *Parser) NextToken(){
	p.currToken = p.peekedToken
	p.peekedToken = p.lexer.NextToken()
}
/*
 This function is to parse a given program
*/

func (p *Parser) Parse() *ast.Program{
	return nil
}
