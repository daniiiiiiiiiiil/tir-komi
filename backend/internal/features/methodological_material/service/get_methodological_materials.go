package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (s *MethodologicalMaterialService) GetMethodologicalMaterials(ctx context.Context, limit, offset int) ([]domain.MethodologicalMaterial, int, error) {
	limit, offset = pagination.LimitOffset(limit, offset)

	materials, total, err := s.repo.GetMethodologicalMaterials(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get methodological materials: %w", err)
	}
	return materials, total, nil
}
