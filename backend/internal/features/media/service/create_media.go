package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *MediaService) CreateMedia(ctx context.Context, m domain.Media) (domain.Media, error) {
	if err := m.Validate(); err != nil {
		return domain.Media{}, fmt.Errorf("validation failed: %w", err)
	}

	created, err := s.repo.CreateMedia(ctx, m)
	if err != nil {
		return domain.Media{}, fmt.Errorf("create media: %w", err)
	}

	return created, nil
}
