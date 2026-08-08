package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

func (r *ReviewRepository) GetReview(ctx context.Context, id int) (domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, name, email, rating, created_at
		FROM reviews
		WHERE id = $1
	`

	var model ReviewModel
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Name,
		&model.Email,
		&model.Rating,
		&model.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return domain.Review{}, fmt.Errorf("review with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.Review{}, fmt.Errorf("get review: %w", err)
	}

	return reviewDomainFromModel(model), nil
}
