-- name: CreateUser :one
INSERT INTO users (email, hashed_password)
VALUES ($1, $2)
RETURNING *;
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: GetUserForUpdate :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1 FOR NO KEY
UPDATE;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;
