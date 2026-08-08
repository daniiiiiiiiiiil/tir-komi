package postgres

import (
	"context"
	"fmt"
)

type ReviewStats struct {
	AvgRating  float64 `json:"avg_rating"`
	TotalVotes int     `json:"total_votes"`
	Rating1    int     `json:"rating_1"`
	Rating2    int     `json:"rating_2"`
	Rating3    int     `json:"rating_3"`
	Rating4    int     `json:"rating_4"`
	Rating5    int     `json:"rating_5"`
}

func (r *ReviewRepository) GetReviewStats(ctx context.Context) (ReviewStats, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var stats ReviewStats

	query := `
		SELECT 
			COALESCE(AVG(rating), 0) as avg_rating,
			COUNT(*) as total_votes,
			COUNT(*) FILTER (WHERE rating = 1) as rating_1,
			COUNT(*) FILTER (WHERE rating = 2) as rating_2,
			COUNT(*) FILTER (WHERE rating = 3) as rating_3,
			COUNT(*) FILTER (WHERE rating = 4) as rating_4,
			COUNT(*) FILTER (WHERE rating = 5) as rating_5
		FROM reviews
	`

	err := r.pool.QueryRow(ctx, query).Scan(
		&stats.AvgRating,
		&stats.TotalVotes,
		&stats.Rating1,
		&stats.Rating2,
		&stats.Rating3,
		&stats.Rating4,
		&stats.Rating5,
	)
	if err != nil {
		return ReviewStats{}, fmt.Errorf("get review stats: %w", err)
	}

	return stats, nil
}
