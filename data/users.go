package data

import (
	"context"
	"errors"
	"time"
	"todo/db"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `bun:"id,pk,autoincrement"`            // Primary key with autoincrement
	UUID      uuid.UUID `bun:"uuid, notnull,unique,type:uuid"` // Unique and not null constraint
	Username  string    `bun:"username,notnull"`               // Not null constraint
	Email     string    `bun:"email,unique,notnull"`           // Unique and not null constraint
	Password  string    `bun:"password,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"` // Use current timestamp as default
}

var _ bun.AfterCreateTableHook = (*User)(nil)

type Users []*User

// UsersAll fetches all users from the database.
func UsersAll() (Users, error) {
	ctx := context.Background()
	var users Users
	err := db.Bun.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UserCreateTable creates the users table in the database.
func UserCreateTable() error {
	ctx := context.Background()
	_, err := db.Bun.NewCreateTable().Model((*User)(nil)).Exec(ctx)
	return err
}

// Make index for uuid user table column
func (*User) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := db.Bun.NewCreateIndex().Model((*User)(nil)).Index("idx_uuid").Column("uuid").Exec(ctx)
	return err
}

// UserDropTable drops the users table in the database.
func UserDropTable() error {
	ctx := context.Background()
	_, err := db.Bun.NewDropTable().Table("users").IfExists().Exec(ctx)
	return err
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Create new user with given password string
func UserAdd(username, email, password string) error {
	newHash, err := HashPassword(password)
	if err != nil {
		return err
	}
	ctx := context.Background()

	if _, err := db.Bun.NewInsert().Model(&User{UUID: uuid.New(), Username: username, Email: email, Password: newHash}).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func UserPasswordMatches(plainText string, user *User) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword): // password does not match
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

// User get user by username
func UserGet(username string) (*User, error) {
	ctx := context.Background()
	user := new(User)
	if err := db.Bun.NewSelect().Model(user).Where("username = ?", username).Scan(ctx); err != nil {
		return nil, err
	}
	return user, nil
}

// User get user by ID uuid
func UserGetByUUID(id uuid.UUID) (*User, error) {
	ctx := context.Background()
	user := new(User)
	if err := db.Bun.NewSelect().Model(user).Where("uuid = ?", id).Scan(ctx); err != nil {
		return nil, err
	}
	return user, nil
}
