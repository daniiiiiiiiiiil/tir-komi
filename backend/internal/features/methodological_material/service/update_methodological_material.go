package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *MethodologicalMaterialService) UpdateMethodologicalMaterial(ctx context.Context, id int, patch domain.MethodologicalMaterialPatch) (domain.MethodologicalMaterial, error) {
	material, err := s.repo.GetMethodologicalMaterial(ctx, id)
	if err != nil {
		return domain.MethodologicalMaterial{}, fmt.Errorf("methodologicalMaterialRepository.GetMethodologicalMaterial: %w", err)
	}

	if err := material.ApplyPatch(patch); err != nil {
		return domain.MethodologicalMaterial{}, fmt.Errorf("apply patch: %w", err)
	}

	updatedMaterial, err := s.repo.UpdateMethodologicalMaterial(ctx, id, patch)
	if err != nil {
		return domain.MethodologicalMaterial{}, fmt.Errorf("methodologicalMaterialRepository.UpdateMethodologicalMaterial: %w", err)
	}

	return updatedMaterial, nil
}
