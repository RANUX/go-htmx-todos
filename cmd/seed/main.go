package main

import (
	"todo/data"
	"todo/db"

	"github.com/sirupsen/logrus"
)

func AddTodos() {
	// get user from db by username
	user, err := data.UserGet("admin")
	if err != nil {
		logrus.Error(err)
		return
	}
	// create new todo
	_, err = data.TodoCreate("Buy milk", user)
	if err != nil {
		logrus.Error(err)
		return
	}
	_, err = data.TodoCreate("Buy bread", user)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func AddUsers() error {
	data.UserAdd("admin", "admin@admin.com", "admin")
	data.UserAdd("user", "user@user.com", "user")
	return nil
}

func main() {
	db.Init()
	AddUsers()
	AddTodos()
}
