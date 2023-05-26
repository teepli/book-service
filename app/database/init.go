package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() (*sql.DB, error) {
	file, err := os.Create("books.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	return sql.Open("sqlite3", "./books.db")
}

func PrepareTables(db *sql.DB) {
	createBooksTable := `CREATE TABLE if not exists books (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		year INT,
		publisher TEXT,
		description TEXT
	);`

	statement, err := db.Prepare(createBooksTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}
