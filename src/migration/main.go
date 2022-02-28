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

// Run migrate change database schema
func Run(dbDriver string, dbSource string, migrationPath string) {
	db, err := sql.Open(dbDriver, dbSource)
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
		dbDriver, driver)
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
		return
	}
	m.Up()
}
