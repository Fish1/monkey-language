package lexer

import (
	"monkey-language/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) NextToken() token.Token {
	var current_token token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		current_token = newToken(token.ASSIGN, l.ch)
	case ';':
		current_token = newToken(token.SEMICOLON, l.ch)
	case '(':
		current_token = newToken(token.LPAREN, l.ch)
	case ')':
		current_token = newToken(token.RPAREN, l.ch)
	case ',':
		current_token = newToken(token.COMMA, l.ch)
	case '+':
		current_token = newToken(token.PLUS, l.ch)
	case '{':
		current_token = newToken(token.LBRACE, l.ch)
	case '}':
		current_token = newToken(token.RBRACE, l.ch)
	case 0:
		current_token.Literal = ""
		current_token.Type = token.EOF
	default:
		if isLetter(l.ch) {
			current_token.Literal = l.readIdentifier()
			current_token.Type = token.LookupIdent(current_token.Literal)
			return current_token
		} else if isInteger(l.ch) {
			current_token.Literal = l.readInteger()
			current_token.Type = token.INT
			return current_token
		} else {
			current_token = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return current_token
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readInteger() string {
	position := l.position
	for isInteger(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isInteger(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
