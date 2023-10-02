package app

import (
	"fmt"
	"github.com/ao0000/jp/ast"
	"github.com/ao0000/jp/lexer"
	"github.com/ao0000/jp/parser"
)

func App(str string) {
	l := lexer.NewLexer(str)
	p := parser.NewParser(l)
	node, errs := p.Parse()
	if len(errs) != 0 {
		for _, err := range p.Errors() {
			fmt.Println(err.Error())
		}
	}
	run(node.Value())
}

func run(node ast.Value) ast.Value {
	switch node.(type) {
	case *ast.Object:
		if n, ok := node.(*ast.Object); ok {
			printObject(n)
		}
		return nil
	case *ast.Array:
		panic("not implementation : *ast.Array case")
		return nil
	default:
		return node
	}
}

func printObject(object *ast.Object) {
	rng := len(object.Keys)
	if rng < len(object.Values) {
		rng = len(object.Values)
	}

	if rng == 0 {
		fmt.Println("{}")
		return
	}

	for i := 0; i < rng; i++ {
		k := run(object.Keys[i])
		v := run(object.Values[i])
		fmt.Printf("{%s : %s}", k, v)
	}
}
