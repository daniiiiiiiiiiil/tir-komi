package http

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

var emailRegexp = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

type UpdateReviewRequest struct {
	Name        domain.Nullable[string] `json:"name"`
	Email       domain.Nullable[string] `json:"email"`
	Description domain.Nullable[string] `json:"description"`
	Rating      domain.Nullable[int]    `json:"rating"`
}

func (r *UpdateReviewRequest) Validate() error {
	if r.Name.Set {
		if r.Name.Value == nil {
			return fmt.Errorf("Name can`t be NULL")
		}

		nameLen := len([]rune(*r.Name.Value))
		if nameLen < 1 || nameLen > 100 {
			return fmt.Errorf("Name length must be between 1 and 100")
		}
	}

	if r.Email.Set {
		if r.Email.Value == nil {
			return fmt.Errorf("Email can`t be NULL")
		}
		if !emailRegexp.MatchString(*r.Email.Value) {
			return fmt.Errorf("Email is invalid")
		}
	}

	if r.Description.Set {
		if r.Description.Value == nil {
			return fmt.Errorf("Description can`t be NULL")
		}

		descriptionLen := len([]rune(*r.Description.Value))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf("Description length must be between 1 and 1000")
		}
	}

	if r.Rating.Set {
		if r.Rating.Value == nil {
			return fmt.Errorf("Rating can`t be NULL")
		}
		if *r.Rating.Value < 1 || *r.Rating.Value > 5 {
			return fmt.Errorf("Rating must be between 1 and 5")
		}
	}

	return nil
}

type UpdateReviewResponse ReviewDto

// UpdateReview 	godoc
// @Summary     Изменение отзыва
// @Description Изменение информации об отзыве
// @Description ### Логика обновления полей:
// @Description 1.**Поле не переданно** значение в бд не меняется
// @Description 2.**Передан null** недопустимо ни для одного поля
// @Description 3.**Передано значение** обновление в бд
// @Tags 		Reviews
// @Accept 		json
// @Produce 	json
// @Param 		id path int true 					"ID отзыва"
// @Param 		request body UpdateReviewRequest true "Тело запроса"
// @Success 	200 {object} UpdateReviewResponse 	"Успешно измененный отзыв"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Review not found"
// @Failure 	409 {object} response.ErrorResponse "Conflict"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/reviews/{id} [patch]
func (handler *ReviewHandler) UpdateReview(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	reviewId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get reviewId path value")
		return
	}

	var req UpdateReviewRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request body")
		return
	}

	patch := domain.NewReviewPatch(req.Name, req.Email, req.Description, req.Rating)

	updated, err := handler.reviewService.UpdateReview(ctx, reviewId, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update review")
		return
	}

	response := UpdateReviewResponse(convertReviewDtoFromDomain(updated))
	responseHandler.JsonResponse(response, http.StatusOK)
}
