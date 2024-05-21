package migrator

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"log"
)

func Run(connStr, migrDir string) {
	driver := "postgres"

	db, err := sqlx.Open(driver, connStr)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer db.Close()

	if err := goose.SetDialect(driver); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, migrDir); err != nil {
		panic(err)
	}
}
