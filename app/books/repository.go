package books

import (
	"database/sql"
	"errors"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) repository {
	return repository{db}
}

func (r repository) createBook(bq BookQuery) (int, error) {
	bookId := 0
	err := r.db.QueryRow(`
		INSERT INTO books(title, author, year, publisher, description)
		VALUES($1, $2, $3, $4, $5)
		returning id`,
		bq.Title, bq.Author, bq.Year, bq.Publisher, bq.Description).Scan(&bookId)
	return bookId, err
}

func (r repository) getBook(id string) (BookResponse, error) {
	b := BookResponse{}
	err := r.db.QueryRow(`
		SELECT id, title, author, year, publisher, description
		FROM books
		WHERE id = $1`,
		id).Scan(&b.Id, &b.Title, &b.Author, &b.Year, &b.Publisher, &b.Description)
	return b, err
}

func (r repository) getAllBooks() ([]BookResponse, error) {
	rows, err := r.db.Query(
		`SELECT id, title, author, year, publisher, description FROM books`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []BookResponse{}

	for rows.Next() {
		var b BookResponse
		if err := rows.Scan(&b.Id, &b.Title, &b.Author, &b.Year, &b.Publisher, &b.Description); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}
func (r repository) deleteBook(id string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
}
