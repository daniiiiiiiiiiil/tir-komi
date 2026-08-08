package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *VacantPositionRepository) UpdateVacantPosition(ctx context.Context, id int, patch domain.VacantPositionPatch) (domain.VacantPosition, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	// Сначала получаем текущую вакансию
	current, err := r.GetVacantPosition(ctx, id)
	if err != nil {
		return domain.VacantPosition{}, err
	}

	// Применяем патч
	if err := current.ApplyPatch(patch); err != nil {
		return domain.VacantPosition{}, fmt.Errorf("apply patch: %w", err)
	}

	query := `
		UPDATE vacant_positions SET
		title = $1,
		description = $2,
		date = $3
		WHERE id = $4
		RETURNING id, title, description, date
	`

	var model VacantPositionModel
	err = r.pool.QueryRow(ctx, query,
		current.Title,
		current.Description,
		current.Date,
		id,
	).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.Date,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.VacantPosition{}, fmt.Errorf("vacant position with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.VacantPosition{}, fmt.Errorf("update vacant position: %w", err)
	}

	return vacantPositionDomainFromModel(model), nil
}
