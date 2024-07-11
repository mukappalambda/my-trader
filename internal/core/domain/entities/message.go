package entities

import "time"

type Message struct {
	MessageID int32
	Topic     string
	Message   string
	CreatedAt time.Time
}
