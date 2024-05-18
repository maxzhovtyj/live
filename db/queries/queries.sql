-- name: CreateUser :one
INSERT INTO users (first_name, last_name, email, password_hash)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetAuthorizedUser :one
SELECT *
FROM users
WHERE email = $1
  AND password_hash = $2
LIMIT 1;

-- name: GetConversation :one
SELECT *
FROM conversations
WHERE id = $1
LIMIT 1;

-- name: GetUserConversations :many
SELECT *
FROM conversation_participants
WHERE user_id = $1;

-- name: GetConversationMessages :many
SELECT m.id,
       m.conversation_id,
       m.sender_id,
       m.body,
       m.created_at,
       u.id,
       concat(u.first_name, ' ', u.last_name)
FROM messages m
         LEFT JOIN users u ON m.sender_id = u.id
WHERE conversation_id = $1;

-- name: InsertMessageIntoConversation :exec
INSERT INTO messages (conversation_id, sender_id, body)
VALUES ($1, $2, $3);

