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

func TestApi(t *testing.T) {
	db, err := database.InitDatabase("file://../../migrations")
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
		method           string
		path             string
		body             []byte
		code             int
		expectedResponse string
		responseModel    interface{}
	}{
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_ONE), 200, (BOOK_ONE_RESPONSE), CreateBookResponse{}},
		{"GET", BOOK_BASE_ROUTE, nil, 200, "", nil},

		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_TWO), 200, (BOOK_TWO_RESPONSE), CreateBookResponse{}},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_THREE), 200, (BOOK_THREE_RESPONSE), CreateBookResponse{}},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_FOUR), 200, (BOOK_FOUR_RESPONSE), CreateBookResponse{}},

		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_MISSING_TITLE), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_MISSING_YEAR), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_MISSING_YEAR_TWO), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_EMPTY_AUTHOR), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_EMPTY_TITLE), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_NON_INTEGER_YEAR), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_NON_INTEGER_YEAR_TWO), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_EMPTY_PUBLISHER), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(BOOK_WITH_NON_UNIQUE_VALUES), 400, "", nil},
		{"POST", BOOK_BASE_ROUTE, []byte(INVALID_JSON), 400, "", nil},

		{"GET", BOOK_BASE_ROUTE, nil, 200, (ALL_BOOKS_RESPONSE), []BookResponse{}},
		{"GET", BOOK_BASE_ROUTE + "?author=J%2EK%2E%20Rowling", nil, 200, (BOOKS_FILTERED_BY_AUTHOR_RESPONSE), []BookResponse{}},
		{"GET", BOOK_BASE_ROUTE + "?year=1997", nil, 200, (BOOKS_FILTERD_BY_YEAR_RESPONSE), []BookResponse{}},
		{"GET", BOOK_BASE_ROUTE + "?publisher=Otava", nil, 200, (BOOKS_FILTERED_BY_NONEXISTING_PUBLISHER_RESPONSE), []BookResponse{}},
		{"GET", BOOK_BASE_ROUTE + "?year=1997&publisher=scholastic%20point", nil, 200, (BOOKS_FILTERED_BY_YEAR_AND_PUBLISHER_RESPONSE), []BookResponse{}},

		{"GET", BOOK_BASE_ROUTE + "/1", nil, 200, BOOK_ONE_GET_RESPONSE, BookResponse{}},
		{"GET", BOOK_BASE_ROUTE + "/0", nil, 400, "", nil},
		{"GET", BOOK_BASE_ROUTE + "/x", nil, 400, "", nil},
		{"GET", BOOK_BASE_ROUTE + "/1.5", nil, 400, "", nil},

		{"DELETE", BOOK_BASE_ROUTE + "/1", nil, 204, "", nil},
		{"DELETE", BOOK_BASE_ROUTE + "/1", nil, 400, "", nil},
		{"DELETE", BOOK_BASE_ROUTE + "/0", nil, 400, "", nil},
		{"DELETE", BOOK_BASE_ROUTE + "/x", nil, 400, "", nil},
		{"DELETE", BOOK_BASE_ROUTE + "/1.5", nil, 400, "", nil},

		{"GET", BOOK_BASE_ROUTE, nil, 200, (BOOKS_FINAL_STATE), []BookResponse{}},
	}

	for _, route := range routes {
		req, err := http.NewRequest(route.method, route.path, bytes.NewBuffer(route.body))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		g.ServeHTTP(rr, req)
		if status := rr.Code; status != route.code {
			t.Errorf(`Wrong status code in %v-request - path %v: got %v want %v, body: %v`,
				route.method, route.path, status, route.code, rr.Body.String())
		}

		if route.responseModel != nil {
			response := route.responseModel
			if err := json.Unmarshal([]byte(rr.Body.String()), &response); err != nil {
				t.Fatal(err)
			}

			expected := route.responseModel
			if err := json.Unmarshal([]byte(route.expectedResponse), &expected); err != nil {
				t.Fatal(err)
			}

			a := assert.New(t)
			a.Equal(response, expected)
		}
	}
}
