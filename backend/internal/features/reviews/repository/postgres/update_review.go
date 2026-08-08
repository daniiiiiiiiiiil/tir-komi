package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *ReviewRepository) UpdateReview(ctx context.Context, id int, patch domain.ReviewPatch) (domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	current, err := r.GetReview(ctx, id)
	if err != nil {
		return domain.Review{}, err
	}

	if err := current.ApplyPatch(patch); err != nil {
		return domain.Review{}, fmt.Errorf("apply patch: %w", err)
	}

	query := `
		UPDATE reviews SET
		name = $1,
		email = $2,
		rating = $3
		WHERE id = $4
		RETURNING id, name, email, rating, created_at
	`

	var model ReviewModel
	err = r.pool.QueryRow(ctx, query,
		current.Name,
		current.Email,
		current.Rating,
		id,
	).Scan(
		&model.ID,
		&model.Name,
		&model.Email,
		&model.Rating,
		&model.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Review{}, fmt.Errorf("review with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.Review{}, fmt.Errorf("update review: %w", err)
	}

	return reviewDomainFromModel(model), nil
}
