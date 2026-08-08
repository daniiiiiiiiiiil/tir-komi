package http

import (
	"net/http"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type CreateMaterialRequest struct {
	Title       string  `json:"title" validate:"required,min=1,max=200" example:"Методичка по Go"`
	Description string  `json:"description" validate:"omitempty,max=1000" example:"Базовые принципы разработки на Go"`
	Pdf         *[]byte `json:"pdf"`
}

type CreateMaterialResponse MaterialDto

// CreateMaterial 	godoc
// @Summary 	Создание методического материала
// @Description Создать новый методический материал в системе
// @Tags 		Materials
// @Accept 		json
// @Produce 	json
// @Param 		request body CreateMaterialRequest true "Тело запроса"
// @Success 	201 {object} CreateMaterialResponse 	"Успешно созданный материал"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/materials [post]
func (handler *MaterialHandler) CreateMaterial(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	var req CreateMaterialRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode the request")
		return
	}

	var pdfPath *string
	if req.Pdf != nil {
		path, err := handler.fileStorage.Save(ctx, *req.Pdf)
		if err != nil {
			responseHandler.ErrorResponse(err, "failed to save pdf")
			return
		}
		pdfPath = &path
	}

	description := &req.Description
	date := time.Now()
	materialDomain := domain.NewMethodologicalMaterialUninitialized(req.Title, description, &date, pdfPath)
	if err := materialDomain.Validate(); err != nil {
		responseHandler.ErrorResponse(err, "invalid material")
		return
	}

	created, err := handler.materialService.CreateMethodologicalMaterial(ctx, materialDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create material")
		return
	}

	response := CreateMaterialResponse(convertMaterialDtoFromDomain(created))
	responseHandler.JsonResponse(response, http.StatusCreated)
}
