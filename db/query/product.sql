-- name: CreateProduct :one
INSERT INTO products (name, cost, quantity)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetProduct :one
SELECT *
FROM products
WHERE id = $1
LIMIT 1;
-- name: GetProductForUpdate :one
SELECT *
FROM products
WHERE id = $1
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListProducts :many
SELECT *
FROM products
ORDER BY id
LIMIT $1 OFFSET $2;
