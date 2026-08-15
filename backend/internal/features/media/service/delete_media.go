package service

import (
	"context"
	"fmt"
)

func (s *MediaService) DeleteMedia(ctx context.Context, id int) error {
	_, err := s.repo.GetMedia(ctx, id)
	if err != nil {
		return fmt.Errorf("get media: %w", err)
	}
	if err := s.repo.DeleteMedia(ctx, id); err != nil {
		return fmt.Errorf("delete media: %w", err)
	}
	return nil
}
