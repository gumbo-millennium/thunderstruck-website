-- name: GetOne :one
SELECT id, type, email, created_at, updated_at, deleted_at
FROM tickets
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetAll :many
SELECT id, type, email, created_at, updated_at, deleted_at
FROM tickets
WHERE deleted_at IS NULL;

-- name: Create :one
INSERT INTO tickets (
  id, type, email, updated_at, deleted_at
)
VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: Update :one
UPDATE tickets SET
  type = $2,
  email = $3,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: Delete :one
UPDATE tickets SET
  deleted_at = NOW()
WHERE id = $1
RETURNING *;
