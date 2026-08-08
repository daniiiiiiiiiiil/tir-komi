package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

// DeleteReview 	godoc
// @Summary 	Удаление отзыва
// @Description Удалить отзыв в системе
// @Tags 		Reviews
// @Produce 	json
// @Param 		id path int true 					"ID отзыва"
// @Success 	204 								"Успешно удаленный отзыв по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Review not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/reviews/{id} [delete]
func (handler *ReviewHandler) DeleteReview(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	reviewId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get reviewId path value")
		return
	}

	if err := handler.reviewService.DeleteReview(ctx, reviewId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete review")
		return
	}
	responseHandler.NoContentResponse()
}
