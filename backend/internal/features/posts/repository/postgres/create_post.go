package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (r *PostRepository) CreatePost(ctx context.Context, post domain.Post) (domain.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO posts (title, description, pdf, image, date, type)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id int
	err := r.pool.QueryRow(ctx, query,
		post.Title,
		post.Description,
		post.Pdf,
		post.Image,
		post.Date,
		string(post.Type),
	).Scan(&id)
	if err != nil {
		return domain.Post{}, fmt.Errorf("create post: %w", err)
	}

	post.ID = id
	return post, nil
}
