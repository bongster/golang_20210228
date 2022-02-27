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

func Run(dbPath string, migrationPath string) {
	db, err := sql.Open("postgres", dbPath)
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
		migrationPath,
		"postgres", driver)
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
		return
	}
	m.Up()
}
