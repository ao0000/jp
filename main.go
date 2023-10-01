package main

import (
	"fmt"
	"github.com/ao0000/jp/lexer"
	"github.com/ao0000/jp/parser"
)

func main() {
	l := lexer.NewLexer(`{}`)
	p := parser.NewParser(l)
	ast := p.Parse()
	fmt.Printf("%+v\n", ast.Literal())
	if len(p.Errors()) != 0 {
		for _, err := range p.Errors() {
			fmt.Println(err.Error())
		}
	}
}
