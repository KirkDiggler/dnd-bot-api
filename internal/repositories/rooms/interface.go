package rooms

import (
	"context"
	"github.com/KirkDiggler/dnd-bot-api/internal/entities"
)

type Repository interface {
	GetRoom(ctx context.Context, id string) (*entities.Room, error)
	CreateRoom(ctx context.Context, room *entities.Room) error
}
