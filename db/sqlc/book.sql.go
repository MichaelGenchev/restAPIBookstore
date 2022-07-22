// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: book.sql

package db

import (
	"context"
	_ "github.com/lib/pq"

)

const createBook = `-- name: CreateBook :one
INSERT INTO books (
  title, author, category, price
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, title, author, category, price
`

type CreateBookParams struct {
	Title    string  `json:"title"`
	Author   int64   `json:"author"`
	Category int64   `json:"category"`
	Price    float64 `json:"price"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, createBook,
		arg.Title,
		arg.Author,
		arg.Category,
		arg.Price,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Category,
		&i.Price,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING id, title, author, category, price
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, deleteBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Category,
		&i.Price,
	)
	return i, err
}

const getBook = `-- name: GetBook :one
SELECT id, title, author, category, price FROM books
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRowContext(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Category,
		&i.Price,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT id, title, author, category, price FROM books
ORDER BY id
`

func (q *Queries) ListBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.QueryContext(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Book
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Author,
			&i.Category,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBook = `-- name: UpdateBook :one
UPDATE books
set price = $2
WHERE id = $1
RETURNING id, title, author, category, price
`

type UpdateBookParams struct {
	ID    int64           `json:"id"`
	Price float64 `json:"price"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) (Book, error) {
	row := q.db.QueryRowContext(ctx, updateBook, arg.ID, arg.Price)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Author,
		&i.Category,
		&i.Price,
	)
	return i, err
}
