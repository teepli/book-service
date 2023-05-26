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

func (r repository) createBook(b string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
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
