package postgres

import "github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"

type MediaRepository struct {
	pool postgres.Pool
}

func NewMediaRepository(p postgres.Pool) *MediaRepository {
	return &MediaRepository{pool: p}
}
