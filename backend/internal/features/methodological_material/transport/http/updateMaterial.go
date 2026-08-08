package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type UpdateMaterialRequest struct {
	Title       domain.Nullable[string] `json:"title"`
	Description domain.Nullable[string] `json:"description"`
	Pdf         domain.Nullable[[]byte] `json:"pdf"`
}

func (r *UpdateMaterialRequest) Validate() error {
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

type UpdateMaterialResponse MaterialDto

// UpdateMaterial 	godoc
// @Summary     Изменение методического материала
// @Description Изменение информации о методическом материале
// @Description ### Логика обновления полей:
// @Description 1.**Поле не переданно** значение в бд не меняется
// @Description 2.**Передан null** удаление поля из бд (для title недопустимо)
// @Description 3.**Передано значение** обновление в бд
// @Tags 		Materials
// @Accept 		json
// @Produce 	json
// @Param 		id path int true 					"ID материала"
// @Param 		request body UpdateMaterialRequest true "Тело запроса"
// @Success 	200 {object} UpdateMaterialResponse 	"Успешно измененный материал"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	404 {object} response.ErrorResponse "Material not found"
// @Failure 	409 {object} response.ErrorResponse "Conflict"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/materials/{id} [patch]
func (handler *MaterialHandler) UpdateMaterial(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	materialId, err := requests.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get materialId path value")
		return
	}

	var req UpdateMaterialRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate request body")
		return
	}

	pdfPatch, err := handler.savePdfPatch(ctx, req.Pdf)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to save pdf")
		return
	}

	patch := domain.NewMethodologicalMaterialPatch(req.Title, req.Description, domain.Nullable[time.Time]{}, pdfPatch)

	updated, err := handler.materialService.UpdateMethodologicalMaterial(ctx, materialId, patch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update material")
		return
	}

	response := UpdateMaterialResponse(convertMaterialDtoFromDomain(updated))
	responseHandler.JsonResponse(response, http.StatusOK)
}

func (handler *MaterialHandler) savePdfPatch(ctx context.Context, patch domain.Nullable[[]byte]) (domain.Nullable[string], error) {
	if !patch.Set {
		return domain.Nullable[string]{}, nil
	}
	if patch.Value == nil {
		return domain.Nullable[string]{Set: true, Value: nil}, nil
	}

	path, err := handler.fileStorage.Save(ctx, *patch.Value)
	if err != nil {
		return domain.Nullable[string]{}, err
	}
	return domain.Nullable[string]{Set: true, Value: &path}, nil
}
