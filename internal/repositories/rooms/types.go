package rooms

import "github.com/KirkDiggler/dnd-bot-api/internal/entities"

type UpdateInput struct {
	Room      *entities.Room
	InputMask []string
}

type UpdateOutput struct {
	Room *entities.Room
}
