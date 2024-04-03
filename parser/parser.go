package parser

import (
	"github.com/Fish1/monkey-language/ast"
	"github.com/Fish1/monkey-language/lexer"
	"github.com/Fish1/monkey-language/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	parser := &Parser{l: l}
	parser.NextToken()
	parser.NextToken()
	return parser
}

func (p *Parser) NextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.NextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	if p.currToken.Type == token.LET {
		return p.parseLetStatement()
	}
	return nil
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.currToken,
	}

	if p.peekToken.Type != token.IDENT {
		return nil
	}
	p.NextToken()
	stmt.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	if p.peekToken.Type != token.ASSIGN {
		return nil
	}

	for p.currToken.Type != token.SEMICOLON {
		p.NextToken()
	}

	return stmt
}
