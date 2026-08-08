package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
)

func (r *VacantPositionRepository) GetVacantPositions(ctx context.Context, limit, offset int) ([]domain.VacantPosition, int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	limit, offset = pagination.LimitOffset(limit, offset)

	var total int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM vacant_positions`).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("count vacant positions: %w", err)
	}

	query := `
		SELECT id, title, description, date
		FROM vacant_positions
		ORDER BY date DESC NULLS LAST
		LIMIT $1 OFFSET $2
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("get vacant positions: %w", err)
	}
	defer rows.Close()

	var positions []domain.VacantPosition
	for rows.Next() {
		var model VacantPositionModel
		err := rows.Scan(
			&model.ID,
			&model.Title,
			&model.Description,
			&model.Date,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan vacant position: %w", err)
		}
		positions = append(positions, vacantPositionDomainFromModel(model))
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("rows error: %w", err)
	}

	return positions, total, nil
}
