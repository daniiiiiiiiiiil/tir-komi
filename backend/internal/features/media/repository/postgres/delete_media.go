package postgres

import (
	"context"
	"fmt"

	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *MediaRepository) DeleteMedia(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	cmdTag, err := r.pool.Exec(ctx, `DELETE FROM media WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete media: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("media with id %d: %w", id, errors_core.ErrNotFound)
	}

	return nil
}
