package books

import "github.com/gin-gonic/gin"

type BookService interface {
	createBook(bq BookQuery) (int, error)
	getAllBooks(params BookFilterParams) ([]BookResponse, error)
	getBook(id string) (BookResponse, error)
	deleteBook(id string) error
}

type BookRepository interface {
	createBook(bq BookQuery) (int, error)
	getAllBooks(params BookFilterParams) ([]BookResponse, error)
	getBook(id string) (BookResponse, error)
	deleteBook(id string) (int64, error)
}

type BookApi interface {
	createBook(c *gin.Context)
	getAllBooks(c *gin.Context)
	getBook(c *gin.Context)
	deleteBook(c *gin.Context)
}
