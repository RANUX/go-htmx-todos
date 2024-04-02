package handler

import (
	"fmt"
	"net/http"
	"todo/data"
	"todo/view/todos"

	"github.com/anthdm/slick"
	"github.com/sirupsen/logrus"
)

func HandleTodoPage(c *slick.Context) error {
	user, err := data.UserGetUser(c)
	if err != nil {
		logrus.Error(err)
		return err
	}
	todoList, err := data.TodoAllByUser(user)
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

	user, err := data.UserGetUser(c)
	if err != nil {
		logrus.Error(err)
		return err
	}

	_, err = data.TodoCreate(text, user)
	if err != nil {
		logrus.Error(err)
		return err
	}

	todoList, err := data.TodoAllByUser(user)
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
	todo, err := data.TodoGetById(id)
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
	user, err := data.UserGetUser(c)
	if err != nil {
		logrus.Error(err)
		return err
	}
	id := c.FormValue("id")
	text := c.FormValue("text")

	data.TodoUpdate(id, text)

	todoList, err := data.TodoAllByUser(user)
	if err != nil {
		fmt.Printf("Error fetching todos: %v\n", err)
		return err
	}
	TodoListProps := todos.TodoListProps{
		Todos: todoList,
	}

	return c.Render(todos.TodoContainer(TodoListProps))
}
