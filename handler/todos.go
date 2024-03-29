package handler

import (
	"fmt"
	"net/http"
	"todo/data"
	"todo/view/todos"

	"github.com/anthdm/slick"
)

func HandleTodoPage(c *slick.Context) error {
	todoList, err := data.TodosAll()
	if err != nil {
		fmt.Printf("Error fetching todos: %v\n", err)
		return err
	}
	TodoListProps := todos.TodoListProps{
		Todos: todoList,
	}

	return c.Render(todos.TodoPage(TodoListProps))
}

// handle remove todo
func HandleTodoRemove(c *slick.Context) error {
	id := c.Param("id")
	data.TodoRemove(id)

	// return empty string
	return c.Text(http.StatusOK, "")
}

// hande add new todo
func HandleTodoAdd(c *slick.Context) error {
	text := c.FormValue("text")
	data.TodoAdd(text)

	todoList, err := data.TodosAll()
	if err != nil {
		fmt.Printf("Error fetching todos: %v\n", err)
		return err
	}
	TodoListProps := todos.TodoListProps{
		Todos: todoList,
	}

	return c.Render(todos.TodoList(TodoListProps))
}

func HandleTodoEditGet(c *slick.Context) error {
	id := c.Param("id")
	todo, err := data.TodoGet(id)
	if err != nil {
		fmt.Printf("Error fetching todo: %v\n", err)
		return err
	}
	props := todos.InputFormProps{
		Todo:   *todo,
		Action: "edit",
	}
	return c.Render(todos.InputForm(props))
}

func HandleTodoEditPost(c *slick.Context) error {
	id := c.FormValue("id")
	text := c.FormValue("text")
	data.TodoUpdate(id, text)

	todoList, err := data.TodosAll()
	if err != nil {
		fmt.Printf("Error fetching todos: %v\n", err)
		return err
	}
	TodoListProps := todos.TodoListProps{
		Todos: todoList,
	}

	return c.Render(todos.TodoContainer(TodoListProps))
}
