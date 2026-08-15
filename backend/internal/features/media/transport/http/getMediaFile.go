package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

// GetMediaFile 	godoc
// @Summary 	Получение файла фото/видео
// @Description Отдаёт бинарное содержимое файла (для <img>/<video> src)
// @Tags 		Media
// @Produce 	application/octet-stream
// @Param 		id path int true "ID элемента"
// @Success 	200 {file} binary
// @Failure 	404 {object} response.ErrorResponse "Media not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/media/{id}/file [get]
func (handler *MediaHandler) GetMediaFile(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	mediaId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get mediaId path value")
		return
	}

	m, err := handler.mediaService.GetMedia(ctx, mediaId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get media")
		return
	}

	rw.Header().Set("Content-Type", m.MimeType)
	rw.Header().Set("Content-Length", strconv.FormatInt(m.FileSize, 10))
	rw.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, m.FileName))
	rw.WriteHeader(http.StatusOK)
	rw.Write(m.FileData)
}
