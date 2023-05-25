package books

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

const BOOK_BASE_ROUTE = "/books"

const EXAMPLE_BOOK = `{"title": "Harry Potter and the Philosophers Stone","author": "J.K. Rowling","year": 1997,"publisher": "Bloomsbury (UK)","description": "A book about a wizard boy"}`

func TestApi(t *testing.T) {
	r := gin.Default()
	NewBookRoutes(r)
	routes := []struct {
		method string
		path   string
		body   []byte
		code   int
	}{
		{"POST", BOOK_BASE_ROUTE, []byte(EXAMPLE_BOOK), 200},
		{"GET", BOOK_BASE_ROUTE, nil, 200},
		{"GET", BOOK_BASE_ROUTE + "/1", nil, 200},
		{"DEL", BOOK_BASE_ROUTE + "/1", nil, 204},
	}

	for _, route := range routes {
		req, err := http.NewRequest(route.method, route.path, bytes.NewBuffer(route.body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
		if status := rr.Code; status != route.code {
			t.Errorf(`wrong status code in %v-request - path %v: got %v want %v, body %v`,
				route.method, route.path, status, route.code, rr.Body.String())
		}
	}
}
