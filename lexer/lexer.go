package lexer

import (
	"github.com/Fish1/monkey-language/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	lexer.skipWhitespace()
	var t token.Token
	switch lexer.ch {
	case '=':
		if lexer.peekChar() == '=' {
			t = newToken(token.EQUAL, "==")
			lexer.readChar()
		} else {
			t = newToken(token.ASSIGN, lexer.ch)
		}
	case '!':
		if lexer.peekChar() == '=' {
			t = newToken(token.NOTEQUAL, "!=")
			lexer.readChar()
		} else {
			t = newToken(token.EXCLAMATION, lexer.ch)
		}
	case '+':
		t = newToken(token.PLUS, lexer.ch)
	case '(':
		t = newToken(token.LPAREN, lexer.ch)
	case ')':
		t = newToken(token.RPAREN, lexer.ch)
	case '{':
		t = newToken(token.LBRACE, lexer.ch)
	case '}':
		t = newToken(token.RBRACE, lexer.ch)
	case ',':
		t = newToken(token.COMMA, lexer.ch)
	case ';':
		t = newToken(token.SEMICOLON, lexer.ch)
	case '-':
		t = newToken(token.MINUS, lexer.ch)
	case '*':
		t = newToken(token.ASTERISK, lexer.ch)
	case '/':
		t = newToken(token.BSLASH, lexer.ch)
	case '<':
		t = newToken(token.LESSTHAN, lexer.ch)
	case '>':
		t = newToken(token.GREATERTHAN, lexer.ch)
	case 0:
		t = newToken(token.EOF, "")
	default:
		if isLetter(lexer.ch) {
			literal := lexer.readIdentifier()
			ident := token.LookupIdent(literal)
			return newToken(ident, literal)
		} else if isDigit(lexer.ch) {
			literal := lexer.readDigit()
			return newToken(token.INT, literal)
		} else {
			return newToken(token.ILLEGAL, lexer.ch)
		}
	}
	lexer.readChar()
	return t
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) skipWhitespace() {
	for isWhitespace(lexer.ch) {
		lexer.readChar()
	}
}

func (lexer *Lexer) readIdentifier() string {
	res := ""
	for isLetter(lexer.ch) {
		res = res + string(lexer.ch)
		lexer.readChar()
	}
	return res
}

func (lexer *Lexer) readDigit() string {
	res := ""
	for isDigit(lexer.ch) {
		res = res + string(lexer.ch)
		lexer.readChar()
	}
	return res
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch == '_')
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func newToken[LITERAL byte | string](t token.TokenType, c LITERAL) token.Token {
	return token.Token{
		Type:    t,
		Literal: string(c),
	}
}
