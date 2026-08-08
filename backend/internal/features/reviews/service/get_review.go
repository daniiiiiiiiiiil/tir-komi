package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *ReviewService) GetReview(ctx context.Context, id int) (domain.Review, error) {
	review, err := s.repo.GetReview(ctx, id)
	if err != nil {
		return domain.Review{}, fmt.Errorf("get review: %w", err)
	}
	return review, nil
}
