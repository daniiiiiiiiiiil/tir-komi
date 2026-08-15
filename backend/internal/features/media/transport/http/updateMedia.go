package http

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type UpdateMediaResponse MediaDto

// UpdateMedia 	godoc
// @Summary     Изменение фото/видео
// @Description Изменить подпись/порядок, и опционально заменить сам файл новым (загрузка с ПК).
// @Tags 		Media
// @Accept 		multipart/form-data
// @Produce 	json
// @Param 		id          path     int    true  "ID элемента"
// @Param 		title       formData string false "Название"
// @Param 		description formData string false "Описание"
// @Param 		sort_order  formData int    false "Порядок сортировки"
// @Param 		file        formData file   false "Новый файл (если нужно заменить)"
// @Success 	200 {object} UpdateMediaResponse
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Media not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/media/{id} [patch]
func (handler *MediaHandler) UpdateMedia(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	mediaId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get mediaId path value")
		return
	}

	r.Body = http.MaxBytesReader(rw, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		responseHandler.ErrorResponse(err, "failed to parse multipart form")
		return
	}

	var patch domain.MediaPatch

	if r.Form.Has("title") {
		v := r.FormValue("title")
		patch.Title = domain.Nullable[string]{Set: true, Value: &v}
	}

	if r.Form.Has("description") {
		v := r.FormValue("description")
		patch.Description = domain.Nullable[string]{Set: true, Value: &v}
	}

	if r.Form.Has("sort_order") {
		if v, err := parseIntOrZero(r.FormValue("sort_order")); err == nil {
			patch.SortOrder = domain.Nullable[int]{Set: true, Value: &v}
		}
	}

	if file, header, err := r.FormFile("file"); err == nil {
		defer file.Close()

		fileData, err := io.ReadAll(file)
		if err != nil {
			responseHandler.ErrorResponse(err, "failed to read file content")
			return
		}

		mimeType := header.Header.Get("Content-Type")
		if mimeType == "" || mimeType == "application/octet-stream" {
			ext := filepath.Ext(header.Filename)
			switch ext {
			case ".png":
				mimeType = "image/png"
			case ".jpg", ".jpeg":
				mimeType = "image/jpeg"
			case ".gif":
				mimeType = "image/gif"
			case ".mp4":
				mimeType = "video/mp4"
			case ".webm":
				mimeType = "video/webm"
			}
		}

		patch.FileData = domain.Nullable[[]byte]{Set: true, Value: &fileData}
		patch.FileName = domain.Nullable[string]{Set: true, Value: &header.Filename}
		patch.MimeType = domain.Nullable[string]{Set: true, Value: &mimeType}
	}

	updated, err := handler.mediaService.UpdateMedia(ctx, mediaId, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update media")
		return
	}

	response := UpdateMediaResponse(convertMediaDtoFromDomain(updated))
	responseHandler.JsonResponse(response, http.StatusOK)
}
