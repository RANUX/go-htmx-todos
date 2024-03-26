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
	app.Get("/", handler.HandleAuthPage)
	app.Get("/login", handler.HandleAuthPage)
	app.Post("/login", handler.HandleLoginPost)
	app.Get("/todos", handler.HandleTodosPage)
	app.Get("/todos/:id", handler.HandleTodoRemove)
	app.Post("/todos/add", handler.HandleTodoAdd)
	app.Get("/todos/:id/edit", handler.HandleTodoEditGet)
	app.Post("/todos/edit", handler.HandleTodoEditPost)
	log.Fatal(app.Start())
}
