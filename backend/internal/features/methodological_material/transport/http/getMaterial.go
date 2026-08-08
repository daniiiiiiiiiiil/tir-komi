package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type MaterialResponse MaterialDto

// GetMaterial 	godoc
// @Summary 	Получение методического материала
// @Description Получение методического материала по id
// @Tags 		Materials
// @Produce 	json
// @Param 		id path int true 					"ID материала"
// @Success 	200 {object} MaterialResponse 		"Успешно найденный материал по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Material not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/materials/{id} [get]
func (handler *MaterialHandler) GetMaterial(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	materialId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get materialId path value")
		return
	}

	materialDomain, err := handler.materialService.GetMethodologicalMaterial(ctx, materialId)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get material")
		return
	}

	response := MaterialResponse(convertMaterialDtoFromDomain(materialDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
