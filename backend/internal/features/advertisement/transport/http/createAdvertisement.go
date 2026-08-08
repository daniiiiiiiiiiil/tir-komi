package http

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type CreateAdvertisementRequest struct {
	Title       string  `json:"title" validate:"required,min=1,max=200" example:"Скидка 20% на всё"`
	Description string  `json:"description" validate:"omitempty,max=1000" example:"Акция действует до конца месяца"`
	Image       *[]byte `json:"image,omitempty"`
	Pdf         *[]byte `json:"pdf,omitempty"`
	Url         *string `json:"url" validate:"omitempty,max=250"`
}

type CreateAdvertisementResponse AdvertisementDto

// CreateAdvertisement 	godoc
// @Summary 	Создание объявления
// @Description Создать новое объявление в системе
// @Tags 		Advertisements
// @Accept 		json
// @Produce 	json
// @Param 		request body CreateAdvertisementRequest true "Тело запроса"
// @Success 	201 {object} CreateAdvertisementResponse 	"Успешно созданное объявление"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/advertisements [post]
func (handler *AdvertisementHandler) CreateAdvertisement(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	var req CreateAdvertisementRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode the request")
		return
	}

	var imagePath, pdfPath *string
	if req.Image != nil {
		path, err := handler.fileStorage.Save(ctx, *req.Image)
		if err != nil {
			responseHandler.ErrorResponse(err, "failed to save image")
			return
		}
		imagePath = &path
	}
	if req.Pdf != nil {
		path, err := handler.fileStorage.Save(ctx, *req.Pdf)
		if err != nil {
			responseHandler.ErrorResponse(err, "failed to save pdf")
			return
		}
		pdfPath = &path
	}

	description := &req.Description
	adDomain := domain.NewAdvertisementUninitialized(req.Title, description, imagePath, pdfPath, req.Url)
	if err := adDomain.Validate(); err != nil {
		responseHandler.ErrorResponse(err, "invalid advertisement")
		return
	}

	created, err := handler.advertisementService.CreateAdvertisement(ctx, adDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create advertisement")
		return
	}

	response := CreateAdvertisementResponse(convertAdvertisementDtoFromDomain(created))
	responseHandler.JsonResponse(response, http.StatusCreated)
}
