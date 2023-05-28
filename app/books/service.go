package books

import (
	"database/sql"
	"fmt"

	"github.com/teepli/book-service/app/common"
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
		return id, common.NewDatabaseError(err.Error(), "")
	}

	return id, nil
}

func (s service) getBook(id int) (BookResponse, error) {
	book, err := s.repo.getBook(id)
	if err != nil {
		if err == sql.ErrNoRows {
			msg := fmt.Sprintf("Book with id %d not found", id)
			return book, common.NewDatabaseError(msg, "")
		}
		return book, common.NewDatabaseError(err.Error(), "")
	}

	return book, nil
}

func (s service) getAllBooks(params BookFilterParams) ([]BookResponse, error) {
	books, err := s.repo.getAllBooks(params)
	if err != nil {
		return books, common.NewDatabaseError(err.Error(), "")
	}

	return books, nil
}

func (s service) deleteBook(id int) error {
	count, err := s.repo.deleteBook(id)
	if err != nil {
		return common.NewDatabaseError(err.Error(), "")
	}
	if count < 1 {
		msg := fmt.Sprintf("Book with id %d not found", id)
		return common.NewDatabaseError(msg, "")
	}

	return nil
}
