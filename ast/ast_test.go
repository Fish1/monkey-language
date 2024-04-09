package ast

import (
	"testing"

	"github.com/Fish1/monkey-language/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	programString := program.String()

	if programString != "let myVar = anotherVar;" {
		t.Fatalf("have: \"%s\" , want: \"let myVar = anotherVar;\"", programString)
	}
}
