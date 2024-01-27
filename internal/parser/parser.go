package parser

import (
	"fmt"
	"strconv"

	ast "github.com/houcine7/JIPL/internal/AST"
	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/token"
)

type Parser struct {
	lexer  *lexer.Lexer // points to the lexer
	errors []*Error     // parser errors

	currToken   token.Token // the current token in examination
	peekedToken token.Token // the next token to parse

	prefixParseFuncs map[token.TokenType]prefixParse // function used for prefix parsing
	infixParseFuncs  map[token.TokenType]infixParse  // function used for infix parsing
}

type Error struct {
	Message string
	Token   token.Token
}

/*
types of expression parsing functions
*/
type (
	prefixParse func() ast.Expression
	// takes param as the left operand of the infix operation
	infixParse func(ast.Expression) ast.Expression
)

func InitParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: make([]*Error, 0),
	}

	p.Next() // to peek the first token
	p.Next() // first token in the currentToken

	// init ParseFunctions maps
	p.prefixParseFuncs = make(map[token.TokenType]prefixParse)
	// registering prefix parsing functions
	p.addPrefixFn(token.IDENTIFIER, p.parseIdentifier)
	p.addPrefixFn(token.INT, p.parseInt)
	p.addAllPrefixFn([]token.TokenType{
		token.TRUE,
		token.FALSE,
	}, p.parseBoolean)
	p.addPrefixFn(token.LP, p.parseGroupExpression)
	p.addAllPrefixFn([]token.TokenType{
		token.EX_MARK,
		token.MINUS,
	}, p.parsePrefixExpression)
	p.addPrefixFn(token.IF, p.parseIfExpression)
	p.addPrefixFn(token.FUNCTION, p.parseFunctionExpression)
	p.addPrefixFn(token.FOR, p.parseForLoopExpression)
	p.addPrefixFn(token.STRING, p.parseStringLit)
	p.addPrefixFn(token.LB, p.parseArrayLit)

	// infix expression parser functions
	p.infixParseFuncs = make(map[token.TokenType]infixParse)
	infixParseTokens := []token.TokenType{
		token.PLUS,
		token.MINUS,
		token.SLASH,
		token.STAR,
		token.MODULO,

		token.EQUAL,
		token.NOT_EQUAL,

		token.LT,
		token.GT,
		token.LT_OR_EQ,
		token.GT_OR_EQ,
		token.AND,
		token.OR,
	}
	p.addALlInfixFn(infixParseTokens, p.parseInfixExpression)
	p.addInfixFn(token.LP, p.parseFunctionCallExp)
	p.addALlInfixFn([]token.TokenType{
		token.DECREMENT,
		token.INCREMENT,
	}, p.parsePostFixExpression)
	p.addInfixFn(token.LB, p.parseIndexExp)

	return p
}
func (p *Parser) printTrace(a ...any) {
	const dots = ". . . . . . . . . . . . . . . . . . . . . . "
	const lnDots = len(dots)

	fmt.Print(dots[0:4])
	fmt.Println(a...)
}

func trace(msg string, p *Parser) *Parser {
	p.printTrace(msg, "(")
	return p
}

/*
Helper function to move the token pointers of the parser
*/
func (p *Parser) Next() {
	p.currToken = p.peekedToken
	p.peekedToken = p.lexer.NextToken()
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.currentTokenEquals(token.FILE_ENDED) {
		stm := p.parseStmt()
		program.Statements = append(program.Statements, stm)
		// Advance with token
		p.Next()
	}
	return program
}

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
	if !p.expectedNextToken(token.CreateToken(token.IDENTIFIER, "IDENT")) {
		return nil
	}
	stm.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Value,
	}
	if !p.expectedNextToken(token.CreateToken(token.ASSIGN, "=")) {
		return nil
	}

	p.Next() // move to the expression after =

	stm.Value = p.parseExpression(LOWEST)

	if p.peekTokenEquals(token.S_COLON) {
		p.Next()
	}

	return stm
}

func (p *Parser) parseDefStmtInForLoop() *ast.DefStatement {
	stm := &ast.DefStatement{Token: p.currToken}

	// syntax error's
	if !p.expectedNextToken(token.CreateToken(token.IDENTIFIER, "IDENT")) {
		return nil
	}
	stm.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Value,
	}
	if !p.expectedNextToken(token.CreateToken(token.ASSIGN, "=")) {
		return nil
	}

	p.Next()

	stm.Value = p.parseExpression(LOWEST)

	return stm
}

