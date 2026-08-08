package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type ReviewsResponse struct {
	Items []ReviewDto `json:"items"`
	Total int         `json:"total"`
}

// GetReviews 	godoc
// @Summary 	Получение отзывов
// @Description Получение списка отзывов с паггинацией
// @Tags 		Reviews
// @Produce 	json
// @Param 		limit query int false 				"Размер страницы с отзывами"
// @Param 		offset query int false 			"Смещение страницы с отзывами"
// @Success 	200 {object} ReviewsResponse 		"Успешное получение списка отзывов"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/reviews [get]
func (handler *ReviewHandler) GetReviews(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset query params")
		return
	}

	reviews, total, err := handler.reviewService.GetReviews(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get reviews")
		return
	}

	response := ReviewsResponse{
		Items: convertReviewDtosFromDomains(reviews),
		Total: total,
	}
	responseHandler.JsonResponse(response, http.StatusOK)
}
