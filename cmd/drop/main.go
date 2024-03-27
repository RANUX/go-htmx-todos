package main

import (
	"todo/data"
	"todo/db"
)

func main() {
	db.Init()
	data.TodoDropTable()
	data.UserDropTable()
}
