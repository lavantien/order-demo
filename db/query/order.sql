-- name: CreateOrder :one
INSERT INTO orders (user_id, product_id, quantity, price)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetOrder :one
SELECT *
FROM orders
WHERE id = $1
LIMIT 1;
-- name: GetOrderForUpdate :one
SELECT *
FROM orders
WHERE id = $1
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListOrders :many
SELECT *
FROM orders
ORDER BY id
LIMIT $1 OFFSET $2;
