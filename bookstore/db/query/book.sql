
-- name: CreateBook :one
INSERT INTO books (
  title, author, category, price
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;


-- name: GetBook :one
SELECT * FROM books
WHERE id = $1 LIMIT 1;

-- name: ListBooks :many
SELECT * FROM books
ORDER BY id;


-- name: UpdateBook :one
UPDATE books
set price = $2
WHERE id = $1
RETURNING *;


-- name: DeleteBook :one
DELETE FROM books
WHERE id = $1
RETURNING *;