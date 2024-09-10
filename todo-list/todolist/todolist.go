package todolist

import (
	"todo-list/io"
	"todo-list/types"
)

func List() ([]types.ToDo, error) {
	return io.ReadFile()
}

func Add(text string) ([]types.ToDo, error) {
	todoList, err := List()
	if err != nil {
		return nil, err
	}

	todo := types.ToDo{
		Id:        len(todoList) + 1,
		Text:      text,
		Completed: false,
	}

	todoList = append(todoList, todo)

	err = io.SaveFile(todoList)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func Remove(id int) ([]types.ToDo, error) {
	todoList, err := List()
	if err != nil {
		return nil, err
	}

	var newTodoList []types.ToDo

	for _, todo := range todoList {
		if todo.Id != id {
			newTodoList = append(newTodoList, todo)
		}
	}

	err = io.SaveFile(newTodoList)
	if err != nil {
		return nil, err
	}

	return newTodoList, nil
}

func Update(id int) ([]types.ToDo, error) {
	todoList, err := List()
	if err != nil {
		return nil, err
	}

	for i, todo := range todoList {
		if todo.Id == id {
			todoList[i].Completed = !todoList[i].Completed
		}
	}

	err = io.SaveFile(todoList)
	if err != nil {
		return nil, err
	}

	return todoList, nil
}

func Clear() error {
	return io.DeleteFile()
}
