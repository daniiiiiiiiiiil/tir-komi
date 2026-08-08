package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *VacantPositionService) GetVacantPosition(ctx context.Context, id int) (domain.VacantPosition, error) {
	position, err := s.repo.GetVacantPosition(ctx, id)
	if err != nil {
		return domain.VacantPosition{}, fmt.Errorf("get vacant position: %w", err)
	}
	return position, nil
}
