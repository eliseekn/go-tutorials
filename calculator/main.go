package main

import (
	"calculator/operator"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 4 {
		log.Fatal("not enough arguments")
	}

	switch args[2] {
	case "+":
		result := operator.Sum(args)
		fmt.Println(result)
	case "-":
		result := operator.Substract(args)
		fmt.Println(result)
	case "x":
		result := operator.Multiply(args)
		fmt.Println(result)
	case "/":
		result := operator.Divide(args)
		fmt.Println(result)
	default:
		log.Fatalf("invalid operator: %v", args[2])
	}
}
