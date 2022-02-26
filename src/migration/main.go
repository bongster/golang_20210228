package migration

import (
	"database/sql"
	"os"

	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Run() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@db:5432/payments?sslmode=disable")
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
		return
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migration",
		"postgres", driver)
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
		return
	}
	err = m.Up()
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
		return
	}
}
