package http

import (
	"net/http"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type CreatePostRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=1000" example:"Новый пост"`
	Description string `json:"description" validate:"omitempty,max=10000" example:"Описание поста"`
	Pdf         []byte `json:"pdf,omitempty"`
	Image       []byte `json:"image,omitempty"`
	Type        string `json:"type" validate:"required" example:"information"`
}

type CreatePostResponse PostDto

// CreatePost godoc
// @Summary 	Создание поста
// @Description Создать новый пост в системе
// @Tags 		Posts
// @Accept 		json
// @Produce 	json
// @Param 		request body CreatePostRequest true "Тело запроса"
// @Success 	201 {object} CreatePostResponse "Успешно созданный пост"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/posts [post]
func (handler *PostHandler) CreatePost(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	var req CreatePostRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode the request")
		return
	}

	description := &req.Description
	date := time.Now()
	postType := domain.PostType(req.Type)

	postDomain := domain.NewPostUninitialized(
		req.Title,
		description,
		req.Pdf,
		req.Image,
		&date,
		postType,
	)
	if err := postDomain.Validate(); err != nil {
		responseHandler.ErrorResponse(err, "invalid post")
		return
	}

	created, err := handler.postService.CreatePost(ctx, postDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create post")
		return
	}

	response := CreatePostResponse(convertPostDtoFromDomain(created))
	responseHandler.JsonResponse(response, http.StatusCreated)
}
