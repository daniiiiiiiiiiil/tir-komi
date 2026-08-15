package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *PostService) UpdatePost(ctx context.Context, id int, patch domain.PostPatch) (domain.Post, error) {
	post, err := s.repo.GetPost(ctx, id)
	if err != nil {
		return domain.Post{}, fmt.Errorf("postRepository.GetPost: %w", err)
	}

	if err := post.ApplyPatch(patch); err != nil {
		return domain.Post{}, fmt.Errorf("apply patch: %w", err)
	}

	updatedPost, err := s.repo.UpdatePost(ctx, id, patch)
	if err != nil {
		return domain.Post{}, fmt.Errorf("postRepository.UpdatePost: %w", err)
	}

	return updatedPost, nil
}
