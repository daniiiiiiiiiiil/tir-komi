package http_post

import (
	"fmt"
	"net/http"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type UpdatePostRequest struct {
	Title       domain.Nullable[string]  `json:"title"`
	Description domain.Nullable[*string] `json:"description"`
	Pdf         domain.Nullable[[]byte]  `json:"pdf"`
	Image       domain.Nullable[[]byte]  `json:"image"`
	Type        domain.Nullable[*string] `json:"type"`
}

func (r *UpdatePostRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("title cannot be NULL")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 1000 {
			return fmt.Errorf("title length must be between 1 and 1000")
		}
	}

	if r.Pdf.Set && r.Pdf.Value != nil {
		if len(*r.Pdf.Value) > 100*1024*1024 { // 50 MB
			return fmt.Errorf("pdf size must not exceed 50 MB")
		}
	}

	if r.Image.Set && r.Image.Value != nil {
		if len(*r.Image.Value) > 100*1024*1024 { // 10 MB
			return fmt.Errorf("image size must not exceed 10 MB")
		}
	}

	return nil
}

type UpdatePostResponse PostDto

// UpdatePost godoc
// @Summary     Изменение поста
// @Description Изменение информации о посте
// @Description ### Логика обновления полей:
// @Description 1. **Поле не передано** - значение в БД не меняется
// @Description 2. **Передан null** - удаление поля из БД (для title и type недопустимо)
// @Description 3. **Передано значение** - обновление в БД
// @Tags 		Posts
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "ID поста"
// @Param 		request body UpdatePostRequest true "Тело запроса"
// @Success 	200 {object} UpdatePostResponse "Успешно измененный пост"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Post not found"
// @Failure 	409 {object} response.ErrorResponse "Conflict"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/posts/{id} [patch]
func (handler *PostHandler) UpdatePost(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	postId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get postId path value")
		return
	}

	var req UpdatePostRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request body")
		return
	}

	var typeNullable domain.Nullable[*domain.PostType]
	if req.Type.Set {
		if req.Type.Value != nil {
			pt := domain.PostType(**req.Type.Value)
			ptr := &pt
			typeNullable = domain.Nullable[*domain.PostType]{Set: true, Value: &ptr}
		} else {
			var nilPtr *domain.PostType
			typeNullable = domain.Nullable[*domain.PostType]{Set: true, Value: &nilPtr}
		}
	}

	patch := domain.NewPostPatch(
		req.Title,
		req.Description,
		req.Pdf,
		req.Image,
		domain.Nullable[*time.Time]{},
		typeNullable,
	)

	updated, err := handler.postService.UpdatePost(ctx, postId, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update post")
		return
	}

	response := UpdatePostResponse(convertPostDtoFromDomain(updated))
	responseHandler.JsonResponse(response, http.StatusOK)
}
