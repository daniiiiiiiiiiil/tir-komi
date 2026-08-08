package service

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type VacantPositionRepository interface {
	CreateVacantPosition(ctx context.Context, position domain.VacantPosition) (domain.VacantPosition, error)
	GetVacantPosition(ctx context.Context, id int) (domain.VacantPosition, error)
	GetVacantPositions(ctx context.Context, limit, offset int) ([]domain.VacantPosition, int, error)
	UpdateVacantPosition(ctx context.Context, id int, patch domain.VacantPositionPatch) (domain.VacantPosition, error)
	DeleteVacantPosition(ctx context.Context, id int) error
}

type VacantPositionService struct {
	repo VacantPositionRepository
}

func NewVacantPositionService(repo VacantPositionRepository) *VacantPositionService {
	return &VacantPositionService{
		repo: repo,
	}
}
