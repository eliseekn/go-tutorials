package main

import (
	"fmt"
	"os"
	"todo-list/parser"
)

func main() {
	_, err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
