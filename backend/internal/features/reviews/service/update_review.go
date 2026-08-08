package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *ReviewService) UpdateReview(ctx context.Context, id int, patch domain.ReviewPatch) (domain.Review, error) {
	review, err := s.repo.GetReview(ctx, id)
	if err != nil {
		return domain.Review{}, fmt.Errorf("reviewRepository.GetReview: %w", err)
	}

	if err := review.ApplyPatch(patch); err != nil {
		return domain.Review{}, fmt.Errorf("apply patch: %w", err)
	}

	updatedReview, err := s.repo.UpdateReview(ctx, id, patch)
	if err != nil {
		return domain.Review{}, fmt.Errorf("reviewRepository.UpdateReview: %w", err)
	}

	return updatedReview, nil
}
