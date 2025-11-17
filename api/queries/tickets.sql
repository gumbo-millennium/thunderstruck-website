-- name: GetOneTicket :one
SELECT id, type, state, value, email, created_at, updated_at, deleted_at
FROM tickets
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetAllTickets :many
SELECT id, type, state, value, email, created_at, updated_at, deleted_at
FROM tickets
WHERE deleted_at IS NULL;

-- name: CreateTicket :one
INSERT INTO tickets (
  type, state, email, value
)
VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateTicket :one
UPDATE tickets SET
  type = $2,
  state = $3,
  email = $4,
  value = $5,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTicket :one
UPDATE tickets SET
  deleted_at = NOW()
WHERE id = $1
RETURNING *;
