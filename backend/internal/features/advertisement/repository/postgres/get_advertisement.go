package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	errors_core "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

func (r *AdvertisementRepository) GetAdvertisement(ctx context.Context, id int) (domain.Advertisement, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, title, image, pdf, url, created_at
		FROM advertisement
		WHERE id = $1
	`

	var model AdvertisementModel
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&model.ID,
		&model.Title,
		&model.Image,
		&model.Pdf,
		&model.Url,
		&model.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return domain.Advertisement{}, fmt.Errorf("advertisement with id %d: %w", id, errors_core.ErrNotFound)
		}
		return domain.Advertisement{}, fmt.Errorf("get advertisement: %w", err)
	}

	return advertisementDomainFromModel(model), nil
}
