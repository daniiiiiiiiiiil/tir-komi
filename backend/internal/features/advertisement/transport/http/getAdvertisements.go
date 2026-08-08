package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

const (
	defaultLimit  = 20
	defaultOffset = 0
)

type AdvertisementsResponse struct {
	Items []AdvertisementDto `json:"items"`
	Total int                `json:"total"`
}

// GetAdvertisements 	godoc
// @Summary 	Получение объявлений
// @Description Получение списка объявлений с паггинацией
// @Tags 		Advertisements
// @Produce 	json
// @Param 		limit query int false 				"Размер страницы с объявлениями"
// @Param 		offset query int false 			"Смещение страницы с объявлениями"
// @Success 	200 {object} AdvertisementsResponse 	"Успешное получение списка объявлений"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/advertisements [get]
func (handler *AdvertisementHandler) GetAdvertisements(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset query params")
		return
	}

	ads, total, err := handler.advertisementService.GetAdvertisements(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get advertisements")
		return
	}

	response := AdvertisementsResponse{
		Items: convertAdvertisementDtosFromDomains(ads),
		Total: total,
	}
	responseHandler.JsonResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (int, int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	limit := defaultLimit
	if limitPtr, err := requests.GetIntQueryParams(r, limitQueryParamKey); err != nil {
		return 0, 0, err
	} else if limitPtr != nil {
		limit = *limitPtr
	}

	offset := defaultOffset
	if offsetPtr, err := requests.GetIntQueryParams(r, offsetQueryParamKey); err != nil {
		return 0, 0, err
	} else if offsetPtr != nil {
		offset = *offsetPtr
	}

	return limit, offset, nil
}
