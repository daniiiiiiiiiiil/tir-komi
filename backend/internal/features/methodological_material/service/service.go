package service

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type MethodologicalMaterialRepository interface {
	CreateMethodologicalMaterial(ctx context.Context, material domain.MethodologicalMaterial) (domain.MethodologicalMaterial, error)
	GetMethodologicalMaterial(ctx context.Context, id int) (domain.MethodologicalMaterial, error)
	GetMethodologicalMaterials(ctx context.Context, limit, offset int) ([]domain.MethodologicalMaterial, int, error)
	UpdateMethodologicalMaterial(ctx context.Context, id int, patch domain.MethodologicalMaterialPatch) (domain.MethodologicalMaterial, error)
	DeleteMethodologicalMaterial(ctx context.Context, id int) error
}

type MethodologicalMaterialService struct {
	repo MethodologicalMaterialRepository
}

func NewMethodologicalMaterialService(repo MethodologicalMaterialRepository) *MethodologicalMaterialService {
	return &MethodologicalMaterialService{
		repo: repo,
	}
}
