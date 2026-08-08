package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *ReviewService) CreateReview(ctx context.Context, review domain.Review) (domain.Review, error) {
	if err := review.Validate(); err != nil {
		return domain.Review{}, fmt.Errorf("validation failed: %w", err)
	}

	created, err := s.repo.CreateReview(ctx, review)
	if err != nil {
		return domain.Review{}, fmt.Errorf("create review: %w", err)
	}

	return created, nil
}
