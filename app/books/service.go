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

func (s service) createBook(b string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
}
func (s service) getBook(id string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
}
func (s service) getAllBooks() ([]string, error) {
	err := errors.New("not implemented")
	strArr := []string{"bookRepo"}
	return strArr, err
}
func (s service) deleteBook(id string) (string, error) {
	err := errors.New("not implemented")
	return "bookRepo", err
}
