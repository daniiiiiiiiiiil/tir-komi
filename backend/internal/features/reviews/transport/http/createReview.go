package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type CreateReviewRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=100" example:"Иван Иванов"`
	Email       *string `json:"email" validate:"required,email,max=250" example:"ivan@example.com"`
	Description *string `json:"description" validate:"required,min=1,max=1000" example:"Отличный сервис, всё понравилось"`
	Rating      int     `json:"rating" validate:"required,min=1,max=5" example:"5"`
}

type CreateReviewResponse ReviewDto

// CreateReview 	godoc
// @Summary 	Создание отзыва
// @Description Создать новый отзыв в системе
// @Tags 		Reviews
// @Accept 		json
// @Produce 	json
// @Param 		request body CreateReviewRequest true "Тело запроса"
// @Success 	201 {object} CreateReviewResponse 	"Успешно созданный отзыв"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/reviews [post]
func (handler *ReviewHandler) CreateReview(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	var req CreateReviewRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode the request")
		return
	}

	reviewDomain := domain.NewReviewUninitialized(req.Name, req.Email, req.Description, req.Rating)
	if err := reviewDomain.Validate(); err != nil {
		responseHandler.ErrorResponse(err, "invalid review")
		return
	}

	created, err := handler.reviewService.CreateReview(ctx, reviewDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create review")
		return
	}

	response := CreateReviewResponse(convertReviewDtoFromDomain(created))
	responseHandler.JsonResponse(response, http.StatusCreated)
}
