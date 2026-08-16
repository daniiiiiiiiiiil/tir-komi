package http_post

import (
	"errors"
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type GetPostsByTypeResponse []PostDto

// GetPostsByType godoc
// @Summary 	Получение постов по типу
// @Description Получение списка постов, отфильтрованных по типу
// @Tags 		Posts
// @Produce 	json
// @Param 		type query string true "Тип поста"
// @Success 	200 {object} GetPostsByTypeResponse "Успешно найденные посты по типу"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/posts [get]
func (handler *PostHandler) GetPostsByType(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	postType := domain.PostType(r.URL.Query().Get("type"))
	if postType == "" {
		responseHandler.ErrorResponse(errors.New("type query parameter is required"), "failed to get type query value")
		return
	}

	postsDomain, err := handler.postService.GetPostsByType(ctx, postType)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get posts by type")
		return
	}

	dtos := make(GetPostsByTypeResponse, len(postsDomain))
	for i, p := range postsDomain {
		dtos[i] = PostDto(convertPostDtoFromDomain(p))
	}

	responseHandler.JsonResponse(dtos, http.StatusOK)
}
