package cmd

import (
	"fmt"
	"github.com/ao0000/jp/app"
	"github.com/ao0000/jp/lexer"
	"github.com/ao0000/jp/parser"
)

func Execute(input string) error {
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)

	err := app.Run(p)
	if err != nil {
		return fmt.Errorf("failed to application: %w", err)
	}

	return nil
}
