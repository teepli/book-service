package books

import "github.com/gin-gonic/gin"

type BookService interface {
	createBook(b string) (string, error)
	getAllBooks() ([]string, error)
	getBook(id string) (string, error)
	deleteBook(id string) (string, error)
}

type BookRepository interface {
	createBook(b string) (string, error)
	getAllBooks() ([]string, error)
	getBook(id string) (string, error)
	deleteBook(id string) (string, error)
}

type BookApi interface {
	createBook(c *gin.Context)
	getAllBooks(c *gin.Context)
	getBook(c *gin.Context)
	deleteBook(c *gin.Context)
}
