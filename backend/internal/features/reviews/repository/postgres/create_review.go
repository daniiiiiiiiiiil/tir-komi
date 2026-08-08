package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (r *ReviewRepository) CreateReview(ctx context.Context, review domain.Review) (domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO reviews (name, email, description, rating)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`

	var createdReview domain.Review
	err := r.pool.QueryRow(ctx, query,
		review.Name,
		review.Email,
		review.Description,
		review.Rating,
	).Scan(
		&createdReview.ID,
		&createdReview.CreatedAt,
	)
	if err != nil {
		return domain.Review{}, fmt.Errorf("create review: %w", err)
	}

	createdReview.Name = review.Name
	createdReview.Email = review.Email
	createdReview.Description = review.Description
	createdReview.Rating = review.Rating

	return createdReview, nil
}
