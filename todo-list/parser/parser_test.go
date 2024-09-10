package parser_test

import (
	"testing"
	"todo-list/parser"
)

func Test_Can_Parse_Not_Enough_Args(t *testing.T) {
	args := []string{}

	_, err := parser.Parse(args)
	if err.Error() != "not enough arguments" {
		t.Fatal(err)
	}
}

func Test_Can_Parse_Invalid_Args(t *testing.T) {
	args := []string{"", "-w"}

	_, err := parser.Parse(args)
	if err.Error() != "invalid arguments" {
		t.Fatal(err)
	}
}

func Test_Can_Parse_Valid_Args(t *testing.T) {
	args := []string{"", "-l"}

	_, err := parser.Parse(args)
	if err != nil {
		t.Fatal(err)
	}
}
