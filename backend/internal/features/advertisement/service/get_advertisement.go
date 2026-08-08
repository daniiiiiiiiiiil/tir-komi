package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *AdvertisementService) GetAdvertisement(ctx context.Context, id int) (domain.Advertisement, error) {
	ad, err := s.repo.GetAdvertisement(ctx, id)
	if err != nil {
		return domain.Advertisement{}, fmt.Errorf("get advertisement: %w", err)
	}
	return ad, nil
}
