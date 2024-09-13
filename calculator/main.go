package main

import (
	"calculator/operator"
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
		operator.Sum(args)
	case "-":
		operator.Substract(args)
	case "x":
		operator.Multiply(args)
	case "/":
		operator.Divide(args)
	default:
		log.Fatalf("invalid operator: %v", args[2])
	}
}
