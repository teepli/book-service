package books

import (
	"net/http"

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
	body := BookQuery{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := a.service.createBook(body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func (a api) getBook(c *gin.Context) {
	id := c.Param("id")
	b, err := a.service.getBook(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, b)
}
func (a api) getAllBooks(c *gin.Context) {
	books, err := a.service.getAllBooks()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(200, books)
}
func (a api) deleteBook(c *gin.Context) {
	id := c.Param("id")
	err := a.service.deleteBook(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Status(http.StatusNoContent)
}
