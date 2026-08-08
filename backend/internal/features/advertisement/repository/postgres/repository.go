package postgres

import "github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"

type AdvertisementRepository struct {
	pool postgres.Pool
}

func NewAdvertisementRepository(p postgres.Pool) *AdvertisementRepository {
	return &AdvertisementRepository{pool: p}
}
