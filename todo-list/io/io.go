package io

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"todo-list/types"
)

func Filename() string {
	dir, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}

	filename := dir + "/todos.txt"

	_, err = os.OpenFile(filename, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	return filename
}

func ReadFile() ([]types.ToDo, error) {
	filename := Filename()

	data, err := os.ReadFile(filename)
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
	filename := Filename()

	for _, todo := range todoList {
		line := strconv.Itoa(todo.Id) + "," + todo.Text + "," + strconv.FormatBool(todo.Completed)
		lines = append(lines, line)
	}

	data := strings.Join(lines, "\n")

	return os.WriteFile(filename, []byte(data), 0644)
}

func DeleteFile() error {
	filename := Filename()
	return os.Remove(filename)
}
