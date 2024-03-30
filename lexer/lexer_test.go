package lexer

import (
	"testing"

	"github.com/Fish1/monkey-language/token"
)

func TestNextToken(t *testing.T) {
	input := `let x = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
-/*<>
! != ==
let result = add(five, ten);`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.MINUS, "-"},
		{token.BSLASH, "/"},
		{token.ASTERISK, "*"},
		{token.LESSTHAN, "<"},
		{token.GREATERTHAN, ">"},
		{token.EXCLAMATION, "!"},
		{token.NOTEQUAL, "!="},
		{token.EQUAL, "=="},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)
	for index, tt := range tests {
		token := lexer.NextToken()

		if tt.expectedType != token.Type {
			t.Fatalf("want: %s , have: %s  || %d", tt.expectedType, token.Type, index)
		}

		if tt.expectedLiteral != token.Literal {
			t.Fatalf("want: %s , have: %s || %d", tt.expectedLiteral, token.Literal, index)
		}
	}
}
