package postgres

import (
	"context"
	"fmt"

	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *AdvertisementRepository) DeleteAdvertisement(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	cmdTag, err := r.pool.Exec(ctx, `DELETE FROM advertisement WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete advertisement: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("advertisement with id %d: %w", id, errors_core.ErrNotFound)
	}

	return nil
}
