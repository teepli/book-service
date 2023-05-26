package books

import (
	"github.com/gin-gonic/gin"
)

func failFunction(c *gin.Context) {
	c.AbortWithStatus(404)
}

type api struct {
	service BookService
}

func NewApi(g *gin.Engine, s BookService) BookApi {
	a := api{s}
	rg := g.Group("/books")
	rg.GET("", a.getAllBooks)
	rg.POST("", a.createBook)
	rg.GET("/:id", a.getBook)
	rg.DELETE("/:id", a.deleteBook)

	return a
}

func (a api) createBook(c *gin.Context) {
	b, err := a.service.createBook("1")
	if err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, b)
}
func (a api) getBook(c *gin.Context) {
	b, err := a.service.getBook("1")
	if err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, b)
}
func (a api) getAllBooks(c *gin.Context) {
	b, err := a.service.getAllBooks()
	if err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, b)
}
func (a api) deleteBook(c *gin.Context) {
	b, err := a.service.deleteBook("1")
	if err != nil {
		c.Status(404)
		return
	}
	c.JSON(200, b)
}
