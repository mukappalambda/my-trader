package ports

import (
	"context"

	"github.com/mukappalambda/my-trader/internal/core/domain/entities"
)

type MessageService interface {
	PublishMessage(ctx context.Context, msg *entities.Message) (v interface{}, err error)
}
