package service_post

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *PostService) CreatePost(ctx context.Context, post domain.Post) (domain.Post, error) {
	if err := post.Validate(); err != nil {
		return domain.Post{}, fmt.Errorf("validation failed: %w", err)
	}

	created, err := s.repo.CreatePost(ctx, post)
	if err != nil {
		return domain.Post{}, fmt.Errorf("create post: %w", err)
	}

	return created, nil
}
