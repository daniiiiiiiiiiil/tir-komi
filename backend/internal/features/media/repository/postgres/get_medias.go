package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

func (r *MediaRepository) GetMedias(ctx context.Context, mediaType *domain.MediaType, limit, offset int) ([]domain.Media, int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var (
		rows  postgres.Rows
		err   error
		total int
	)

	if mediaType != nil {
		if err = r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM media WHERE type = $1`, string(*mediaType)).Scan(&total); err != nil {
			return nil, 0, fmt.Errorf("count media: %w", err)
		}
		rows, err = r.pool.Query(ctx, `
			SELECT id, type, title, description, file_name, mime_type, file_size, sort_order, created_at
			FROM media
			WHERE type = $1
			ORDER BY sort_order, created_at DESC
			LIMIT $2 OFFSET $3
		`, string(*mediaType), limit, offset)
	} else {
		if err = r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM media`).Scan(&total); err != nil {
			return nil, 0, fmt.Errorf("count media: %w", err)
		}
		rows, err = r.pool.Query(ctx, `
			SELECT id, type, title, description, file_name, mime_type, file_size, sort_order, created_at
			FROM media
			ORDER BY sort_order, created_at DESC
			LIMIT $1 OFFSET $2
		`, limit, offset)
	}
	if err != nil {
		return nil, 0, fmt.Errorf("query media: %w", err)
	}
	defer rows.Close()

	var models []MediaModel
	for rows.Next() {
		var model MediaModel
		if err := rows.Scan(
			&model.ID,
			&model.Type,
			&model.Title,
			&model.Description,
			&model.FileName,
			&model.MimeType,
			&model.FileSize,
			&model.SortOrder,
			&model.CreatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("scan media: %w", err)
		}
		models = append(models, model)
	}
	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows error: %w", err)
	}

	items := make([]domain.Media, len(models))
	for i, model := range models {
		items[i] = mediaDomainFromModel(model)
	}

	return items, total, nil
}
