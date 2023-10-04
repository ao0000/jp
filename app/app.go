package app

import (
	"errors"
	"fmt"
	"github.com/ao0000/jp/ast"
	"github.com/ao0000/jp/parser"
)

func Run(p parser.Parser) error {
	node, errs := p.Parse()
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	printNode(node.Value(), 0, true)
	return nil
}

func printNode(node ast.Value, nest int, nested bool) ast.Value {
	switch node.(type) {
	case *ast.Object:
		if n, ok := node.(*ast.Object); ok {
			if len(n.Keys) == 0 {
				fmt.Print("{}")
				return nil
			}
			if nested {
				for i := 0; i < nest; i++ {
					fmt.Print("  ")
				}
			}
			fmt.Println("{")

			printObject(n, nest)

			for i := 0; i < nest; i++ {
				fmt.Print("  ")
			}
			fmt.Print("}")
		}
		return nil
	case *ast.Array:
		if n, ok := node.(*ast.Array); ok {
			if len(n.Values) == 0 {
				fmt.Print("[]")
				return nil
			}
			if nested {
				for i := 0; i < nest; i++ {
					fmt.Print("  ")
				}
			}
			fmt.Println("[")

			printArray(n, nest)

			for i := 0; i < nest; i++ {
				fmt.Print("  ")
			}
			fmt.Print("]")
		}
		return nil
	default:
		if nested == true {
			for i := 0; i < nest; i++ {
				fmt.Print("  ")
			}
		}
		fmt.Printf("%s", node.Literal())
		return node
	}
}

func printObject(object *ast.Object, nest int) {
	rng := len(object.Keys)
	if rng < len(object.Values) {
		rng = len(object.Values)
	}

	if rng == 0 {
		return
	}

	for i := 0; i < rng; i++ {
		if i != 0 {
			fmt.Println(",")
		}
		printNode(object.Keys[i], nest+1, true)
		fmt.Printf(": ")
		printNode(object.Values[i], nest+1, false)
		if i == rng-1 {
			fmt.Println()
		}
	}
}

func printArray(array *ast.Array, nest int) {
	for i, v := range array.Values {
		if i != 0 {
			fmt.Println(",")
		}
		printNode(v, nest+1, true)
		if i == len(array.Values)-1 {
			fmt.Println()
		}
	}
}
