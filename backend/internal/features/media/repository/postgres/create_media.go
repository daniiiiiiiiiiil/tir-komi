package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (r *MediaRepository) CreateMedia(ctx context.Context, m domain.Media) (domain.Media, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO media (type, title, description, file_data, file_name, mime_type, file_size, sort_order)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at
	`

	var created domain.Media
	err := r.pool.QueryRow(ctx, query,
		string(m.Type),
		m.Title,
		m.Description,
		m.FileData,
		m.FileName,
		m.MimeType,
		m.FileSize,
		m.SortOrder,
	).Scan(
		&created.ID,
		&created.CreatedAt,
	)
	if err != nil {
		return domain.Media{}, fmt.Errorf("create media: %w", err)
	}

	created.Type = m.Type
	created.Title = m.Title
	created.Description = m.Description
	created.FileData = m.FileData
	created.FileName = m.FileName
	created.MimeType = m.MimeType
	created.FileSize = m.FileSize
	created.SortOrder = m.SortOrder

	return created, nil
}
