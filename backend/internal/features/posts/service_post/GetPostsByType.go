package service_post

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *PostService) GetPostsByType(ctx context.Context, postType domain.PostType) ([]domain.Post, error) {
	if !domain.ValidPostTypes[postType] {
		return nil, fmt.Errorf("invalid post type: %s", postType)
	}

	posts, err := s.repo.GetPostsByType(ctx, postType)
	if err != nil {
		return nil, fmt.Errorf("get posts by type: %w", err)
	}
	return posts, nil
}
