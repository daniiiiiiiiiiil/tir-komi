package http

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type ReviewDto struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	Rating      int       `json:"rating"`
	CreateAt    time.Time `json:"create_at"`
}

func convertReviewDtoFromDomain(review domain.Review) ReviewDto {
	return ReviewDto{
		Id:          review.ID,
		Name:        review.Name,
		Description: review.Description,
		Email:       review.Email,
		Rating:      review.Rating,
		CreateAt:    review.CreatedAt,
	}
}

func convertReviewDtosFromDomains(reviews []domain.Review) []ReviewDto {
	dtos := make([]ReviewDto, len(reviews))
	for i, review := range reviews {
		dtos[i] = convertReviewDtoFromDomain(review)
	}
	return dtos
}
