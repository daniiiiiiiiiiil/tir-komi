package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

// DeleteMedia 	godoc
// @Summary 	Удаление фото/видео
// @Description Удалить элемент из галереи
// @Tags 		Media
// @Produce 	json
// @Param 		id path int true 					"ID элемента"
// @Success 	204 								"Успешно удалённый элемент"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Media not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/media/{id} [delete]
func (handler *MediaHandler) DeleteMedia(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	mediaId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get mediaId path value")
		return
	}

	if err := handler.mediaService.DeleteMedia(ctx, mediaId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete media")
		return
	}
	responseHandler.NoContentResponse()
}
