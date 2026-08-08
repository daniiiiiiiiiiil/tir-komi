package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *AdvertisementService) CreateAdvertisement(ctx context.Context, ad domain.Advertisement) (domain.Advertisement, error) {
	if err := ad.Validate(); err != nil {
		return domain.Advertisement{}, fmt.Errorf("validation failed: %w", err)
	}

	created, err := s.repo.CreateAdvertisement(ctx, ad)
	if err != nil {
		return domain.Advertisement{}, fmt.Errorf("create advertisement: %w", err)
	}

	return created, nil
}
