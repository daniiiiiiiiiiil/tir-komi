package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *MediaService) UpdateMedia(ctx context.Context, id int, patch domain.MediaPatch) (domain.Media, error) {
	m, err := s.repo.GetMedia(ctx, id)
	if err != nil {
		return domain.Media{}, fmt.Errorf("mediaRepository.GetMedia: %w", err)
	}

	if err := m.ApplyPatch(patch); err != nil {
		return domain.Media{}, fmt.Errorf("apply patch: %w", err)
	}

	updated, err := s.repo.UpdateMedia(ctx, id, patch)
	if err != nil {
		return domain.Media{}, fmt.Errorf("mediaRepository.UpdateMedia: %w", err)
	}

	return updated, nil
}
