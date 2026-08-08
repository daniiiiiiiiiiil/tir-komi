package service

import (
	"context"
	"fmt"
)

func (s *VacantPositionService) DeleteVacantPosition(ctx context.Context, id int) error {
	if err := s.repo.DeleteVacantPosition(ctx, id); err != nil {
		return fmt.Errorf("delete vacant position: %w", err)
	}
	return nil
}
