package main

import (
	"github.com/ao0000/jp/cmd"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		log.Fatal("should set input JSON")
	}

	err := cmd.Execute(args[1])
	if err != nil {
		log.Fatal(err)
	}
}
