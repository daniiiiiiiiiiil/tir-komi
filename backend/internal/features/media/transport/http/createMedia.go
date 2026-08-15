package http

import (
	"io"
	"net/http"
	"path/filepath"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

const maxUploadSize = 1000 << 20

type CreateMediaResponse MediaDto

// CreateMedia 	godoc
// @Summary 	Добавление фото/видео
// @Description Загрузить файл (фото или видео) с ПК в галерею. Поддерживаются любые форматы.
// @Tags 		Media
// @Accept 		multipart/form-data
// @Produce 	json
// @Param 		type        formData string true  "photo или video"
// @Param 		title       formData string false "Название"
// @Param 		description formData string false "Описание"
// @Param 		sort_order  formData int    false "Порядок сортировки"
// @Param 		file        formData file   true  "Файл фото/видео"
// @Success 	201 {object} CreateMediaResponse
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/media [post]
func (handler *MediaHandler) CreateMedia(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	r.Body = http.MaxBytesReader(rw, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		responseHandler.ErrorResponse(err, "failed to parse multipart form (file too large or malformed)")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to read uploaded file")
		return
	}
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
	mediaTypeRaw := r.FormValue("type")

	var title *string
	if v := r.FormValue("title"); v != "" {
		title = &v
	}

	var description *string
	if v := r.FormValue("description"); v != "" {
		description = &v
	}

	sortOrder := 0
	if v := r.FormValue("sort_order"); v != "" {
		if parsed, err := parseIntOrZero(v); err == nil {
			sortOrder = parsed
		}
	}

	mediaDomain := domain.NewMediaUninitialized(
		domain.MediaType(mediaTypeRaw),
		title,
		description,
		fileData,
		header.Filename,
		mimeType,
		sortOrder,
	)

	if err := mediaDomain.Validate(); err != nil {
		responseHandler.ErrorResponse(err, "invalid media")
		return
	}

	created, err := handler.mediaService.CreateMedia(ctx, mediaDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create media")
		return
	}

	response := CreateMediaResponse(convertMediaDtoFromDomain(created))
	responseHandler.JsonResponse(response, http.StatusCreated)
}
