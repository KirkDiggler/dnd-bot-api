package v1alpha1

import (
	"context"
	"errors"
	"github.com/KirkDiggler/dnd-bot-api/internal/repositories/rooms"
	dndbotv1alpha1api "github.com/KirkDiggler/dnd-bot-api/protos/gen-external/go/api/admin"
	"log/slog"
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
	return &dndbotv1alpha1api.ListRoomsResponse{
		Rooms: []*dndbotv1alpha1api.Room{
			{
				Id:   "1",
				Name: "room1",
			},
		},
	}, nil
}

func (h *Handler) CreateRoom(ctx context.Context, req *dndbotv1alpha1api.CreateRoomRequest) (*dndbotv1alpha1api.CreateRoomResponse, error) {
	slog.Info("CreateRoom")
	return nil, errors.New("not implemented")
}

func (h *Handler) UpdateRoom(ctx context.Context, req *dndbotv1alpha1api.UpdateRoomRequest) (*dndbotv1alpha1api.UpdateRoomResponse, error) {
	slog.Info("UpdateRoom")
	return nil, errors.New("not implemented")
}

func (h *Handler) DeleteRoom(ctx context.Context, req *dndbotv1alpha1api.DeleteRoomRequest) (*dndbotv1alpha1api.DeleteRoomResponse, error) {
	slog.Info("DeleteRoom")
	return nil, errors.New("not implemented")
}
