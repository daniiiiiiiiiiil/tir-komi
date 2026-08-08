package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

// DeleteAdvertisement 	godoc
// @Summary 	Удаление объявления
// @Description Удалить объявление в системе
// @Tags 		Advertisements
// @Produce 	json
// @Param 		id path int true 					"ID объявления"
// @Success 	204 								"Успешно удаленное объявление по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Advertisement not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/advertisements/{id} [delete]
func (handler *AdvertisementHandler) DeleteAdvertisement(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	adId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get advertisementId path value")
		return
	}

	if err := handler.advertisementService.DeleteAdvertisement(ctx, adId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete advertisement")
		return
	}
	responseHandler.NoContentResponse()
}
