package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/teepli/book-service/app"
	"github.com/teepli/book-service/app/database"
)

func main() {
	opts := database.DBOptions{
		MigrationSource: "file://migrations",
		DBPath:          "./book.db",
	}

	db, err := database.InitDatabase(opts)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	a := app.NewApp(db)
	a.Initialize()
}
