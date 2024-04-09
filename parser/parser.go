package parser

import (
	"errors"
	"fmt"

	"github.com/Fish1/monkey-language/ast"
	"github.com/Fish1/monkey-language/lexer"
	"github.com/Fish1/monkey-language/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	l      *lexer.Lexer
	Errors []string

	currToken token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	parser := &Parser{l: l}
	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.registerPrefix(token.IDENT, parser.parseIdentifier) // when an IDENT is encouterd then parse it
	parser.NextToken()
	parser.NextToken()
	return parser
}

func (p *Parser) NextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.NextToken()
		return true
	}
	s := fmt.Sprintf("expect next TOKEN: %s found TOKEN: %s instead", t, p.peekToken.Type)
	p.Errors = append(p.Errors, s)
	return false
}

func (p *Parser) HasErrors() bool {
	return len(p.Errors) > 0
}

func (p *Parser) PrintErrors() {
	if len(p.Errors) == 0 {
		return
	}
	s := ""

	for i, error := range p.Errors {
		s = fmt.Sprintf("%s%d: %s\n", s, i+1, error)
	}

	fmt.Println(s)
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	for p.currToken.Type != token.EOF {
		if stmt, err := p.parseStatement(); err == nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.NextToken()
	}

	return program
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := &ast.LetStatement{
		Token: p.currToken,
	}

	if p.expectPeek(token.IDENT) == false {
		return nil, errors.New("expected ident")
	}

	stmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if p.expectPeek(token.ASSIGN) == false {
		return nil, errors.New("expected assign")
	}

	for p.currToken.Type != token.SEMICOLON {
		p.NextToken()
	}

	return stmt, nil
}

func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := &ast.ReturnStatement{
		Token: p.currToken,
	}

	for p.currToken.Type != token.SEMICOLON {
		p.NextToken()
	}

	return stmt, nil
}

func (p *Parser) parseExpressionStatement() (*ast.ExpressionStatement, error) {
	stmt := &ast.ExpressionStatement{
		Token:      p.currToken,
		Expression: p.parseExpression(LOWEST),
	}

	if p.peekToken.Type == token.SEMICOLON {
		p.NextToken()
	}

	return stmt, nil
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		return nil
	}
	leftExpression := prefix()
	return leftExpression
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
