package http

import (
	"net/http"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type CreateVacantRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=200" example:"Go разработчик"`
	Description string `json:"description" validate:"omitempty,max=1000" example:"Ищем опытного Go разработчика в команду"`
}

type CreateVacantResponse VacantDto

// CreateVacant 	godoc
// @Summary 	Создание вакансии
// @Description Создать новую вакансию в системе
// @Tags 		Vacancies
// @Accept 		json
// @Produce 	json
// @Param 		request body CreateVacantRequest true "Тело запроса"
// @Success 	201 {object} CreateVacantResponse 	"Успешно созданная вакансия"
// @Failure 	400 {object} response.ErrorResponse "Bad request"
// @Failure 	500 {object} response.ErrorResponse "Internal server error"
// @Router 		/vacancies [post]
func (handler *VacantHandler) CreateVacant(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	var req CreateVacantRequest
	if err := requests.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode the request")
		return
	}

	description := &req.Description
	date := time.Now()
	vacantDomain := domain.NewVacantPositionUninitialized(req.Title, description, &date)
	if err := vacantDomain.Validate(); err != nil {
		responseHandler.ErrorResponse(err, "invalid vacant position")
		return
	}

	created, err := handler.vacantService.CreateVacantPosition(ctx, vacantDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create vacant position")
		return
	}

	response := CreateVacantResponse(convertVacantDtoFromDomain(created))
	responseHandler.JsonResponse(response, http.StatusCreated)
}
