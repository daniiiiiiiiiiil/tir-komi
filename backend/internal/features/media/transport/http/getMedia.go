package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/pagination"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type GetMediaResponse MediaDto

// GetMedia 	godoc
// @Summary 	Получение фото/видео
// @Tags 		Media
// @Produce 	json
// @Param 		id path int true "ID элемента"
// @Success 	200 {object} GetMediaResponse
// @Failure 	404 {object} response.ErrorResponse "Media not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/media/{id} [get]
func (handler *MediaHandler) GetMedia(rw http.ResponseWriter, r *http.Request) {
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

	response := GetMediaResponse(convertMediaDtoFromDomain(m))
	responseHandler.JsonResponse(response, http.StatusOK)
}

type GetMediasResponse struct {
	Items []MediaDto `json:"items"`
	Total int        `json:"total"`
}

// GetMedias 	godoc
// @Summary 	Список фото/видео
// @Description Получить список элементов галереи, можно отфильтровать по type=photo|video
// @Tags 		Media
// @Produce 	json
// @Param 		type   query string false "Фильтр по типу: photo или video"
// @Param 		limit  query int    false "Лимит"
// @Param 		offset query int    false "Смещение"
// @Success 	200 {object} GetMediasResponse
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/media [get]
func (handler *MediaHandler) GetMedias(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := pagination.GetPagination(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get pagination params")
		return
	}

	var mediaType *domain.MediaType
	if t := r.URL.Query().Get("type"); t != "" {
		mt := domain.MediaType(t)
		mediaType = &mt
	}

	items, total, err := handler.mediaService.GetMedias(ctx, mediaType, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get medias")
		return
	}

	response := GetMediasResponse{
		Items: convertMediaDtosFromDomains(items),
		Total: total,
	}
	responseHandler.JsonResponse(response, http.StatusOK)
}
