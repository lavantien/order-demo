-- name: CreateOrder :one
INSERT INTO
    orders (owner, product_id, quantity, price)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetOrder :one
SELECT
    *
FROM
    orders
WHERE
    id = $1
LIMIT
    1;

-- name: ListOrders :many
SELECT
    *
FROM
    orders
ORDER BY
    id
LIMIT
    $1
OFFSET
    $2;

-- name: ListOrdersByOwner :many
SELECT
    *
FROM
    orders
WHERE
    owner = $3
ORDER BY
    id
LIMIT
    $1
OFFSET
    $2;
