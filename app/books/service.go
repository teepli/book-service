package books

import (
	"fmt"
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
	return id, nil
}
func (s service) getBook(id string) (BookResponse, error) {
	book, err := s.repo.getBook(id)
	if err != nil {
		return book, err
	}
	return book, nil
}

func (s service) getAllBooks(params BookFilterParams) ([]BookResponse, error) {
	books, err := s.repo.getAllBooks(params)
	if err != nil {
		return books, err
	}
	return books, nil
}
func (s service) deleteBook(id string) error {
	count, err := s.repo.deleteBook(id)
	if err != nil {
		return err
	}
	if count < 1 {
		return fmt.Errorf(`Book with id %s not found`, id)
	}
	return nil
}
