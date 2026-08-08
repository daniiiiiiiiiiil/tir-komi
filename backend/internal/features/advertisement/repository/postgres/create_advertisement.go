package postgres

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (r *AdvertisementRepository) CreateAdvertisement(ctx context.Context, ad domain.Advertisement) (domain.Advertisement, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO advertisement (title, description, image, pdf, url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	var createdAd domain.Advertisement
	err := r.pool.QueryRow(ctx, query,
		ad.Title,
		ad.Description,
		ad.Image,
		ad.Pdf,
		ad.Url,
	).Scan(
		&createdAd.ID,
		&createdAd.CreatedAt,
	)
	if err != nil {
		return domain.Advertisement{}, fmt.Errorf("create advertisement: %w", err)
	}

	createdAd.Title = ad.Title
	createdAd.Description = ad.Description
	createdAd.Image = ad.Image
	createdAd.Pdf = ad.Pdf
	createdAd.Url = ad.Url

	return createdAd, nil
}
