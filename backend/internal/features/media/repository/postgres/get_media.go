package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *MediaRepository) GetMedia(ctx context.Context, id int) (domain.Media, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, type, title, description, file_data, file_name, mime_type, file_size, sort_order, created_at
		FROM media
		WHERE id = $1
	`

	var model MediaModel
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Type,
		&model.Title,
		&model.Description,
		&model.FileData,
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
		return domain.Media{}, fmt.Errorf("get media: %w", err)
	}

	return mediaDomainFromModel(model), nil
}
