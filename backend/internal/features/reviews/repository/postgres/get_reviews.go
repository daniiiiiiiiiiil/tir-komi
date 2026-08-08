package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (r *ReviewRepository) GetReviews(ctx context.Context, limit, offset int) ([]domain.Review, int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	limit, offset = pagination.LimitOffset(limit, offset)

	var total int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM reviews`).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count reviews: %w", err)
	}

	query := `
		SELECT id, name, email, description, rating, created_at
		FROM reviews
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get reviews: %w", err)
	}
	defer rows.Close()

	var reviews []domain.Review
	for rows.Next() {
		var model ReviewModel
		err := rows.Scan(
			&model.ID,
			&model.Name,
			&model.Email,
			&model.Description,
			&model.Rating,
			&model.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan review: %w", err)
		}
		reviews = append(reviews, reviewDomainFromModel(model))
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows error: %w", err)
	}

	return reviews, total, nil
}
