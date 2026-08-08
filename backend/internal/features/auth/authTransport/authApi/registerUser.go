package authApi

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=64" example:"Никита"`
	Email    string `json:"email" validate:"required,email" example:"bgeheh@bk.ru"`
	Password string `json:"password" validate:"required,min=8" example:"12345678"`
}

type RegisterUserResponse UserDTOResponse

// RegisterUser 	godoc
// @Summary 		Регистрация пользователя
// @Description 	Регистрация пользователя в систему
// @Tags 			Auth
// @Accept 			json
// @Produce 		json
// @Param 			request body RegisterUserRequest true "Тело запроса"
// @Success 		201 {object} RegisterUserResponse "Успешно регистрации пользователя пользователь"
// @Failure 		400 {object} response.ErrorResponse "Bad request"
// @Failure 		500 {object} response.ErrorResponse "Internal server error"
// @Router 			/auth/register [post]
func (c *AuthController) Register(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	var request RegisterUserRequest
	if err := requests.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domainFromDto(request)
	userDomain, err := c.authService.RegisterUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	err = c.issueAuthCookies(rw, userDomain.Id, domain.RoleUser)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create jwt")
		return
	}

	response := RegisterUserResponse(convertUserDTOFromDomain(userDomain))
	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domainFromDto(dto RegisterUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.Email, dto.Password)
}
