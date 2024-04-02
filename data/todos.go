package data

import (
	"context"
	"todo/db"

	"github.com/google/uuid"
)

type Todo struct {
	ID     int64     `bun:",pk,autoincrement"`
	UUID   uuid.UUID `bun:"uuid,unique,type:uuid"` // Unique and not null constraint
	Text   string
	Done   bool // default false
	UserId int64
}

type Todos []*Todo

// TodoAll fetches all todos from the database.
func TodoAll() (Todos, error) {
	ctx := context.Background()
	var todos Todos
	err := db.Bun.NewSelect().Model(&todos).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// fetch all users todos
func TodoAllByUser(user *User) (Todos, error) {
	ctx := context.Background()
	var todos Todos
	err := db.Bun.NewSelect().Model(&todos).Where("user_id = ?", user.ID).Scan(ctx)
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

// Todo removes a todo from the database.
func TodoRemove(id string) error {
	ctx := context.Background()
	if _, err := db.Bun.NewDelete().Model(&Todo{}).Where("id = ?", id).Exec(ctx); err != nil {
		return err
	}
	return nil
}

// TodoGetById fetches a todo from the database.
func TodoGetById(id string) (*Todo, error) {
	ctx := context.Background()
	todo := new(Todo)
	if err := db.Bun.NewSelect().Model(todo).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return todo, nil
}

// TodoGetByUUID fetches a todo from the database.
func TodoGetByUUID(uuid uuid.UUID) (*Todo, error) {
	ctx := context.Background()
	todo := new(Todo)
	if err := db.Bun.NewSelect().Model(todo).Where("uuid = ?", uuid).Scan(ctx); err != nil {
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

// Create new todo
func TodoCreate(text string, user *User) (uuid.UUID, error) {
	ctx := context.Background()
	uuid := uuid.New()
	_, err := db.Bun.NewInsert().Model(&Todo{UUID: uuid, Text: text, UserId: user.ID}).Exec(ctx)
	if err != nil {
		return uuid, err
	}
	return uuid, nil
}
