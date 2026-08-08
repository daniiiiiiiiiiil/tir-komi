package service

import (
	"context"
	"fmt"
)

func (s *AdvertisementService) DeleteAdvertisement(ctx context.Context, id int) error {
	_, err := s.repo.GetAdvertisement(ctx, id)
	if err != nil {
		return fmt.Errorf("get advertisement: %w", err)
	}
	if err := s.repo.DeleteAdvertisement(ctx, id); err != nil {
		return fmt.Errorf("delete advertisement: %w", err)
	}
	return nil
}
