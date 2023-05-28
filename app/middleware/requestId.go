package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const xRequestIdHeader = "x-request-id"

func RequestIdInjector() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New()
		c.Request.Header.Add(xRequestIdHeader, id.String())
		c.Header(xRequestIdHeader, id.String())
		c.Next()
	}
}
