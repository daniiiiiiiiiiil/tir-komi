package http

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type ReviewHandler struct {
	reviewService ReviewService
}

type ReviewService interface {
	CreateReview(ctx context.Context, review domain.Review) (domain.Review, error)
	DeleteReview(ctx context.Context, id int) error
	GetReviewsByRating(ctx context.Context, rating int, limit, offset int) ([]domain.Review, int, error)
	GetReview(ctx context.Context, id int) (domain.Review, error)
	GetReviews(ctx context.Context, limit, offset int) ([]domain.Review, int, error)
	UpdateReview(ctx context.Context, id int, patch domain.ReviewPatch) (domain.Review, error)
}

func NewReviewHandler(reviewService ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
	}
}

func (handler *ReviewHandler) Routers() []server.Route {
	return []server.Route{
		{
			Method:  "POST",
			Path:    "/reviews",
			Handler: handler.CreateReview,
		},
		{
			Method:  "GET",
			Path:    "/reviews",
			Handler: handler.GetReviews,
		},
		{
			Method:  "GET",
			Path:    "/reviews/{id}",
			Handler: handler.GetReview,
		},
		{
			Method:  "DELETE",
			Path:    "/reviews/{id}",
			Handler: handler.DeleteReview,
		},
		{
			Method:  "PATCH",
			Path:    "/reviews/{id}",
			Handler: handler.UpdateReview,
		},
	}
}
