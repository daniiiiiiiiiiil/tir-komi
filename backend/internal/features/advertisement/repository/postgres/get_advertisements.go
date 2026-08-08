package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (r *AdvertisementRepository) GetAdvertisements(ctx context.Context, limit, offset int) ([]domain.Advertisement, int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	limit, offset = pagination.LimitOffset(limit, offset)

	var total int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM advertisement`).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count advertisements: %w", err)
	}

	query := `
		SELECT id, title, description, image, pdf, url, created_at
		FROM advertisement
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get advertisements: %w", err)
	}
	defer rows.Close()

	var advertisements []domain.Advertisement
	for rows.Next() {
		var model AdvertisementModel
		err := rows.Scan(
			&model.ID,
			&model.Title,
			&model.Description,
			&model.Image,
			&model.Pdf,
			&model.Url,
			&model.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan advertisement: %w", err)
		}
		advertisements = append(advertisements, advertisementDomainFromModel(model))
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows error: %w", err)
	}

	return advertisements, total, nil
}
