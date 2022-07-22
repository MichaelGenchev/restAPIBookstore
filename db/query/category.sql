
-- name: CreateCategory :one
INSERT INTO categories (
  name
) VALUES (
  $1
)
RETURNING *;


-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY id;


-- name: UpdateCategory :one
UPDATE categories
set name = $2
WHERE id = $1
RETURNING *;


-- name: DeleteCategory :one
DELETE FROM categories
WHERE id = $1
RETURNING *;