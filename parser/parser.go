package parser

import (
	"fmt"
	"strconv"

	ast "github.com/houcine7/JIPL/AST"
	"github.com/houcine7/JIPL/lexer"
	"github.com/houcine7/JIPL/token"
)

/*
* The parser struct defining the parser type
* which is used to parse the tokens
 */
type Parser struct {
	lexer  *lexer.Lexer // points to the lexer
	errors []string     // parser errors

	currToken   token.Token // the current token in examination
	peekedToken token.Token // the next token after the current one

	prefixParseFuncs map[token.TokenType]prefixParse
	infixParseFuncs  map[token.TokenType]infixParse
}

/*
Types of expression parsing functions
*/
type (
	prefixParse func() ast.Expression
	// takes param as the left operand of the infix operator
	infixParse func(ast.Expression) ast.Expression
)

func InitParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	p.NextToken() // to peek the first token
	p.NextToken() // first token in the currentToken

	// init ParseFunctions maps
	p.prefixParseFuncs = make(map[token.TokenType]prefixParse)
	// registering prefix parsing functions
	p.addPrefixFn(token.IDENTIFIER, p.parseIdentifier)
	p.addPrefixFn(token.INT, p.parseInt)

	// prefix expression parser
	prefixParseTokens := []token.TokenType{
		token.EX_MARK,
		token.MINUS,
	}
	p.addAllPrefixFn(prefixParseTokens, p.parsePrefixExpression)

	// infix expresion parseres
	infixParseTokens := []token.TokenType{
		tokne.PLUS,
		token.MINUS,
		token.SLASH,
		token.STAR,

		token.EQUAL,
		token.NOT_EQUAL,

		token.LT,
		token.GT,
		token.LT_OR_EQ,
		token.GT_OR_EQ,
	}

	p.addALlInfixFn(infixParseTokens, p.parseInfixExpression)

	return p
}

/*
Helper function to move the pointer of the token in the lexer
reads the next token stores it on the peek but before store the previous
peekedToken on currentToken
*/
func (p *Parser) NextToken() {
	p.currToken = p.peekedToken
	p.peekedToken = p.lexer.NextToken()
}

/*
This function is to parse a given program
*/
func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenEquals(token.FILE_ENDED) {
		stm := p.parseStmt()
		//fmt.Println(stm.TokenLiteral())
		// if stm !=nil {
		program.Statements = append(program.Statements, stm)
		// }
		// Advance with token
		p.NextToken()
	}

	return program
}

/*
* a parser function to parse statements
*  and returns the parsed statement
 */
func (p *Parser) parseStmt() ast.Statement {
	switch p.currToken.Type {
	case token.DEF:
		return p.parseDefStmt()
	case token.RETURN:
		return p.parseReturnStmt()
	// left are expression statement
	default:
		return p.parseExpressionStatement()
	}
}

/*
* function used to parse def statement
 */
func (p *Parser) parseDefStmt() *ast.DefStatement {
	stm := &ast.DefStatement{Token: p.currToken}

	// syntax error's
	if !p.expectedNextToken(token.NewToken(token.IDENTIFIER, "IDENT")) {
		return nil
	}
	stm.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Value,
	}
	if !p.expectedNextToken(token.NewToken(token.ASSIGN, "=")) {
		return nil
	}

	for !p.currentTokenEquals(token.S_COLON) {
		p.NextToken()
	}

	return stm
}

// to parse Return statement
func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
	stm := &ast.ReturnStatement{Token: p.currToken}

	p.NextToken()

	//TODO:
	for !p.currentTokenEquals(token.S_COLON) {
		p.NextToken()
	}

	return stm
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {

	fmt.Println("------->", p.currToken)
	stm := &ast.ExpressionStatement{Token: p.currToken}

	// fmt.Println(p.parseExpression(token.IDENTIFIER))
	stm.Expression = p.parseExpression(LOWEST)

	// fmt.Println(stm)
	if p.peekTokenEquals(token.S_COLON) {
		p.NextToken()
	}

	return stm
}

