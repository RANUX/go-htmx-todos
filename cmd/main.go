package main

import (
	"log"
	"todo/db"
	"todo/handler"
	plug "todo/middleware"

	"github.com/anthdm/slick"

	_ "todo/configs"
)

func main() {
	db.Init()

	app := slick.New()
	app.Get("/", handler.HandleLandingPage)
	app.Get("/login", handler.HandleLoginPage)
	app.Post("/login", handler.HandleLoginPost)
	app.Get("/logout", handler.HandleLogoutGet)
	app.Get("/signup", handler.HandleRegisterPage)
	app.Get("/profile", handler.HandleProfilePage, plug.AuthMiddleware)
	app.Post("/profile/save", handler.HandleProfileSave, plug.AuthMiddleware)
	app.Get("/todos", handler.HandleTodoPage, plug.AuthMiddleware)
	app.Get("/todos/:id", handler.HandleTodoRemove, plug.AuthMiddleware)
	app.Post("/todos/add", handler.HandleTodoAdd, plug.AuthMiddleware)
	app.Get("/todos/:id/edit", handler.HandleTodoEditGet, plug.AuthMiddleware)
	app.Post("/todos/edit", handler.HandleTodoEditPost, plug.AuthMiddleware)
	app.Get("/public/*filepath", handler.HandlePublicStaticFile)
	app.Get("/alerts", handler.HandleAlerts)
	log.Fatal(app.Start())
}
