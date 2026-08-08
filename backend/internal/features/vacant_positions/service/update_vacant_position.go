package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *VacantPositionService) UpdateVacantPosition(ctx context.Context, id int, patch domain.VacantPositionPatch) (domain.VacantPosition, error) {
	position, err := s.repo.GetVacantPosition(ctx, id)
	if err != nil {
		return domain.VacantPosition{}, fmt.Errorf("vacantPositionRepository.GetVacantPosition: %w", err)
	}

	if err := position.ApplyPatch(patch); err != nil {
		return domain.VacantPosition{}, fmt.Errorf("apply patch: %w", err)
	}

	updatedPosition, err := s.repo.UpdateVacantPosition(ctx, id, patch)
	if err != nil {
		return domain.VacantPosition{}, fmt.Errorf("vacantPositionRepository.UpdateVacantPosition: %w", err)
	}

	return updatedPosition, nil
}
