-- name: GetOne :one
SELECT id, type, state, email, value, created_at, updated_at, deleted_at
FROM tickets
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetAll :many
SELECT id, type, state, email, value, created_at, updated_at, deleted_at
FROM tickets
WHERE deleted_at IS NULL;

-- name: Create :one
INSERT INTO tickets (
  type, state, email, value, updated_at, deleted_at
)
VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: Update :one
UPDATE tickets SET
  type = $2,
  state = $3,
  email = $4,
  value = $5,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: Delete :one
UPDATE tickets SET
  deleted_at = NOW()
WHERE id = $1
RETURNING *;
