package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

func (r *VacantPositionRepository) GetVacantPosition(ctx context.Context, id int) (domain.VacantPosition, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, title, description, date
		FROM vacant_positions
		WHERE id = $1
	`

	var model VacantPositionModel
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.Date,
	)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return domain.VacantPosition{}, fmt.Errorf("vacant position with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.VacantPosition{}, fmt.Errorf("get vacant position: %w", err)
	}

	return vacantPositionDomainFromModel(model), nil
}
