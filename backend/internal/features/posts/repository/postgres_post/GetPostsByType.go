package postgres_post

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (r *PostRepository) GetPostsByType(ctx context.Context, postType domain.PostType) ([]domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, title, description, pdf, image, date, type
		FROM posts
		WHERE type = $1
		ORDER BY date DESC
	`

	rows, err := r.pool.Query(ctx, query, string(postType))
	if err != nil {
		return nil, fmt.Errorf("get posts by type: %w", err)
	}
	defer rows.Close()

	var models []PostModel
	for rows.Next() {
		var model PostModel
		if err := rows.Scan(
			&model.ID,
			&model.Title,
			&model.Description,
			&model.Pdf,
			&model.Image,
			&model.Date,
			&model.Type,
		); err != nil {
			return nil, fmt.Errorf("scan post: %w", err)
		}
		models = append(models, model)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get posts by type: %w", err)
	}

	posts := make([]domain.Post, len(models))
	for i, model := range models {
		posts[i] = postDomainFromModel(model)
	}

	return posts, nil
}
