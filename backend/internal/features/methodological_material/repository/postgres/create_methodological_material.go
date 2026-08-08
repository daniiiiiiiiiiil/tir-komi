package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (r *MethodologicalMaterialRepository) CreateMethodologicalMaterial(ctx context.Context, material domain.MethodologicalMaterial) (domain.MethodologicalMaterial, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO methodological_material (title, description, date, pdf)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int
	err := r.pool.QueryRow(ctx, query,
		material.Title,
		material.Description,
		material.Date,
		material.Pdf,
	).Scan(&id)
	if err != nil {
		return domain.MethodologicalMaterial{}, fmt.Errorf("create methodological material: %w", err)
	}

	material.ID = id
	return material, nil
}
