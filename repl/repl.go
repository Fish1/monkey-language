package repl

import (
	"fmt"

	"github.com/Fish1/monkey-language/lexer"
	"github.com/Fish1/monkey-language/token"
)

func Start() {
	var input string
	for {
		fmt.Scanln(&input)
		lexer := lexer.New(input)
		for {
			t := lexer.NextToken()
			if t.Type == token.EOF || t.Type == token.ILLEGAL {
				break
			}
			fmt.Printf("{type: %s, literal: %s}\n", t.Type, t.Literal)
		}
	}
}