// to parse for loop expressions
func (p *Parser) parseForLoopExpression() ast.Expression {

	exp := &ast.ForLoopExpression{Token: p.currToken}

	if !p.expectedNextToken(token.CreateToken(token.LP, "(")) {
		return nil
	}
	p.Next() // advance to init parseStmt
	exp.InitStm = p.parseDefStmtInForLoop()

	if !p.expectedNextToken(token.CreateToken(token.S_COLON, ";")) {
		return nil
	}
	p.Next() // advance to the condition expression
	exp.Condition = p.parseExpression(LOWEST)

	if !p.expectedNextToken(token.CreateToken(token.S_COLON, ";")) {
		return nil
	}
	p.Next() // advance to post iteration
	exp.PostIteration = p.parseExpression(LOWEST)

	if !p.expectedNextToken(token.CreateToken(token.RP, ")")) {
		return nil
	}

	if !p.expectedNextToken(token.CreateToken(token.LCB, "{")) {
		return nil
	}
	exp.Body = p.parseBlocStatements()

	return exp

}

// to parse functionExpression
func (p *Parser) parseFunctionExpression() ast.Expression {
	exp := &ast.FunctionExp{Token: p.currToken}

	if !p.expectedNextToken(token.CreateToken(token.IDENTIFIER, "")) {
		return nil
	}
	exp.Name = p.parseIdentifier().(*ast.Identifier)
	if !p.expectedNextToken(token.CreateToken(token.LP, "(")) {
		return nil
	}

	// fn params
	exp.Parameters = p.parsePramas()

	if !p.expectedNextToken(token.CreateToken(token.LCB, "{")) {
		return nil
	}
	// fn body should start with  {
	exp.FnBody = p.parseBlocStatements()

	return exp
}

func (p *Parser) parsePramas() []*ast.Identifier {
	var params = []*ast.Identifier{}

	if p.peekTokenEquals(token.RP) { // in case of no param
		p.Next()
		return params
	}

	p.Next() // advance to params p1,p2....
	params = append(params, p.parseIdentifier().(*ast.Identifier))

	for p.peekTokenEquals(token.COMMA) {
		p.Next()
		p.Next() //advance to next param
		params = append(params, p.parseIdentifier().(*ast.Identifier))
	}
	//end of params
	if !p.expectedNextToken(token.CreateToken(token.RP, ")")) {
		return nil
	}

	return params
}

func (p *Parser) parseFunctionCallExp(fn ast.Expression) ast.Expression {
	exp := &ast.FunctionCall{
		Token:    p.currToken,
		Function: fn,
	}
	exp.Arguments = p.parseExpressionList(token.CreateToken(token.RP, ")"))

	return exp
}

func (p *Parser) parseStringLit() ast.Expression {
	exp := &ast.StringLiteral{Token: p.currToken, Value: p.currToken.Value}
	return exp
}

func (p *Parser) parseReturnStmt() *ast.ReturnStatement {
	stm := &ast.ReturnStatement{Token: p.currToken}

	p.Next()
	stm.ReturnValue = p.parseExpression(LOWEST)
	//
	for p.peekTokenEquals(token.S_COLON) {
		p.Next()
	}

	return stm
}

// group expression
func (p *Parser) parseGroupExpression() ast.Expression {
	p.Next()
	grpExp := p.parseExpression(LOWEST)

	if !p.expectedNextToken(token.CreateToken(token.RP, ")")) {
		return nil
	}
	return grpExp
}

// parser if expressions
func (p *Parser) parseIfExpression() ast.Expression {

	exp := &ast.IfExpression{Token: p.currToken}

	// going to ( token
	if !p.expectedNextToken(token.CreateToken(token.LP, "(")) {
		return nil
	}

	//advances to token right after (
	p.Next()
	exp.Condition = p.parseExpression(LOWEST)

	if !p.expectedNextToken(token.CreateToken(token.RP, ")")) {
		return nil
	}
	if !p.expectedNextToken(token.CreateToken(token.LCB, "{")) {
		return nil
	}
	exp.Body = p.parseBlocStatements()

	if p.peekTokenEquals(token.ELSE) {
		p.Next()
		if !p.expectedNextToken(token.CreateToken(token.LCB, "{")) {
			return nil
		}
		exp.ElseBody = p.parseBlocStatements()
	}

	return exp
}

func (p *Parser) parseBlocStatements() *ast.BlockStm {
	blockStm := &ast.BlockStm{Token: p.currToken}
	stms := []ast.Statement{}
	p.Next()

	for !p.currentTokenEquals(token.RCB) &&
		!p.currentTokenEquals(token.FILE_ENDED) {
		stm := p.parseStmt()
		stms = append(stms, stm)
		p.Next()
	}
	blockStm.Statements = stms
	return blockStm

}

// expression statments parsing
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {

	stm := &ast.ExpressionStatement{Token: p.currToken}

	stm.Expression = p.parseExpression(LOWEST)

	if p.peekTokenEquals(token.S_COLON) {
		p.Next()
	}

	return stm
}

