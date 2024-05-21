-- name: GetMessage :one
SELECT * FROM messages
WHERE message_id = $1 LIMIT 1;

-- name: CreateMessage :one
INSERT INTO messages (
  topic, message
) VALUES (
  $1, $2
)
RETURNING *;
