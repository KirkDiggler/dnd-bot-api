package dnd5e

import "github.com/KirkDiggler/dnd-bot-api/internal/entities"

type Client interface {
	ListClasses() ([]*entities.Class, error)
	ListRaces() ([]*entities.Race, error)
	GetClass(key string) (*entities.Class, error)
}
