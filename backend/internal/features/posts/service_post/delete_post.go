package service_post

import (
	"context"
	"fmt"
)

func (s *PostService) DeletePost(ctx context.Context, id int) error {
	_, err := s.repo.GetPost(ctx, id)
	if err != nil {
		return fmt.Errorf("get post: %w", err)
	}
	if err := s.repo.DeletePost(ctx, id); err != nil {
		return fmt.Errorf("delete post: %w", err)
	}
	return nil
}
