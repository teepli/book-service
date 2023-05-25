package books

import (
	"github.com/gin-gonic/gin"
)

func failFunction(c *gin.Context) {
	c.AbortWithStatus(404)
}

func NewBookRoutes(r *gin.Engine) *gin.RouterGroup {

	rg := r.Group("/books")
	rg.GET("", failFunction)
	rg.POST("", failFunction)
	rg.GET("/:id", failFunction)
	rg.DELETE("/:id", failFunction)
	return rg
}
