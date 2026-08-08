package postgres

import (
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

type ReviewRepository struct {
	pool postgres.Pool
}

func NewReviewRepository(p postgres.Pool) *ReviewRepository {
	return &ReviewRepository{pool: p}
}
