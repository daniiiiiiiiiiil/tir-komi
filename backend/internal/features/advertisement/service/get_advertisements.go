package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (s *AdvertisementService) GetAdvertisements(ctx context.Context, limit, offset int) ([]domain.Advertisement, int, error) {
	limit, offset = pagination.LimitOffset(limit, offset)

	ads, total, err := s.repo.GetAdvertisements(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get advertisements: %w", err)
	}
	return ads, total, nil
}
