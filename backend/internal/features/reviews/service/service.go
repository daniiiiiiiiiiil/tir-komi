package service

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type ReviewRepository interface {
	CreateReview(ctx context.Context, review domain.Review) (domain.Review, error)
	GetReview(ctx context.Context, id int) (domain.Review, error)
	GetReviews(ctx context.Context, limit, offset int) ([]domain.Review, int, error)
	GetReviewsByRating(ctx context.Context, rating int, limit, offset int) ([]domain.Review, int, error)
	UpdateReview(ctx context.Context, id int, patch domain.ReviewPatch) (domain.Review, error)
	DeleteReview(ctx context.Context, id int) error
}

type ReviewStats struct {
	AvgRating  float64 `json:"avg_rating"`
	TotalVotes int     `json:"total_votes"`
	Rating1    int     `json:"rating_1"`
	Rating2    int     `json:"rating_2"`
	Rating3    int     `json:"rating_3"`
	Rating4    int     `json:"rating_4"`
	Rating5    int     `json:"rating_5"`
}

type ReviewService struct {
	repo ReviewRepository
}

func NewReviewService(repo ReviewRepository) *ReviewService {
	return &ReviewService{
		repo: repo,
	}
}
