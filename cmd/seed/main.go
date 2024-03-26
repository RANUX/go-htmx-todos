package main

import (
	"todo/data"
	"todo/db"
)

func AddTodos() error {
	data.TodoAdd("Buy milk")
	data.TodoAdd("Buy eggs")
	data.TodoAdd("Buy bread")
	return nil
}

func AddUsers() error {
	data.UserAdd("admin", "admin@admin.com", "admin")
	data.UserAdd("user", "user@user.com", "user")
	return nil
}

func main() {
	db.Init()
	AddTodos()
	AddUsers()
}
