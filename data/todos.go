package data

import (
	"context"
	"todo/db"
)

type Todo struct {
	ID   int64 `bun:",pk,autoincrement"`
	Text string
	Done bool // default false
}

type Todos []*Todo

// TodosAll fetches all todos from the database.
func TodosAll() (Todos, error) {
	ctx := context.Background()
	var todos Todos
	err := db.Bun.NewSelect().Model(&todos).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// TodoCreateTable creates the todos table in the database.
func TodoCreateTable() error {
	ctx := context.Background()
	_, err := db.Bun.NewCreateTable().Model((*Todo)(nil)).Exec(ctx)
	return err
}

// TodoDropTable drops the todos table in the database.
func TodoDropTable() error {
	ctx := context.Background()
	_, err := db.Bun.NewDropTable().Table("todos").IfExists().Exec(ctx)
	return err
}

// Todo inserts a new todo into the database.
func TodoAdd(text string) error {
	ctx := context.Background()
	if _, err := db.Bun.NewInsert().Model(&Todo{Text: text}).Exec(ctx); err != nil {
		return err
	}
	return nil
}

// Todo removes a todo from the database.
func TodoRemove(id string) error {
	ctx := context.Background()
	if _, err := db.Bun.NewDelete().Model(&Todo{}).Where("id = ?", id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

// TodoGet fetches a todo from the database.
func TodoGet(id string) (*Todo, error) {
	ctx := context.Background()
	todo := new(Todo)
	if err := db.Bun.NewSelect().Model(todo).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return todo, nil
}

// TodoUpdate updates a todo in the database.
func TodoUpdate(id, text string) error {
	ctx := context.Background()
	if _, err := db.Bun.NewUpdate().Model(&Todo{}).Where("id = ?", id).Set("text = ?", text).Exec(ctx); err != nil {
		return err
	}
	return nil
}
