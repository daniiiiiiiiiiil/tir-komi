package http

import (
	"fmt"
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type UpdateAdvertisementRequest struct {
	Title       domain.Nullable[string] `json:"title"`
	Description domain.Nullable[string] `json:"description"`
	Image       domain.Nullable[[]byte] `json:"image"`
	Pdf         domain.Nullable[[]byte] `json:"pdf"`
	Url         domain.Nullable[string] `json:"url"`
}

func (r *UpdateAdvertisementRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("Title can`t be NULL")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 200 {
			return fmt.Errorf("Title length must be between 1 and 200")
		}
	}

	if r.Url.Set {
		if r.Url.Value != nil {
			urlLen := len([]rune(*r.Url.Value))
			if urlLen < 1 || urlLen > 250 {
				return fmt.Errorf("Url length must be between 1 and 250")
			}
		}
	}

	return nil
}

type UpdateAdvertisementResponse AdvertisementDto

// UpdateAdvertisement 	godoc
// @Summary     Изменение объявления
// @Description Изменение информации об объявлении
// @Description ### Логика обновления полей:
// @Description 1.**Поле не переданно** значение в бд не меняется
// @Description 2.**Передан null** удаление поля из бд (для title недопустимо)
// @Description 3.**Передано значение** обновление в бд
// @Tags 		Advertisements
// @Accept 		json
// @Produce 	json
// @Param 		id path int true 					"ID объявления"
// @Param 		request body UpdateAdvertisementRequest true "Тело запроса"
// @Success 	200 {object} UpdateAdvertisementResponse 	"Успешно измененное объявление"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Advertisement not found"
// @Failure 	409 {object} response.ErrorResponse "Conflict"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/advertisements/{id} [patch]
func (handler *AdvertisementHandler) UpdateAdvertisement(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	adId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get advertisementId path value")
		return
	}

	var req UpdateAdvertisementRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request body")
		return
	}

	patch := domain.NewAdvertisementPatch(
		req.Title,
		req.Description,
		req.Image,
		req.Pdf,
		req.Url,
	)

	updated, err := handler.advertisementService.UpdateAdvertisement(ctx, adId, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update advertisement")
		return
	}

	response := UpdateAdvertisementResponse(convertAdvertisementDtoFromDomain(updated))
	responseHandler.JsonResponse(response, http.StatusOK)
}
