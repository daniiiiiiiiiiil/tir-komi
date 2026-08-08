package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

func (r *MethodologicalMaterialRepository) GetMethodologicalMaterial(ctx context.Context, id int) (domain.MethodologicalMaterial, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, title, description, date, pdf
		FROM methodological_material
		WHERE id = $1
	`

	var model MethodologicalMaterialModel
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.Date,
		&model.Pdf,
	)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return domain.MethodologicalMaterial{}, fmt.Errorf("methodological material with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.MethodologicalMaterial{}, fmt.Errorf("get methodological material: %w", err)
	}

	return methodologicalMaterialDomainFromModel(model), nil
}
