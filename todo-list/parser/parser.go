package parser

import (
	"errors"
	"fmt"
	"strconv"
	"todo-list/todolist"
	"todo-list/types"
)

func Parse(args []string) ([]types.ToDo, error) {
	if len(args) < 2 {
		return nil, errors.New("not enough arguments")
	}

	if args[1] == "--list" || args[1] == "-l" {
		todoList, err := todolist.List()
		if err != nil {
			return nil, err
		}

		for _, todo := range todoList {
			completed := "X"

			if !todo.Completed {
				completed = ""
			}

			fmt.Println("[" + completed + "] " + todo.Text + " (" + strconv.Itoa(todo.Id) + ")")
		}

		return todoList, nil
	}

	if args[1] == "--add" || args[1] == "-a" {
		todoList, err := todolist.Add(args[2])
		if err != nil {
			return nil, err
		}

		return todoList, nil
	}

	if args[1] == "--remove" || args[1] == "-r" {
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return nil, err
		}

		todoList, err := todolist.Remove(id)
		if err != nil {
			return nil, err
		}

		for _, todo := range todoList {
			completed := "X"

			if !todo.Completed {
				completed = ""
			}

			fmt.Println("[" + completed + "] " + todo.Text + " (" + strconv.Itoa(todo.Id) + ")")
		}

		return todoList, nil
	}

	if args[1] == "--update" || args[1] == "-u" {
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return nil, err
		}

		todoList, err := todolist.Update(id)
		if err != nil {
			return nil, err
		}

		for _, todo := range todoList {
			completed := "X"

			if !todo.Completed {
				completed = ""
			}

			fmt.Println("[" + completed + "] " + todo.Text + " (" + strconv.Itoa(todo.Id) + ")")
		}

		return todoList, nil
	}

	if args[1] == "--clear" || args[1] == "-c" {
		return nil, todolist.Clear()
	}

	return nil, errors.New("invalid arguments")
}
