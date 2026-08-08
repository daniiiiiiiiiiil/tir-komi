package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (r *MethodologicalMaterialRepository) GetMethodologicalMaterials(ctx context.Context, limit, offset int) ([]domain.MethodologicalMaterial, int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	limit, offset = pagination.LimitOffset(limit, offset)

	var total int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM methodological_material`).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count methodological materials: %w", err)
	}

	query := `
		SELECT id, title, description, date, pdf
		FROM methodological_material
		ORDER BY date DESC NULLS LAST
		LIMIT $1 OFFSET $2
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get methodological materials: %w", err)
	}
	defer rows.Close()

	var materials []domain.MethodologicalMaterial
	for rows.Next() {
		var model MethodologicalMaterialModel
		err := rows.Scan(
			&model.ID,
			&model.Title,
			&model.Description,
			&model.Date,
			&model.Pdf,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan methodological material: %w", err)
		}
		materials = append(materials, methodologicalMaterialDomainFromModel(model))
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows error: %w", err)
	}

	return materials, total, nil
}
