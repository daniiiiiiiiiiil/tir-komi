package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

// DeleteVacant 	godoc
// @Summary 	Удаление вакансии
// @Description Удалить вакансию в системе
// @Tags 		Vacancies
// @Produce 	json
// @Param 		id path int true 					"ID вакансии"
// @Success 	204 								"Успешно удаленная вакансия по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Vacant position not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/vacancies/{id} [delete]
func (handler *VacantHandler) DeleteVacant(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	vacantId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get vacantId path value")
		return
	}

	if err := handler.vacantService.DeleteVacantPosition(ctx, vacantId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete vacant position")
		return
	}
	responseHandler.NoContentResponse()
}
