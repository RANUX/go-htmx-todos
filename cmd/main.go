package main

import (
	"log"
	"todo/db"
	"todo/handler"

	"github.com/anthdm/slick"
)

func main() {
	db.Init()

	app := slick.New()
	app.Get("/", handler.HandleTodosIndex)
	app.Get("/todos/:id", handler.HandleTodoRemove)
	app.Post("/todos/add", handler.HandleTodoAdd)
	app.Get("/todos/:id/edit", handler.HandleTodoEditGet)
	app.Post("/todos/edit", handler.HandleTodoEditPost)

	log.Fatal(app.Start())
}