package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *MethodologicalMaterialRepository) UpdateMethodologicalMaterial(ctx context.Context, id int, patch domain.MethodologicalMaterialPatch) (domain.MethodologicalMaterial, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	current, err := r.GetMethodologicalMaterial(ctx, id)
	if err != nil {
		return domain.MethodologicalMaterial{}, err
	}

	if err := current.ApplyPatch(patch); err != nil {
		return domain.MethodologicalMaterial{}, fmt.Errorf("apply patch: %w", err)
	}

	query := `
		UPDATE methodological_material SET
		title = $1,
		description = $2,
		date = $3,
		pdf = $4
		WHERE id = $5
		RETURNING id, title, description, date, pdf
	`

	var model MethodologicalMaterialModel
	err = r.pool.QueryRow(ctx, query,
		current.Title,
		current.Description,
		current.Date,
		current.Pdf,
		id,
	).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.Date,
		&model.Pdf,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.MethodologicalMaterial{}, fmt.Errorf("methodological material with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.MethodologicalMaterial{}, fmt.Errorf("update methodological material: %w", err)
	}

	return methodologicalMaterialDomainFromModel(model), nil
}
