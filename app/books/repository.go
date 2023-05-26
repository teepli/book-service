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

func (r repository) getBook(id string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
}
func (r repository) getAllBooks() ([]string, error) {
	err := errors.New("not implemented")
	strArr := []string{"Abc"}
	return strArr, err
}
func (r repository) deleteBook(id string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
}
