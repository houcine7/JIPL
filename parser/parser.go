package parser

import (
	"fmt"

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
	errors []string
}

func InitParser(l *lexer.Lexer) *Parser{
	p := &Parser{
		lexer: l,
		errors: []string{},
	}

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
* function to return the encountered errors 
* while parsing
*/

func (p *Parser) Errors() []string{
	return p.errors
}

/*
 This function is to parse a given program
*/
func (p *Parser) Parse() *ast.Program{
	program := &ast.Program{}
	program.Statements = []ast.Statement{};

	for !p.currentTokenEquals(token.FILE_ENDED) {
		stm := p.parseStmt()
		//fmt.Println(stm.TokenLiteral())

		if stm !=nil {
			program.Statements =append(program.Statements, stm )
		}
		// Advance with token
		p.NextToken()
	}

	return program
}

/*
* a parser function to parse statements 
*  and returns the parsed statement 
*/

func (p *Parser) parseStmt() ast.Statement{
	switch p.currToken.Type {
	case token.DEF:
		return p.parseDefStmt()
	case token.RETURN:
		return p.parserReturnStmt()
	default:
		return nil
	}
}

/*
* function used to parse def statement
*/
func (p *Parser) parseDefStmt() *ast.DefStatement {
	stm := &ast.DefStatement{Token: p.currToken}
	
	// syntax error's
	if !p.expectedNextToken(token.NewToken(token.IDENTIFIER,"IDENT")){
		return nil
	}
	stm.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Value,
	}
	if !p.expectedNextToken(token.NewToken(token.ASSIGN,"=")){
		return nil
	}

	for !p.currentTokenEquals(token.S_COLON){
		p.NextToken();
	}
	
	return stm
}

func (p *Parser) parserReturnStmt() *ast.ReturnStatement {
	stm := &ast.ReturnStatement{Token: p.currToken}

	p.NextToken()

	//TODO:
	for !p.currentTokenEquals(token.S_COLON) {
		p.NextToken()
	}

	return stm

}


func (p *Parser) currentTokenEquals(t token.TokenType) bool{
	return p.currToken.Type == t;
}

func (p *Parser) peekTokenEquals(t token.TokenType) bool {
	return p.peekedToken.Type == t
}

/*
* function checks if the given token is the next token
* returns true and advances the tokens pointers of the parser
* if not returns false
*/
func (p *Parser) expectedNextToken(t token.Token) bool{
	if p.peekTokenEquals(t.Type){
		p.NextToken()
		return true
	}
	// peek errors
	p.peekedError(t)
	return false
}

/*
* function to peek Errors in next token
* takes the peeked wrong token and adds error
* message to the errors array of the parser
*/
func (p *Parser) peekedError(encounteredToken token.Token) {
	
	errorMessage := fmt.Sprintf("wrong next token type expected token is %s instead got %s",
	encounteredToken.Value, p.peekedToken.Value)

	// append message to the errors array
	p.errors =append(p.errors, errorMessage)
}