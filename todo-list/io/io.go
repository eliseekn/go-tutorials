package io

import (
	"os"
	"strconv"
	"strings"
	"todo-list/types"
)

func ReadFile() ([]types.ToDo, error) {
	data, err := os.ReadFile("/tmp/todos.txt")
	if err != nil {
		return nil, err
	}

	var todoList []types.ToDo
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ",")

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		text := parts[1]

		completed, err := strconv.ParseBool(parts[2])
		if err != nil {
			continue
		}

		todoList = append(todoList, types.ToDo{
			Id:        id,
			Text:      text,
			Completed: completed,
		})
	}

	return todoList, nil
}

func SaveFile(todoList []types.ToDo) error {
	var lines []string

	for _, todo := range todoList {
		line := strconv.Itoa(todo.Id) + "," + todo.Text + "," + strconv.FormatBool(todo.Completed)
		lines = append(lines, line)
	}

	data := strings.Join(lines, "\n")

	return os.WriteFile("/tmp/todos.txt", []byte(data), 0644)
}

func DeleteFile() error {
	return os.Remove("/tmp/todos.txt")
}
