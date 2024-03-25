package main

import (
	"todo/data"
	"todo/db"
)

func main() {
	db.Init()
	// add todos to database
	data.TodoAdd("Buy milk")
	data.TodoAdd("Buy eggs")
	data.TodoAdd("Buy bread")

}
