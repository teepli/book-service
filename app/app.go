package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/teepli/book-service/app/books"
)

type App struct {
	DB *sql.DB
}

func NewApp(db *sql.DB) App {
	return App{db}
}

func (a *App) Initialize() {
	r := gin.Default()
	books.NewBookRoutes(r)
	r.Run(":9000")
}
