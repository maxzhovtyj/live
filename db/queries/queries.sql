-- name: GetUser :one
SELECT *
FROM users
WHERE id = ?
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password_hash)
VALUES (?, ?, ?, ?)
RETURNING *;
