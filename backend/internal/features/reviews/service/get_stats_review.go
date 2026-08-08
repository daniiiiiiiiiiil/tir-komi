package service

import (
	"context"
	"fmt"
)

func (s *ReviewService) GetReviewStats(ctx context.Context) (ReviewStats, error) {
	stats, err := s.repo.GetReviewStats(ctx)
	if err != nil {
		return ReviewStats{}, fmt.Errorf("get review stats: %w", err)
	}
	return stats, nil
}
