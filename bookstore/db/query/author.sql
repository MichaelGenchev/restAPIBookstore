
-- name: CreateAuthor :one
INSERT INTO authors (
  name, 
  biography
) VALUES (
  $1, $2
)
RETURNING *;


-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY id;


-- name: UpdateAuthor :one
UPDATE authors
set name = $2,
biography = $3
WHERE id = $1
RETURNING *;


-- name: DeleteAuthor :one
DELETE FROM authors
WHERE id = $1
RETURNING *;