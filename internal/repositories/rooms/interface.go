package rooms

import (
	"context"

	"github.com/KirkDiggler/dnd-bot-api/internal/entities"
)

type Repository interface {
	GetRoom(ctx context.Context, id string) (*entities.Room, error)
	CreateRoom(ctx context.Context, room *entities.Room) error
	UpdateRoom(ctx context.Context, input *UpdateInput) (*UpdateOutput, error)
	ListRooms(ctx context.Context) ([]*entities.Room, error)
}
