package books

import "github.com/gin-gonic/gin"

type BookService interface {
	createBook(bq BookQuery) (int, error)
	getAllBooks(params BookFilterParams) ([]BookResponse, error)
	getBook(id int) (BookResponse, error)
	deleteBook(id int) error
}

type BookRepository interface {
	createBook(bq BookQuery) (int, error)
	getAllBooks(params BookFilterParams) ([]BookResponse, error)
	getBook(id int) (BookResponse, error)
	deleteBook(id int) (int64, error)
}

type BookApi interface {
	createBook(c *gin.Context)
	getAllBooks(c *gin.Context)
	getBook(c *gin.Context)
	deleteBook(c *gin.Context)
}
