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