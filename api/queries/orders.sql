-- name: GetOneOrder :one
SELECT *
FROM orders
WHERE id = $1;

-- name: GetOneOrderByReference :one
SELECT *
FROM orders
WHERE reference = $1;

-- name: GetAllOrders :many
SELECT *
FROM orders;

-- name: CreateOrder :one
INSERT INTO orders (
  ticket_id, reference, checkout, state, email
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateOrder :one
UPDATE orders SET
  ticket_id = $2,
  reference = $3,
  checkout = $4,
  state = $5,
  email = $6,
  updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :one
UPDATE orders SET
  deleted_at = NOW()
WHERE id = $1
RETURNING *;
