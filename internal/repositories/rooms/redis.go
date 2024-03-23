package rooms

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/KirkDiggler/dnd-bot-api/internal/common"

	"github.com/KirkDiggler/dnd-bot-api/internal/entities"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client        redis.UniversalClient
	uuidGenerator common.UUIDGenerator
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
		client:        cfg.Client,
		uuidGenerator: &common.GoogleUUID{},
	}, nil
}

func (r *Redis) ListRooms(ctx context.Context) ([]*entities.Room, error) {
	roomKeys, err := r.client.SMembers(ctx, getRoomListKey()).Result()
	if err != nil {
		return nil, err
	}

	rooms := make([]*entities.Room, 0, len(roomKeys))

	for _, roomKey := range roomKeys {
		roomJson, err := r.client.Get(ctx, roomKey).Result()
		if err != nil {
			return nil, err
		}

		room := &entities.Room{}
		err = json.Unmarshal([]byte(roomJson), room)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
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

	if room.ID != "" {
		return errors.New("room.ID connot be set on create")
	}

	room.ID = r.uuidGenerator.New()

	roomKey := getRoomKey(room.ID)

	roomJson, err := json.Marshal(room)
	if err != nil {
		return err
	}

	r.client.Set(ctx, roomKey, roomJson, 0)
	r.client.SAdd(ctx, getRoomListKey(), getRoomKey(room.ID))

	return err
}

func getRoomListKey() string {
	return "room:list"
}

func getRoomKey(id string) string {
	return "room:" + id
}
