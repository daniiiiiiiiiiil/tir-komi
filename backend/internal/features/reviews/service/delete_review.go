package service

import (
	"context"
	"fmt"
)

func (s *ReviewService) DeleteReview(ctx context.Context, id int) error {
	_, err := s.repo.GetReview(ctx, id)
	if err != nil {
		return fmt.Errorf("get review: %w", err)
	}
	if err := s.repo.DeleteReview(ctx, id); err != nil {
		return fmt.Errorf("delete review: %w", err)
	}
	return nil
}
