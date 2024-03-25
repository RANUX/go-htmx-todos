package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

var Bun *bun.DB

func CreateDatabase() (*sql.DB, error) {
	// godotenv.Load()
	// var (
	// 	dbname = os.Getenv("DB_NAME")
	// 	dbuser = os.Getenv("DB_USER")
	// 	dbpass = os.Getenv("DB_PASS")
	// 	dbhost = os.Getenv("DB_HOST")
	// 	uri    = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbuser, dbpass, dbhost, dbname)
	// )
	// sqlite, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	sqlite, err := sql.Open(sqliteshim.ShimName, "db.sqlite")
	if err != nil {
		panic(err)
	}
	sqlite.SetMaxOpenConns(1)
	return sqlite, nil
}

func Init() error {
	sqlite, err := CreateDatabase()
	if err != nil {
		return err
	}
	Bun = bun.NewDB(sqlite, sqlitedialect.New())
	Bun.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return nil
}
