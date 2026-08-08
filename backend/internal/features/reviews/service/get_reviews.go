package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (s *ReviewService) GetReviews(ctx context.Context, limit, offset int) ([]domain.Review, int, error) {
	limit, offset = pagination.LimitOffset(limit, offset)

	reviews, total, err := s.repo.GetReviews(ctx, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get reviews: %w", err)
	}
	return reviews, total, nil
}
