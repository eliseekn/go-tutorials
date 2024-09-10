package todolist_test

import (
	"testing"
	"todo-list/todolist"
)

func Test_Can_Add_Todo(t *testing.T) {
	text := "Learn Go"

	todoList, err := todolist.Add(text)
	if err != nil {
		t.Fatal(err)
	}

	if todoList[0].Text != text {
		t.FailNow()
	}
}

func Test_Can_List_Todo(t *testing.T) {
	todoList, err := todolist.List()
	if err != nil {
		t.Fatal(err)
	}

	if len(todoList) == 0 {
		t.FailNow()
	}
}

func Test_Can_Remove_Todo(t *testing.T) {
	text := "Learn Rust"

	_, err := todolist.Add(text)
	if err != nil {
		t.Fatal(err)
	}

	todoList, err := todolist.Remove(2)
	if err != nil {
		t.Fatal(err)
	}

	if len(todoList) > 1 {
		t.FailNow()
	}
}

func Test_Can_Update_Todo(t *testing.T) {
	text := "Learn PHP"

	_, err := todolist.Add(text)
	if err != nil {
		t.Fatal(err)
	}

	todoList, err := todolist.Update(2)
	if err != nil {
		t.Fatal(err)
	}

	if todoList[1].Completed != true {
		t.FailNow()
	}
}
