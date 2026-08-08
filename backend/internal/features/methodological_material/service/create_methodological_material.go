package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *MethodologicalMaterialService) CreateMethodologicalMaterial(ctx context.Context, material domain.MethodologicalMaterial) (domain.MethodologicalMaterial, error) {
	if err := material.Validate(); err != nil {
		return domain.MethodologicalMaterial{}, fmt.Errorf("validation failed: %w", err)
	}

	created, err := s.repo.CreateMethodologicalMaterial(ctx, material)
	if err != nil {
		return domain.MethodologicalMaterial{}, fmt.Errorf("create methodological material: %w", err)
	}

	return created, nil
}
