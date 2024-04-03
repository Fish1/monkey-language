package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Fish1/monkey-language/lexer"
	"github.com/Fish1/monkey-language/token"
)

func Start() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		lexer := lexer.New(line)
		for {
			t := lexer.NextToken()
			fmt.Printf("{type: %s, literal: %s}\n", t.Type, t.Literal)
			if t.Type == token.EOF || t.Type == token.ILLEGAL {
				break
			}
		}
	}
}
