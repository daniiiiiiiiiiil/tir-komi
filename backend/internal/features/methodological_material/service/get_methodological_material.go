package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *MethodologicalMaterialService) GetMethodologicalMaterial(ctx context.Context, id int) (domain.MethodologicalMaterial, error) {
	material, err := s.repo.GetMethodologicalMaterial(ctx, id)
	if err != nil {
		return domain.MethodologicalMaterial{}, fmt.Errorf("get methodological material: %w", err)
	}
	return material, nil
}