func (p *Parser) parseExpression(precedence int) ast.Expression {

	prefix := p.prefixParseFuncs[p.currToken.Type]

	if prefix == nil {
		// add error msg
		p.notFoundPrefixFunctionError(p.currToken)
		return nil
	}

	leftExpression := prefix()

	for !p.peekTokenEquals(token.S_COLON) && precedence < p.peekPrecedence() {

		infix := p.infixParseFuncs[p.peekedToken.Type]
		if infix == nil {
			return leftExpression
		}

		p.Next()
		//
		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (p *Parser) parseIdentifier() ast.Expression {
	stm := &ast.Identifier{Token: p.currToken,
		Value: p.currToken.Value}
	if p.peekTokenEquals(token.ASSIGN) {
		p.Next()
		exp := p.parseAssignmentExpr(stm)
		return exp
	}
	return stm
}

func (p *Parser) parseInt() ast.Expression {
	exp := &ast.IntegerLiteral{Token: p.currToken}
	val, err := strconv.ParseInt(p.currToken.Value, 0, 0)

	if err != nil {
		errMsg := fmt.Sprintf("Parsing error, couldn't parse string %s to Integer value",
			p.currToken.Value)
		p.errors = append(p.errors, &Error{Message: errMsg, Token: p.currToken})
		return nil
	}
	exp.Value = int(val)
	return exp
}

func (p *Parser) parseBoolean() ast.Expression {
	exp := &ast.BooleanExp{
		Token: p.currToken,
		Value: p.currentTokenEquals(token.TRUE),
	}

	return exp
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Value,
	}

	p.Next()
	exp.Right = p.parseExpression(PREFIX)
	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Value,
		Left:     left,
	}

	prevPrecedence := p.currentPrecedence()
	p.Next()
	exp.Right = p.parseExpression(prevPrecedence)

	return exp
}

func (p *Parser) parsePostFixExpression(left ast.Expression) ast.Expression {
	exp := &ast.PostfixExpression{
		Token:    p.currToken,
		Operator: p.currToken.Value,
		Left:     left,
	}
	return exp
}

func (p *Parser) parseAssignmentExpr(left *ast.Identifier) ast.Expression {
	exp := &ast.AssignmentExpression{
		Token: p.currToken,
		Left:  left,
	}

	p.Next()
	exp.AssignmentValue = p.parseExpression(LOWEST)

	return exp
}

func (p *Parser) parseArrayLit() ast.Expression {
	exp := &ast.ArrayLiteral{
		Token: p.currToken,
	}

	exp.Values = p.parseExpressionList(token.CreateToken(token.RB, "]"))

	return exp
}

func (p *Parser) parseExpressionList(t token.Token) []ast.Expression {
	var res = []ast.Expression{}

	if p.peekTokenEquals(t.Type) {
		p.Next()
		return res // an empty array
	}

	p.Next()
	res = append(res, p.parseExpression(LOWEST))

	for p.peekTokenEquals(token.COMMA) {
		p.Next()
		p.Next()

		res = append(res, p.parseExpression(LOWEST))
	}

	if !p.expectedNextToken(t) {
		return nil
	}

	return res

}

func (p *Parser) parseIndexExp(left ast.Expression) ast.Expression {

	exp := &ast.IndexExpression{
		Token: p.currToken,
		Left:  left,
	}

	p.Next()
	exp.Index = p.parseExpression(LOWEST)

	if !p.expectedNextToken(token.CreateToken(token.RB, "]")) {
		return nil
	}

	return exp

}

// ERRORS
func (p *Parser) notFoundPrefixFunctionError(t token.Token) {
	msg := fmt.Sprintf("no prefix function for the given tokenType={%d,%s} found", t.Type, t.Value)
	p.errors = append(p.errors, &Error{msg, t})
}

func (p *Parser) Errors() []*Error {
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
* function checks if the given token has the same type of the next token
*  advances the parser if true if not it adds parsing errors
 */
func (p *Parser) expectedNextToken(t token.Token) bool {
	if p.peekTokenEquals(t.Type) {
		p.Next()
		return true
	}
	// peek errors
	p.peekedError(t)
	return false
}

/* Function to peek precedence of the next token*/
func (p *Parser) peekPrecedence() int {
	if prec, ok := precedences[p.peekedToken.Type]; ok {
		return prec
	}
	return LOWEST
}

func (p *Parser) currentPrecedence() int {
	if prec, ok := precedences[p.currToken.Type]; ok {
		return prec
	}
	return LOWEST
}

/*
* function to add Error of wrong type of token
 */
func (p *Parser) peekedError(expectedToken token.Token) {

	errorMessage := fmt.Sprintf("wrong next token type expected token is %s instead got %s",
		expectedToken.Value, p.peekedToken.Value)

	// append message to the errors array
	p.errors = append(p.errors, &Error{errorMessage, p.peekedToken})
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
