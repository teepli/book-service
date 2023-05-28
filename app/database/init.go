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

func InitDatabase(migrationSource string) (*sql.DB, error) {
	con, err := openConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = migrateSQLite(con, migrationSource)
	if err != nil {
		log.Fatal(err.Error())
	}

	return con, nil
}

func openConnection() (*sql.DB, error) {
	file, err := os.Create("books.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	return sql.Open("sqlite3", "./books.db")
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
