package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type GetPostsResponse struct {
	Items []PostDto `json:"items"`
	Total int       `json:"total"`
}

const (
	defaultLimit  = 20
	defaultOffset = 0
)

// GetPosts godoc
// @Summary 	Получение постов
// @Description Получение списка постов с пагинацией
// @Tags 		Posts
// @Produce 	json
// @Param 		limit query int false "Размер страницы с постами"
// @Param 		offset query int false "Смещение страницы с постами"
// @Success 	200 {object} GetPostsResponse "Успешное получение списка постов"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/posts [get]
func (handler *PostHandler) GetPosts(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset query params")
		return
	}

	posts, total, err := handler.postService.GetPosts(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get posts")
		return
	}

	response := GetPostsResponse{
		Items: convertPostDtosFromDomains(posts),
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
