// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package messages

import (
	"context"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO messages (
  topic, message
) VALUES (
  $1, $2
)
RETURNING message_id, topic, message, created_at
`

type CreateMessageParams struct {
	Topic   string
	Message string
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, createMessage, arg.Topic, arg.Message)
	var i Message
	err := row.Scan(
		&i.MessageID,
		&i.Topic,
		&i.Message,
		&i.CreatedAt,
	)
	return i, err
}

const getMessage = `-- name: GetMessage :one
SELECT message_id, topic, message, created_at FROM messages
WHERE message_id = $1 LIMIT 1
`

func (q *Queries) GetMessage(ctx context.Context, messageID int32) (Message, error) {
	row := q.db.QueryRow(ctx, getMessage, messageID)
	var i Message
	err := row.Scan(
		&i.MessageID,
		&i.Topic,
		&i.Message,
		&i.CreatedAt,
	)
	return i, err
}
