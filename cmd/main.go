package main

import (
	"log"
	"todo/db"
	"todo/handler"
	"todo/plug"

	"github.com/anthdm/slick"

	_ "todo/pkg/env"
	_ "todo/pkg/log"
	_ "todo/pkg/session"
)

func main() {
	db.Init()

	app := slick.New()
	app.Get("/", handler.HandleAuthPage)
	app.Get("/login", handler.HandleAuthPage)
	app.Post("/login", handler.HandleLoginPost)
	app.Get("/todos", handler.HandleTodosPage, plug.AuthMiddleware)
	app.Get("/todos/:id", handler.HandleTodoRemove, plug.AuthMiddleware)
	app.Post("/todos/add", handler.HandleTodoAdd, plug.AuthMiddleware)
	app.Get("/todos/:id/edit", handler.HandleTodoEditGet, plug.AuthMiddleware)
	app.Post("/todos/edit", handler.HandleTodoEditPost, plug.AuthMiddleware)
	log.Fatal(app.Start())
}
