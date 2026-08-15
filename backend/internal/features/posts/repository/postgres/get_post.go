package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

func (r *PostRepository) GetPost(ctx context.Context, id int) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, title, description, pdf, image, date, type
		FROM posts
		WHERE id = $1
	`

	var model PostModel
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.Pdf,
		&model.Image,
		&model.Date,
		&model.Type,
	)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return domain.Post{}, fmt.Errorf("post with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.Post{}, fmt.Errorf("get post: %w", err)
	}

	return postDomainFromModel(model), nil
}
