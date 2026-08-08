package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

// GetReviewsByRating 	godoc
// @Summary 	Получение отзывов по рейтингу
// @Description Получение списка отзывов, отфильтрованных по рейтингу, с паггинацией
// @Tags 		Reviews
// @Produce 	json
// @Param 		rating path int true 				"Рейтинг отзыва"
// @Param 		limit query int false 				"Размер страницы с отзывами"
// @Param 		offset query int false 			"Смещение страницы с отзывами"
// @Success 	200 {object} ReviewsResponse 		"Успешное получение списка отзывов"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/reviews/rating/{rating} [get]
func (handler *ReviewHandler) GetReviewsByRating(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	rating, err := requests.GetIntPathValue(r, "rating")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get rating path value")
		return
	}

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset query params")
		return
	}

	reviews, total, err := handler.reviewService.GetReviewsByRating(ctx, rating, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get reviews by rating")
		return
	}

	response := ReviewsResponse{
		Items: convertReviewDtosFromDomains(reviews),
		Total: total,
	}
	responseHandler.JsonResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (int, int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	limit := defaultLimit
	if limitPtr, err := requests.GetIntQueryParams(r, limitQueryParamKey); err != nil {
		return 0, 0, err
	} else if limitPtr != nil {
		limit = *limitPtr
	}

	offset := defaultOffset
	if offsetPtr, err := requests.GetIntQueryParams(r, offsetQueryParamKey); err != nil {
		return 0, 0, err
	} else if offsetPtr != nil {
		offset = *offsetPtr
	}

	return limit, offset, nil
}
