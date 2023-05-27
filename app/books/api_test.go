package books

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/teepli/book-service/app/database"
	"github.com/teepli/book-service/app/tools"
)

const BOOK_BASE_ROUTE = "/books"

const EXAMPLE_BOOK = `{"title": "Harry Potter and the Philosophers Stone","author": "J.K. Rowling","year": 1997,"publisher": "Bloomsbury (UK)","description": "A book about a wizard boy"}`
const EXAMPLE_BOOK_RESPONSE = `{"id": 1, "title": "Harry Potter and the Philosophers Stone","author": "J.K. Rowling","year": 1997,"publisher": "Bloomsbury (UK)","description": "A book about a wizard boy"}`

func TestApi(t *testing.T) {
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	database.PrepareTables(db)
	g := gin.Default()
	tools.RegisterCustomValidators()

	r := NewRepository(db)
	s := NewService(r)
	NewApi(g, s)

	routes := []struct {
		method   string
		path     string
		body     []byte
		code     int
		expected string
		model    interface{}
	}{
		{"POST", BOOK_BASE_ROUTE, []byte(EXAMPLE_BOOK), 200, (`{"id": 1}`), CreateBookResponse{}},
		{"GET", BOOK_BASE_ROUTE, nil, 200, "", nil},
		{"GET", BOOK_BASE_ROUTE + "/1", nil, 200, EXAMPLE_BOOK_RESPONSE, BookResponse{}},
		{"DELETE", BOOK_BASE_ROUTE + "/1", nil, 204, "", nil},
	}

	for _, route := range routes {
		req, err := http.NewRequest(route.method, route.path, bytes.NewBuffer(route.body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		g.ServeHTTP(rr, req)
		if status := rr.Code; status != route.code {
			t.Errorf(`wrong status code in %v-request - path %v: got %v want %v, body %v`,
				route.method, route.path, status, route.code, rr.Body.String())
		}

		if route.model != nil {
			response := route.model
			if err := json.Unmarshal([]byte(rr.Body.String()), &response); err != nil {
				t.Fatal(err)
			}

			expected := route.model
			if err := json.Unmarshal([]byte(route.expected), &expected); err != nil {
				t.Fatal(err)
			}

			a := assert.New(t)
			a.Equal(response, expected)
		}
	}
}
