package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *VacantPositionService) CreateVacantPosition(ctx context.Context, position domain.VacantPosition) (domain.VacantPosition, error) {
	if err := position.Validate(); err != nil {
		return domain.VacantPosition{}, fmt.Errorf("validation failed: %w", err)
	}

	created, err := s.repo.CreateVacantPosition(ctx, position)
	if err != nil {
		return domain.VacantPosition{}, fmt.Errorf("create vacant position: %w", err)
	}

	return created, nil
}
