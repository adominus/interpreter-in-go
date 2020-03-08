package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string {},
	}

	// Populate cur and peek Token
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.isCurrentToken(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	}
	return nil
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.advanceIfPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.advanceIfPeek(token.ASSIGN) {
		return nil
	}

	// TODO: Will fill in later
	// For now skip until end of statement
	for !p.isCurrentToken(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: Will fill in later
	// For now skip until end of statement
	for !p.isCurrentToken(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) peekError( t token.TokenType) {
	message := fmt.Sprintf("Expected next token to be of type %s. Instead received %s",
		t, p.peekToken)

	p.errors = append(p.errors, message)
}

func (p *Parser) isCurrentToken(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) isPeekToken(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) advanceIfPeek(expectedToken token.TokenType) bool {
	if p.isPeekToken(expectedToken) {
		p.nextToken()
		return true
	} else {
		p.peekError(expectedToken)
		return false
	}
}
