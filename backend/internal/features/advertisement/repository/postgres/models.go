package postgres

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type AdvertisementModel struct {
	ID        int
	Title     string
	Image     *string
	Pdf       *string
	Url       *string
	CreatedAt time.Time
}

func advertisementDomainFromModel(model AdvertisementModel) domain.Advertisement {
	return domain.NewAdvertisement(
		model.ID,
		model.Title,
		model.Image,
		model.Pdf,
		model.Url,
		model.CreatedAt,
	)
}

func advertisementModelsFromDomains(ads []domain.Advertisement) []AdvertisementModel {
	models := make([]AdvertisementModel, len(ads))
	for i, ad := range ads {
		models[i] = AdvertisementModel{
			ID:        ad.ID,
			Title:     ad.Title,
			Image:     ad.Image,
			Pdf:       ad.Pdf,
			Url:       ad.Url,
			CreatedAt: ad.CreatedAt,
		}
	}
	return models
}
