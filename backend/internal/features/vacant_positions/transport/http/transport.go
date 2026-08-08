package http

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type VacantHandler struct {
	vacantService VacantService
}

type VacantService interface {
	CreateVacantPosition(ctx context.Context, position domain.VacantPosition) (domain.VacantPosition, error)
	DeleteVacantPosition(ctx context.Context, id int) error
	GetVacantPositions(ctx context.Context, limit, offset int) ([]domain.VacantPosition, int, error)
	GetVacantPosition(ctx context.Context, id int) (domain.VacantPosition, error)
	UpdateVacantPosition(ctx context.Context, id int, patch domain.VacantPositionPatch) (domain.VacantPosition, error)
}

func NewVacantHandler(vacantService VacantService) *VacantHandler {
	return &VacantHandler{
		vacantService: vacantService,
	}
}

func (handler *VacantHandler) Routers() []server.Route {
	return []server.Route{
		{
			Method:  "POST",
			Path:    "/vacancies",
			Handler: handler.CreateVacant,
		},
		{
			Method:  "GET",
			Path:    "/vacancies",
			Handler: handler.GetVacancies,
		},
		{
			Method:  "GET",
			Path:    "/vacancies/{id}",
			Handler: handler.GetVacant,
		},
		{
			Method:  "DELETE",
			Path:    "/vacancies/{id}",
			Handler: handler.DeleteVacant,
		},
		{
			Method:  "PATCH",
			Path:    "/vacancies/{id}",
			Handler: handler.UpdateVacant,
		},
	}
}
