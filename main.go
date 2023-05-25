package main

import (
	"github.com/gin-gonic/gin"
	"github.com/teepli/book-service/app/books"
)

func main() {
	r := gin.Default()
	books.NewBookRoutes(r)
	r.Run(":9000")
}
