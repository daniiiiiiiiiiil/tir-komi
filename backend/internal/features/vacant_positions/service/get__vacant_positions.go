package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (s *VacantPositionService) GetVacantPositions(ctx context.Context, limit, offset int) ([]domain.VacantPosition, int, error) {
	limit, offset = pagination.LimitOffset(limit, offset)

	positions, total, err := s.repo.GetVacantPositions(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get vacant positions: %w", err)
	}
	return positions, total, nil
}
