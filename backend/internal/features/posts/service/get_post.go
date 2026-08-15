package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *PostService) GetPost(ctx context.Context, id int) (domain.Post, error) {
	post, err := s.repo.GetPost(ctx, id)
	if err != nil {
		return domain.Post{}, fmt.Errorf("get post: %w", err)
	}
	return post, nil
}
