package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *MediaService) GetMedia(ctx context.Context, id int) (domain.Media, error) {
	m, err := s.repo.GetMedia(ctx, id)
	if err != nil {
		return domain.Media{}, fmt.Errorf("get media: %w", err)
	}
	return m, nil
}

func (s *MediaService) GetMedias(ctx context.Context, mediaType *domain.MediaType, limit, offset int) ([]domain.Media, int, error) {
	items, total, err := s.repo.GetMedias(ctx, mediaType, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get medias: %w", err)
	}
	return items, total, nil
}
