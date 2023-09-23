package main

import (
	"fmt"
	"github.com/ao0000/jp/lexer"
	"github.com/ao0000/jp/token"
)

func main() {
	l := lexer.NewLexer("{}")
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		fmt.Println(tok)
	}
}
