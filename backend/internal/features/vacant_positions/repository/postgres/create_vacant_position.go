package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (r *VacantPositionRepository) CreateVacantPosition(ctx context.Context, position domain.VacantPosition) (domain.VacantPosition, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO vacant_positions (title, description, date)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	var id int
	err := r.pool.QueryRow(ctx, query,
		position.Title,
		position.Description,
		position.Date,
	).Scan(&id)
	if err != nil {
		return domain.VacantPosition{}, fmt.Errorf("create vacant position: %w", err)
	}

	position.ID = id
	return position, nil
}
