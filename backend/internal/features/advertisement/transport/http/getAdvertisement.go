package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type AdvertisementResponse AdvertisementDto

// GetAdvertisement 	godoc
// @Summary 	Получение объявления
// @Description Получение объявления по id
// @Tags 		Advertisements
// @Produce 	json
// @Param 		id path int true 					"ID объявления"
// @Success 	200 {object} AdvertisementResponse 	"Успешно найденное объявление по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Advertisement not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/advertisements/{id} [get]
func (handler *AdvertisementHandler) GetAdvertisement(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	adId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get advertisementId path value")
		return
	}

	adDomain, err := handler.advertisementService.GetAdvertisement(ctx, adId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get advertisement")
		return
	}

	response := AdvertisementResponse(convertAdvertisementDtoFromDomain(adDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
