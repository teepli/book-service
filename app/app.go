package app

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/teepli/book-service/app/books"
	"github.com/teepli/book-service/app/middleware"
	"github.com/teepli/book-service/app/tools"
)

type App struct {
	DB *sql.DB
}

func NewApp(db *sql.DB) App {
	return App{db}
}

func (a *App) Initialize() {
	router := gin.New()
	router.Use(middleware.RequestIdInjector())
	router.Use(gin.Recovery())

	tools.RegisterCustomValidators()

	repo := books.NewRepository(a.DB)
	service := books.NewService(repo)
	books.NewApi(router, service)

	router.Run(":9000")
}
