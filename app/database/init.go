package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

type DBOptions struct {
	MigrationSource string
	DBPath          string
}

func InitDatabase(opts DBOptions) (*sql.DB, error) {
	con, err := openConnection(opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = migrateSQLite(con, opts.MigrationSource)
	if err != nil {
		log.Fatal(err.Error())
	}

	return con, nil
}

func openConnection(opts DBOptions) (*sql.DB, error) {
	file, err := os.Create(opts.DBPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	return sql.Open("sqlite3", opts.DBPath)
}

func migrateSQLite(db *sql.DB, migrationSource string) error {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationSource,
		"sqlite", driver)

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	return nil
}
