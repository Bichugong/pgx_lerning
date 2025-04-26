package repository

import (
	"context"
	"second_simple_server/pkg/models"
)

func (r *PGRepo) GetBooks() ([]models.Book, error) {
	rows, err := r.pool.Query(context.Background(), `
	SELECT id, name, author_id, genre_id, price
	FROM books;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Book
	for rows.Next() {
		var item models.Book
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.AuthorID,
			&item.GenreID,
			&item.Price,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}

	return data, nil
}

func (r *PGRepo) NewBook(item models.Book) error {

	_, err := r.pool.Exec(context.Background(),
	`INSERT INTO books (name, author_id, genre_id, price)
	VALUES ($1, $2, $3, $4);`,
	item.Name,
	item.AuthorID,
	item.GenreID,
	item.Price,
	)
	return err
}

func (r *PGRepo) GetBooksByID(id int) (models.Book, error) {
	var book models.Book
	err := r.pool.QueryRow(context.Background(), `
	SELECT id, name, author_id, genre_id, price
	FROM books
	WHERE id = $1;
	`, id,).Scan(
		&book.ID,
		&book.Name,
		&book.AuthorID,
		&book.GenreID,
		&book.Price,
	)
	return book, err
}
