package postgres

import (
	"context"
	"fmt"

	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *VacantPositionRepository) DeleteVacantPosition(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	cmdTag, err := r.pool.Exec(ctx, `DELETE FROM vacant_positions WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete vacant position: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("vacant position with id %d: %w", id, errors_core.ErrNotFound)
	}

	return nil
}
