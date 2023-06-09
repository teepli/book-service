package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teepli/book-service/app/common"
)

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
		c.AbortWithError(http.StatusBadRequest, common.NewValidationError(err.Error(), ""))
		return
	}

	id, err := a.service.createBook(body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := CreateBookResponse{Id: id}
	c.JSON(200, response)
}

func (a api) getBook(c *gin.Context) {
	param := IdParam{}
	if err := c.ShouldBindUri(&param); err != nil {
		c.AbortWithError(http.StatusNotFound, common.NewValidationError(err.Error(), ""))
		return
	}

	b, err := a.service.getBook(param.Id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.JSON(200, b)
}

func (a api) getAllBooks(c *gin.Context) {
	queryParams := BookFilterParams{}
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.AbortWithError(http.StatusBadRequest, common.NewValidationError(err.Error(), ""))
		return
	}

	books, err := a.service.getAllBooks(queryParams)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(200, books)
}

func (a api) deleteBook(c *gin.Context) {
	param := IdParam{}
	if err := c.ShouldBindUri(&param); err != nil {
		c.AbortWithError(http.StatusNotFound, common.NewValidationError(err.Error(), ""))
		return
	}

	err := a.service.deleteBook(param.Id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusNoContent)
}
