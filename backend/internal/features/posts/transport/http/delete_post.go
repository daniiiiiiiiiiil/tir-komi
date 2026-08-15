package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

// DeletePost godoc
// @Summary 	Удаление поста
// @Description Удалить пост в системе
// @Tags 		Posts
// @Produce 	json
// @Param 		id path int true "ID поста"
// @Success 	204 "Успешно удаленный пост по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Post not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/posts/{id} [delete]
func (handler *PostHandler) DeletePost(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	postId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get postId path value")
		return
	}

	if err := handler.postService.DeletePost(ctx, postId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete post")
		return
	}
	responseHandler.NoContentResponse()
}
