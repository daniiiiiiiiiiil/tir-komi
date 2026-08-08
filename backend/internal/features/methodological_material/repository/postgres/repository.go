package postgres

import (
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

type MethodologicalMaterialRepository struct {
	pool postgres.Pool
}

func NewMethodologicalMaterialRepository(p postgres.Pool) *MethodologicalMaterialRepository {
	return &MethodologicalMaterialRepository{pool: p}
}
