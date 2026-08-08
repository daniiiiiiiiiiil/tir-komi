package postgres

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type ReviewModel struct {
	ID        int
	Name      string
	Email     *string
	Rating    int
	CreatedAt time.Time
}

func reviewDomainFromModel(model ReviewModel) domain.Review {
	return domain.NewReview(
		model.ID,
		model.Name,
		model.Email,
		model.Rating,
		model.CreatedAt,
	)
}

func reviewModelsFromDomains(reviews []domain.Review) []ReviewModel {
	models := make([]ReviewModel, len(reviews))
	for i, review := range reviews {
		models[i] = ReviewModel{
			ID:        review.ID,
			Name:      review.Name,
			Email:     review.Email,
			Rating:    review.Rating,
			CreatedAt: review.CreatedAt,
		}
	}
	return models
}
