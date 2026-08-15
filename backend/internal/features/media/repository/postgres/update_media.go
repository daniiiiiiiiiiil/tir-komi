package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *MediaRepository) UpdateMedia(ctx context.Context, id int, patch domain.MediaPatch) (domain.Media, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	current, err := r.GetMedia(ctx, id)
	if err != nil {
		return domain.Media{}, err
	}

	if err := current.ApplyPatch(patch); err != nil {
		return domain.Media{}, fmt.Errorf("apply patch: %w", err)
	}

	query := `
		UPDATE media SET
		title = $1,
		description = $2,
		file_data = $3,
		file_name = $4,
		mime_type = $5,
		file_size = $6,
		sort_order = $7
		WHERE id = $8
		RETURNING id, type, title, description, file_name, mime_type, file_size, sort_order, created_at
	`

	var model MediaModel
	err = r.pool.QueryRow(ctx, query,
		current.Title,
		current.Description,
		current.FileData,
		current.FileName,
		current.MimeType,
		current.FileSize,
		current.SortOrder,
		id,
	).Scan(
		&model.ID,
		&model.Type,
		&model.Title,
		&model.Description,
		&model.FileName,
		&model.MimeType,
		&model.FileSize,
		&model.SortOrder,
		&model.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Media{}, fmt.Errorf("media with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.Media{}, fmt.Errorf("update media: %w", err)
	}

	return mediaDomainFromModel(model), nil
}
