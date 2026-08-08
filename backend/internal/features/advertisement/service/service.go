package service

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type AdvertisementRepository interface {
	CreateAdvertisement(ctx context.Context, ad domain.Advertisement) (domain.Advertisement, error)
	GetAdvertisement(ctx context.Context, id int) (domain.Advertisement, error)
	GetAdvertisements(ctx context.Context, limit, offset int) ([]domain.Advertisement, int, error)
	UpdateAdvertisement(ctx context.Context, id int, patch domain.AdvertisementPatch) (domain.Advertisement, error)
	DeleteAdvertisement(ctx context.Context, id int) error
}

type AdvertisementService struct {
	repo AdvertisementRepository
}

func NewAdvertisementService(repo AdvertisementRepository) *AdvertisementService {
	return &AdvertisementService{
		repo: repo,
	}
}
