package postgres

import (
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

type VacantPositionRepository struct {
	pool postgres.Pool
}

func NewVacantPositionRepository(p postgres.Pool) *VacantPositionRepository {
	return &VacantPositionRepository{pool: p}
}
