package v1alpha1

import (
	"context"
	"errors"
	"fmt"

	"github.com/KirkDiggler/dnd-bot-api/internal/clients/dnd5e"
	"github.com/KirkDiggler/dnd-bot-api/internal/entities"
	v1alpha1 "github.com/KirkDiggler/dnd-bot-api/protos/gen/go/api/v1alpha1"
)

type Handler struct {
	client dnd5e.Client
}

type HandlerConfig struct {
	Client dnd5e.Client
}

func NewHandler(cfg *HandlerConfig) (*Handler, error) {
	if cfg == nil {
		return nil, errors.New("cfg is required")
	}

	if cfg.Client == nil {
		return nil, errors.New("v1alpha1.HandlerConfig.Client is required")
	}

	return &Handler{
		client: cfg.Client,
	}, nil
}

func (h *Handler) ListRaces(ctx context.Context, req *v1alpha1.ListRacesRequest) (*v1alpha1.ListRacesResponse, error) {
	result, err := h.client.ListRaces()
	if err != nil {
		return nil, fmt.Errorf("error in client.ListRaces: %w", err)
	}

	return &v1alpha1.ListRacesResponse{
		Races: racesToProtos(result),
	}, nil
}

func raceToProto(r *entities.Race) *v1alpha1.Race {
	return &v1alpha1.Race{
		Id:          r.Key,
		Description: r.Name,
	}
}

func racesToProtos(rs []*entities.Race) []*v1alpha1.Race {
	out := make([]*v1alpha1.Race, len(rs))
	for i, r := range rs {
		out[i] = raceToProto(r)
	}

	return out
}

func (h *Handler) ListClasses(ctx context.Context, req *v1alpha1.ListClassesRequest) (*v1alpha1.ListClassesResponse, error) {
	return nil, errors.New("not implemented")
}
