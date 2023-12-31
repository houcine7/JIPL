package parser

import (
	"fmt"
	"strconv"

	ast "github.com/houcine7/JIPL/internal/AST"
	"github.com/houcine7/JIPL/internal/lexer"
	"github.com/houcine7/JIPL/internal/token"
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

	prefixParseFuncs map[token.TokenType]prefixParse // function used for prefix parsing
	infixParseFuncs  map[token.TokenType]infixParse // function used for infix parsing
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
	}, p.parseBoolen)
	p.addPrefixFn(token.LP, p.parseGroupExpression)
	// prefix expression parser
	p.addAllPrefixFn([]token.TokenType{
		token.EX_MARK,
		token.MINUS,
	}, p.parsePrefixExpression)
	p.addPrefixFn(token.IF, p.parseIfExpression)
	p.addPrefixFn(token.FUNCTION, p.parseFunctionExpression)
	p.addPrefixFn(token.FOR, p.parseForLoopExpression)
	p.addPrefixFn(token.STRING, p.parseStringLit)
	// infix expresion parseres
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
	/*
		fmt.Println("-------------------------")
		fmt.Println("Map prefix fns is:", p.prefixParseFuncs)
		fmt.Println("Map infix fns is", p.infixParseFuncs)
	*/
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
Helper function to move the pointer of the token in the lexer
reads the next token stores it on the peek but before store the previous
peekedToken on currentToken
*/
func (p *Parser) Next() {
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
		p.Next()
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

	p.Next() // move to the expression after =

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
	//fmt.Println("------ in Identifier ----")
	exp.Name = p.parseIdentifier().(*ast.Identifier)
	if !p.expectedNextToken(token.CreateToken(token.LP, "(")) {
		return nil
	}

	//fmt.Println("function.Name is ", exp.Name)
	//fmt.Println(p.currToken)

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

	exp.Arguments = p.parseFnArgs()

	return exp
}

func (p *Parser) parseFnArgs() []ast.Expression {

	var ans = []ast.Expression{}
	if p.peekTokenEquals(token.RP) {
		p.Next()
		return ans
	}

	p.Next()
	ans = append(ans, p.parseExpression(LOWEST))

	for p.peekTokenEquals(token.COMMA) {
		p.Next()
		p.Next()
		ans = append(ans, p.parseExpression(LOWEST))
	}

	if !p.expectedNextToken(token.CreateToken(token.RP, ")")) {
		return nil
	}

	return ans
}

// parse string lit
func (p *Parser) parseStringLit() ast.Expression {
	exp := &ast.StringLiteral{Token: p.currToken, Value: p.currToken.Value}
	return exp
}

// to parse Return statement
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
		//advance currToken
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
		if stm != nil {
			stms = append(stms, stm)
		}
		p.Next()
	}
	blockStm.Statements = stms
	return blockStm

}

// expression statments parsing
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {

	//debuggin puropos
	//defer trace("parse expression statements called..", p)
	//fmt.Println("------->", p.currToken)
	stm := &ast.ExpressionStatement{Token: p.currToken}

	// fmt.Println(p.parseExpression(token.IDENTIFIER))
	stm.Expression = p.parseExpression(LOWEST)

	// fmt.Println(stm)
	if p.peekTokenEquals(token.S_COLON) {
		p.Next()
	}

	return stm
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	
	// fmt.Println("parse exp", p.currToken)
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
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Value}
}

func (p *Parser) parseInt() ast.Expression {
	exp := &ast.IntegerLiteral{Token: p.currToken}
	val, err := strconv.ParseInt(p.currToken.Value, 0, 0)

	if err != nil {
		//fmt.Println(p.currToken)
		errMsg := fmt.Sprintf("Parsing error, couldn't parse string %s to Integer value",
			p.currToken.Value)
		p.errors = append(p.errors, errMsg)
		return nil
	}
	exp.Value = int(val)
	return exp
}

func (p *Parser) parseBoolen() ast.Expression {
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
* function checks if the given token has the same type of the next token
* it returns true and advances the tokens pointers of the parser
* if not returns false and adds a parser error
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
func (p *Parser) peekedError(expectedToken token.Token) {

	errorMessage := fmt.Sprintf("wrong next token type expected token is %s instead got %s",
		expectedToken.Value, p.peekedToken.Value)

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
