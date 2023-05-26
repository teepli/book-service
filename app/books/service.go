package books

import (
	"errors"
)

type service struct {
	repo BookRepository
}

func NewService(repo BookRepository) service {
	return service{repo}
}

func (s service) createBook(bq BookQuery) (int, error) {
	id, err := s.repo.createBook(bq)
	if err != nil {
		return id, err
	}
	return id, err
}
func (s service) getBook(id string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
}
func (s service) getAllBooks() ([]BookResponse, error) {
	books, err := s.repo.getAllBooks()
	if err != nil {
		return books, err
	}
	return books, nil
}
func (s service) deleteBook(id string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
}
