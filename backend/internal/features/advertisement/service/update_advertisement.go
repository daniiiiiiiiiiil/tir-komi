package service

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (s *AdvertisementService) UpdateAdvertisement(ctx context.Context, id int, patch domain.AdvertisementPatch) (domain.Advertisement, error) {
	ad, err := s.repo.GetAdvertisement(ctx, id)
	if err != nil {
		return domain.Advertisement{}, fmt.Errorf("advertisementRepository.GetAdvertisement: %w", err)
	}

	if err := ad.ApplyPatch(patch); err != nil {
		return domain.Advertisement{}, fmt.Errorf("apply patch: %w", err)
	}

	updatedAd, err := s.repo.UpdateAdvertisement(ctx, id, patch)
	if err != nil {
		return domain.Advertisement{}, fmt.Errorf("advertisementRepository.UpdateAdvertisement: %w", err)
	}

	return updatedAd, nil
}
