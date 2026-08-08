package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

// DeleteMaterial 	godoc
// @Summary 	Удаление методического материала
// @Description Удалить методический материал в системе
// @Tags 		Materials
// @Produce 	json
// @Param 		id path int true 					"ID материала"
// @Success 	204 								"Успешно удаленный материал по Id"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Material not found"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/materials/{id} [delete]
func (handler *MaterialHandler) DeleteMaterial(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	materialId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get materialId path value")
		return
	}

	if err := handler.materialService.DeleteMethodologicalMaterial(ctx, materialId); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete material")
		return
	}
	responseHandler.NoContentResponse()
}
