package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type UpdateVacantRequest struct {
	Title       domain.Nullable[string] `json:"title"`
	Description domain.Nullable[string] `json:"description"`
}

func (r *UpdateVacantRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("Title can`t be NULL")
		}

		titleLen := len([]rune(*r.Title.Value))
		if titleLen < 1 || titleLen > 200 {
			return fmt.Errorf("Title length must be between 1 and 200")
		}
	}

	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLen := len([]rune(*r.Description.Value))
			if descriptionLen < 1 || descriptionLen > 1000 {
				return fmt.Errorf("Description length must be between 1 and 1000")
			}
		}
	}

	return nil
}

type UpdateVacantResponse VacantDto

// UpdateVacant 	godoc
// @Summary     Изменение вакансии
// @Description Изменение информации о вакансии
// @Description ### Логика обновления полей:
// @Description 1.**Поле не переданно** значение в бд не меняется
// @Description 2.**Передан null** удаление поля из бд (для title недопустимо)
// @Description 3.**Передано значение** обновление в бд
// @Tags 		Vacancies
// @Accept 		json
// @Produce 	json
// @Param 		id path int true 					"ID вакансии"
// @Param 		request body UpdateVacantRequest true "Тело запроса"
// @Success 	200 {object} UpdateVacantResponse 	"Успешно измененная вакансия"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Vacant position not found"
// @Failure 	409 {object} response.ErrorResponse "Conflict"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/vacancies/{id} [patch]
func (handler *VacantHandler) UpdateVacant(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	vacantId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get vacantId path value")
		return
	}

	var req UpdateVacantRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request body")
		return
	}

	patch := domain.NewVacantPositionPatch(req.Title, req.Description, domain.Nullable[time.Time]{})

	updated, err := handler.vacantService.UpdateVacantPosition(ctx, vacantId, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update vacant position")
		return
	}

	response := UpdateVacantResponse(convertVacantDtoFromDomain(updated))
	responseHandler.JsonResponse(response, http.StatusOK)
}
