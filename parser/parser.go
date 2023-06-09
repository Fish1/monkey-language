package parser

import (
	"monkey-language/ast"
	"monkey-language/lexer"
	"monkey-language/token"
)

type Parser struct {
	lexer *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
	}
	// read two tokens, so curToken and peekToken are both set
	parser.nextToken()
	parser.nextToken()
	return parser
}

func (this *Parser) nextToken() {
	this.curToken = this.peekToken
	this.peekToken = this.lexer.NextToken()
}

func (this *Parser) expectPeek(tok token.TokenType) bool {
	if this.peekTokenIs(tok) {
		this.nextToken()
		return true
	} else {
		return false
	}
}

func (this *Parser) curTokenIs(tok token.TokenType) bool {
	return this.curToken.Type == tok
}

func (this *Parser) peekTokenIs(tok token.TokenType) bool {
	return this.peekToken.Type == tok
}

func (this *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for this.curTokenIs(token.EOF) == false {
		stmt := this.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		this.nextToken()
	}

	return program
}

func (this *Parser) parseStatement() ast.Statement {
	switch this.curToken.Type {
	case token.LET:
		return this.parseLetStatement()
	default:
		return nil
	}
}

func (this *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: this.curToken,
	}

	if this.expectPeek(token.IDENT) == false {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: this.curToken,
		Value: this.curToken.Literal,
	}

	if this.expectPeek(token.ASSIGN) == false {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for this.curTokenIs(token.SEMICOLON) {
		this.nextToken()
	}

	return stmt
}
