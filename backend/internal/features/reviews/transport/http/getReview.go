package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type ReviewResponse ReviewDto

// GetReview 	godoc
// @Summary 	Получение отзыва
// @Description Получение отзыва по id
// @Tags 		Reviews
// @Produce 	json
// @Param 		id path int true 					"ID отзыва"
// @Success 	200 {object} ReviewResponse 		"Успешно найденный отзыв по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Review not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/reviews/{id} [get]
func (handler *ReviewHandler) GetReview(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	reviewId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get reviewId path value")
		return
	}

	reviewDomain, err := handler.reviewService.GetReview(ctx, reviewId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get review")
		return
	}

	response := ReviewResponse(convertReviewDtoFromDomain(reviewDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
