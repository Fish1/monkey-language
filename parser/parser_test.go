package parser

import (
	"fmt"
	"testing"

	"github.com/Fish1/monkey-language/ast"
	"github.com/Fish1/monkey-language/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Errorf("ParseProgram() returned nil")
	}

	if p.HasErrors() {
		p.PrintErrors()
		t.Errorf("parsing errors encountered")
	}

	if len(program.Statements) != 3 {
		t.Errorf("Failed to identify three statements")
	}

	tests := []struct{ expectedIdentifier string }{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for index, test := range tests {
		statement := program.Statements[index]
		if testLetStatement(t, statement, test.expectedIdentifier) == false {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("token literal not 'let'. got %s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if ok == false {
		t.Errorf("s not *ast.LetStatement. got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not %s, got %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not %s, got %s",
			name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 123;
	return x;
	return mything();
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Errorf("ParseProgram() returned nil")
	}

	if p.HasErrors() {
		p.PrintErrors()
		t.Errorf("parsing errors encountered")
	}

	if len(program.Statements) != 3 {
		t.Errorf("failed to identify three statements")
	}

	for index := range 3 {
		statement := program.Statements[index]
		if testReturnStatement(t, statement) == false {
			return
		}
	}
}

func testReturnStatement(t *testing.T, s ast.Statement) bool {
	if s.TokenLiteral() != "return" {
		t.Errorf("token literal not 'return'. got %s", s.TokenLiteral())
		return false
	}

	_, ok := s.(*ast.ReturnStatement)
	if ok == false {
		t.Errorf("s not *ast.ReturnStatement. got %T", s)
		return false
	}

	return true
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	p.PrintErrors()

	if len(program.Statements) != 1 {
		t.Errorf("program did not identify 1 statements")
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if ok == false {
		t.Errorf("statement is not an expression statement")
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if ok == false {
		t.Errorf("expression is not an identifier")
	}

	if ident.Value != "foobar" {
		t.Errorf("identifier is not foobar")
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() not %s. got %s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	p.PrintErrors()

	if len(program.Statements) != 1 {
		t.Errorf("program did not identify 1 statements")
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if ok == false {
		t.Errorf("statement is not an expression statement")
	}

	ident, ok := stmt.Expression.(*ast.IntegerLiteral)
	if ok == false {
		t.Errorf("expression is not an integer literal")
	}

	if ident.Value != 5 {
		t.Errorf("value is not 5")
	}

	if ident.TokenLiteral() != "5" {
		t.Errorf("ident.TokenLiteral() not %s. got %s", "5", ident.TokenLiteral())
	}
}

func TestNegationExpression(t *testing.T) {
	input := "-5;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	p.PrintErrors()

	if len(program.Statements) != 1 {
		t.Errorf("program did not identify 1 statements")
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if ok == false {
		t.Errorf("statement is not an expression statement")
	}

	negation, ok := stmt.Expression.(*ast.Negation)
	if ok == false {
		t.Errorf("expression is not a negation")
	}

	literal, ok := negation.Expression.(*ast.IntegerLiteral)
	if ok == false {
		t.Errorf("negation expression is not an integer literal")
	}

	if literal.Value != 5 {
		t.Errorf("literal is not %d. got %d", 5, literal.Value)
	}
}
