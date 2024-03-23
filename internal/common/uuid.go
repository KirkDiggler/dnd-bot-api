package common

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() string
}

type GoogleUUID struct{}

func (g *GoogleUUID) New() string {
	return uuid.New().String()
}
