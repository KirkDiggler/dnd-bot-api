package rooms

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/KirkDiggler/dnd-bot-api/internal/entities"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client redis.UniversalClient
}

type RedisConfig struct {
	Client redis.UniversalClient
}

func NewRedis(cfg *RedisConfig) (*Redis, error) {
	if cfg == nil {
		return nil, errors.New("cfg is required")
	}

	if cfg.Client == nil {
		return nil, errors.New("cfg.Client is required")
	}

	return &Redis{
		client: cfg.Client,
	}, nil
}

func (r *Redis) GetRoom(ctx context.Context, id string) (*entities.Room, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	roomKey := getRoomKey(id)

	roomJson, err := r.client.Get(ctx, roomKey).Result()
	if err != nil {
		return nil, err
	}

	out := &entities.Room{}
	err = json.Unmarshal([]byte(roomJson), out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (r *Redis) CreateRoom(ctx context.Context, room *entities.Room) error {
	if room == nil {
		return errors.New("room is required")
	}

	if room.ID == "" {
		return errors.New("room.ID is required")
	}
	
	roomKey := getRoomKey(room.ID)

	roomJson, err := json.Marshal(room)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, roomKey, roomJson, 0).Err()
}

func getRoomKey(id string) string {
	return "room:" + id
}