func (p *Parser) parseExpression(weight int) ast.Expression {
	prefix := p.prefixParseFuncs[p.currToken.Type]

	if prefix == nil {
		// add error msg
		p.notFoundPrefixFunctionError(p.currToken)
		return nil
	}
	leftExpression := prefix()

	return leftExpression
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Value}
}

func (p *Parser) parseInt() ast.Expression {
	exp := &ast.IntegerLiteral{Token: p.currToken}
	val, err := strconv.ParseInt(p.currToken.Value, 0, 0)

	if err != nil {
		fmt.Println(p.currToken)
		errMsg := fmt.Sprintf("Parsing error, couldn't parse string %s to Integer value",
			p.currToken.Value)
		p.errors = append(p.errors, errMsg)
		return nil
	}
	exp.Value = int(val)
	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Value,
	}

	p.NextToken()
	exp.Right = p.parseExpression(PREFIX)
	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Value,
		Left:     left,
	}

	precedence := p.currentPrecedence()
	p.NextToken()

	exp.Right = p.parseExpression(precedence)

	return exp

}

// ERRORS
func (p *Parser) notFoundPrefixFunctionError(t token.Token) {
	msg := fmt.Sprintf("no prefix function for the given tokenType={%d,%s} found", t.Type, t.Value)
	p.errors = append(p.errors, msg)
}

/*
* function to return the encountered errors
* while parsing
 */
func (p *Parser) Errors() []string {
	return p.errors
}

// Helper functions
func (p *Parser) currentTokenEquals(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) peekTokenEquals(t token.TokenType) bool {
	return p.peekedToken.Type == t
}

/*
* function checks if the given token is the next token
* returns true and advances the tokens pointers of the parser
* if not returns false
 */
func (p *Parser) expectedNextToken(t token.Token) bool {
	if p.peekTokenEquals(t.Type) {
		p.NextToken()
		return true
	}
	// peek errors
	p.peekedError(t)
	return false
}

// precedence map to map tokens with thier precedence
var precedences = map[token.TokenType]int{
	token.EQUAL:     EQUALS,
	token.NOT_EQUAL: EQUALS,

	token.LT:       LESS_OR_GREATER,
	token.GT:       LESS_OR_GREATER,
	token.LT_OR_EQ: LESS_OR_GREATER,
	token.GT_OR_EQ: LESS_OR_GREATER,

	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.SLASH: PRODUCT,
	token.STAR:  PRODUCT,
}

// function to peekPrecedence
/* Function to peek precedence of the next token*/
func (p *Parser) peekPrecedence() int {
	if prec, ok := precedences[p.peekedToken.Type]; ok {
		return prec
	}
	// if not present retrun the lowset(default precedence)
	return LOWEST
}

// to return the current precedence
func (p *Parser) currentPrecedence() int {
	if prec, ok := precedences[p.currToken.Type]; ok {
		return prec
	}
	return LOWEST
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
	p.errors = append(p.errors, errorMessage)
}

// functions to add entries to the prefixParseFun and  infixParseFun
func (p *Parser) addPrefixFn(tokenType token.TokenType, prefixF prefixParse) {
	p.prefixParseFuncs[tokenType] = prefixF
}

func (p *Parser) addAllPrefixFn(tokenTypesSlice []token.TokenType, prefixFn prefixParse) {
	for _, tokenType := range tokenTypesSlice {
		p.addPrefixFn(tokenType, prefixFn)
	}
}

func (p *Parser) addInfixFn(tokenType token.TokenType, infixF infixParse) {
	p.infixParseFuncs[tokenType] = infixF
}

func (p *Parser) addALlInfixFn(tokenTypesSlice []token.TokenType, infixFn infixParse) {
	for _, tokenType := range tokenTypesSlice {
		p.addInfixFn(tokenType, infixFn)
	}
}
