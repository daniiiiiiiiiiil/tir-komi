package service

import (
	"context"
	"fmt"
)

func (s *MethodologicalMaterialService) DeleteMethodologicalMaterial(ctx context.Context, id int) error {
	_, err := s.repo.GetMethodologicalMaterial(ctx, id)
	if err != nil {
		return fmt.Errorf("get methodological material: %w", err)
	}
	if err := s.repo.DeleteMethodologicalMaterial(ctx, id); err != nil {
		return fmt.Errorf("delete methodological material: %w", err)
	}
	return nil
}
