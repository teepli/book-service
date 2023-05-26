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
	router := gin.Default()

	repo := books.NewRepository(a.DB)
	service := books.NewService(repo)
	books.NewApi(router, service)
	router.Run(":9000")
}
