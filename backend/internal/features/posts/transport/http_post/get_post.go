package http_post

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type GetPostResponse PostDto

// GetPost godoc
// @Summary 	Получение поста
// @Description Получение поста по id
// @Tags 		Posts
// @Produce 	json
// @Param 		id path int true "ID поста"
// @Success 	200 {object} GetPostResponse "Успешно найденный пост по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Post not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/posts/{id} [get]
func (handler *PostHandler) GetPost(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	postId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get postId path value")
		return
	}

	postDomain, err := handler.postService.GetPost(ctx, postId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get post")
		return
	}

	response := GetPostResponse(convertPostDtoFromDomain(postDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
