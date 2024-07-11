package repositories

import (
	"context"

	"github.com/mukappalambda/my-trader/internal/core/domain/entities"
)

type MessageRepository interface {
	Create(ctx context.Context, msg *entities.Message) error
	Save(ctx context.Context, msg *entities.Message) error
	Delete(ctx context.Context, id string) error
}
