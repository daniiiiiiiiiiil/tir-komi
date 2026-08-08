package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

func (r *AdvertisementRepository) UpdateAdvertisement(ctx context.Context, id int, patch domain.AdvertisementPatch) (domain.Advertisement, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	current, err := r.GetAdvertisement(ctx, id)
	if err != nil {
		return domain.Advertisement{}, err
	}

	if err := current.ApplyPatch(patch); err != nil {
		return domain.Advertisement{}, fmt.Errorf("apply patch: %w", err)
	}

	query := `
		UPDATE advertisement SET
		title = $1,
		description = $2,
		image = $3,
		pdf = $4,
		url = $5
		WHERE id = $6
		RETURNING id, title, description, image, pdf, url, created_at
	`

	var model AdvertisementModel
	err = r.pool.QueryRow(ctx, query,
		current.Title,
		current.Description,
		current.Image,
		current.Pdf,
		current.Url,
		id,
	).Scan(
		&model.ID,
		&model.Title,
		&model.Description,
		&model.Image,
		&model.Pdf,
		&model.Url,
		&model.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Advertisement{}, fmt.Errorf("advertisement with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.Advertisement{}, fmt.Errorf("update advertisement: %w", err)
	}

	return advertisementDomainFromModel(model), nil
}
