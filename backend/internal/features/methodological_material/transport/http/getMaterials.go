package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type MaterialsResponse struct {
	Items []MaterialDto `json:"items"`
	Total int           `json:"total"`
}

const (
	defaultLimit  = 20
	defaultOffset = 0
)

// GetMaterials 	godoc
// @Summary 	Получение методических материалов
// @Description Получение списка методических материалов с паггинацией
// @Tags 		Materials
// @Produce 	json
// @Param 		limit query int false 				"Размер страницы с материалами"
// @Param 		offset query int false 			"Смещение страницы с материалами"
// @Success 	200 {object} MaterialsResponse 		"Успешное получение списка материалов"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/materials [get]
func (handler *MaterialHandler) GetMaterials(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offset query params")
		return
	}

	materials, total, err := handler.materialService.GetMethodologicalMaterials(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get materials")
		return
	}

	response := MaterialsResponse{
		Items: convertMaterialDtosFromDomains(materials),
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
