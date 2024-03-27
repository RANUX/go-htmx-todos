package main

import (
	"log"
	"todo/db"
	"todo/handler"
	"todo/middleware"

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
	app.Get("/todos", handler.HandleTodosPage, middleware.AuthMiddleware)
	app.Get("/todos/:id", handler.HandleTodoRemove, middleware.AuthMiddleware)
	app.Post("/todos/add", handler.HandleTodoAdd, middleware.AuthMiddleware)
	app.Get("/todos/:id/edit", handler.HandleTodoEditGet, middleware.AuthMiddleware)
	app.Post("/todos/edit", handler.HandleTodoEditPost, middleware.AuthMiddleware)
	log.Fatal(app.Start())
}
