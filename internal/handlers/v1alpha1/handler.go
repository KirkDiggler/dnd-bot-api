package v1alpha1

import (
	"context"
	"errors"
	"log/slog"

	"github.com/KirkDiggler/dnd-bot-api/internal/entities"
	"github.com/KirkDiggler/dnd-bot-api/internal/repositories/rooms"
	dndbotv1alpha1api "github.com/KirkDiggler/dnd-bot-api/protos/gen-external/go/api/admin"
)

type Handler struct {
	roomRepo rooms.Repository
}

type HandlerConfig struct {
	RoomRepo rooms.Repository
}

func NewHandler(cfg *HandlerConfig) (*Handler, error) {
	if cfg == nil {
		return nil, errors.New("cfg is required")
	}

	if cfg.RoomRepo == nil {
		return nil, errors.New("cfg.RoomRepo is required")
	}

	return &Handler{
		roomRepo: cfg.RoomRepo,
	}, nil
}

func (h *Handler) GetRoom(ctx context.Context, req *dndbotv1alpha1api.GetRoomRequest) (*dndbotv1alpha1api.GetRoomResponse, error) {
	if req == nil {
		return nil, errors.New("req is required")
	}

	if req.Id == "" {
		return nil, errors.New("req.Id is required")
	}

	room, err := h.roomRepo.GetRoom(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &dndbotv1alpha1api.GetRoomResponse{
		Room: &dndbotv1alpha1api.Room{
			Id:          room.ID,
			Name:        room.Name,
			Description: room.Description,
		},
	}, nil
}

func (h *Handler) ListRooms(ctx context.Context, req *dndbotv1alpha1api.ListRoomsRequest) (*dndbotv1alpha1api.ListRoomsResponse, error) {
	slog.Info("ListRooms")

	roomsData, err := h.roomRepo.ListRooms(ctx)
	if err != nil {
		return nil, err
	}

	roomProtos := make([]*dndbotv1alpha1api.Room, 0, len(roomsData))
	for _, room := range roomsData {
		roomProtos = append(roomProtos, &dndbotv1alpha1api.Room{
			Id:          room.ID,
			Name:        room.Name,
			Description: room.Description,
		})
	}

	return &dndbotv1alpha1api.ListRoomsResponse{
		Rooms: roomProtos,
	}, nil
}

func (h *Handler) CreateRoom(ctx context.Context, req *dndbotv1alpha1api.CreateRoomRequest) (*dndbotv1alpha1api.CreateRoomResponse, error) {
	slog.Info("CreateRoom")
	if req == nil {
		return nil, errors.New("req is required")
	}

	if req.Room == nil {
		return nil, errors.New("req.Room is required")
	}

	if req.Room.Id != "" {
		return nil, errors.New("req.Room.Id cannot be set on create")
	}

	if req.Room.Name == "" {
		return nil, errors.New("req.Room.Name is required")
	}

	err := h.roomRepo.CreateRoom(ctx, &entities.Room{
		ID:          req.Room.Id,
		Name:        req.Room.Name,
		Description: req.Room.Description,
	})
	if err != nil {
		return nil, err
	}

	return &dndbotv1alpha1api.CreateRoomResponse{}, nil
}

func (h *Handler) UpdateRoom(ctx context.Context, req *dndbotv1alpha1api.UpdateRoomRequest) (*dndbotv1alpha1api.UpdateRoomResponse, error) {
	slog.Info("UpdateRoom")
	if req == nil {
		return nil, errors.New("req is required")
	}

	if req.Room == nil {
		return nil, errors.New("req.Room is required")
	}

	if req.Room.Id == "" {
		return nil, errors.New("req.Room.Id is required")
	}

	result, err := h.roomRepo.UpdateRoom(ctx, &rooms.UpdateInput{
		Room: &entities.Room{
			ID:          req.Room.Id,
			Name:        req.Room.Name,
			Description: req.Room.Description,
		},
		InputMask: []string{"name"},
	})
	if err != nil {
		return nil, err
	}

	return &dndbotv1alpha1api.UpdateRoomResponse{
		Room: &dndbotv1alpha1api.Room{
			Id:          result.Room.ID,
			Name:        result.Room.Name,
			Description: result.Room.Description,
		},
	}, nil
}

func (h *Handler) DeleteRoom(ctx context.Context, req *dndbotv1alpha1api.DeleteRoomRequest) (*dndbotv1alpha1api.DeleteRoomResponse, error) {
	slog.Info("DeleteRoom")
	return nil, errors.New("not implemented")
}
