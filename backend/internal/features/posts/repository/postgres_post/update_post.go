package postgres_post

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *PostRepository) UpdatePost(ctx context.Context, id int, patch domain.PostPatch) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	current, err := r.GetPost(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}

	if err := current.ApplyPatch(patch); err != nil {
		return domain.Post{}, fmt.Errorf("apply patch: %w", err)
	}

	query := `
		UPDATE posts SET
		title = $1,
		description = $2,
		pdf = $3,
		image = $4,
		date = $5,
		type = $6
		WHERE id = $7
		RETURNING id, title, description, pdf, image, date, type
	`

	var model PostModel
	err = r.pool.QueryRow(ctx, query,
		current.Title,
		current.Description,
		current.Pdf,
		current.Image,
		current.Date,
		string(current.Type),
		id,
	).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.Pdf,
		&model.Image,
		&model.Date,
		&model.Type,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Post{}, fmt.Errorf("post with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.Post{}, fmt.Errorf("update post: %w", err)
	}

	return postDomainFromModel(model), nil
}
