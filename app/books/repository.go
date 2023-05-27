package books

import (
	"database/sql"
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
		RETURNING id
		`, bq.Title, bq.Author, bq.Year, bq.Publisher, bq.Description).Scan(&bookId)
	return bookId, err
}

func (r repository) getBook(id int) (BookResponse, error) {
	b := BookResponse{}
	err := r.db.QueryRow(`
		SELECT id, title, author, year, publisher, description
		FROM books
		WHERE id = $1
		`, id).Scan(&b.Id, &b.Title, &b.Author, &b.Year, &b.Publisher, &b.Description)
	return b, err
}

func (r repository) getAllBooks(params BookFilterParams) ([]BookResponse, error) {
	rows, err := r.db.Query(`
		SELECT id, title, author, year, publisher, description
		FROM books
		WHERE
		(year LIKE '%' || $1 || '%' OR $1 IS NULL) AND
		(author LIKE '%' || $2 || '%' OR $2 IS NULL) AND
		(publisher LIKE '%' || $3 || '%' OR $3 IS NULL)
		`, params.Year, params.Author, params.Publisher)
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

func (r repository) deleteBook(id int) (int64, error) {
	count, err := r.db.Exec(`
		DELETE FROM books WHERE id = $1
		`, id)
	if err != nil {
		return -1, err
	}

	return count.RowsAffected()
}
