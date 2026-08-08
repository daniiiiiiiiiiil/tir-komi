package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (s *ReviewService) GetReviewsByRating(ctx context.Context, rating int, limit, offset int) ([]domain.Review, int, error) {
	if rating < 1 || rating > 5 {
		return nil, 0, fmt.Errorf("rating must be between 1 and 5")
	}
	limit, offset = pagination.LimitOffset(limit, offset)

	reviews, total, err := s.repo.GetReviewsByRating(ctx, rating, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get reviews by rating: %w", err)
	}
	return reviews, total, nil
}
