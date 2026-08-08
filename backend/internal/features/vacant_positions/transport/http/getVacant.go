package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type VacantResponse VacantDto

// GetVacant 	godoc
// @Summary 	Получение вакансии
// @Description Получение вакансии по id
// @Tags 		Vacancies
// @Produce 	json
// @Param 		id path int true 					"ID вакансии"
// @Success 	200 {object} VacantResponse 		"Успешно найденная вакансия по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Vacant position not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/vacancies/{id} [get]
func (handler *VacantHandler) GetVacant(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	vacantId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get vacantId path value")
		return
	}

	vacantDomain, err := handler.vacantService.GetVacantPosition(ctx, vacantId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get vacant position")
		return
	}

	response := VacantResponse(convertVacantDtoFromDomain(vacantDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
