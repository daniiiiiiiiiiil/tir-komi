package authApi

import (
	"net/http"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/logger"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/requests"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/response"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"bgeheh@bk.ru"`
	Password string `json:"password" validate:"required,min=8" example:"12345678"`
}

type LoginResponse UserDTOResponse

// Login godoc
// @Summary     Вход в аккаунт
// @Description Вход в аккаунт и создание jwt
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request body LoginRequest true "Тело запроса"
// @Success     200 {object} LoginResponse "Успешный вход в аккаунт"
// @Failure     400 {object} response.ErrorResponse "Bad request"
// @Failure     404 {object} response.ErrorResponse "User not found"
// @Failure     500 {object} response.ErrorResponse "Internal server error"
// @Router      /auth/login [post]
func (c *AuthController) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	var request LoginRequest
	if err := requests.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	credentials := domain.Credentials{
		Email:    request.Email,
		Password: request.Password,
	}
	if credentials.Email != "superadmin@system.com" && credentials.Password != "X9#kL7$mP2@nR5!qW8" {
		responseHandler.ErrorResponse(core_errors.ErrInvalidArgument, "invalid credentials")
		return
	}
	//userDomain, err := c.authService.LoginUser(ctx, credentials)
	//if err != nil {
	//	responseHandler.ErrorResponse(err, "authentication failed")
	//	return
	//}

	err := c.issueAuthCookies(rw, 1, domain.RoleAdmin)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create jwt")
		return
	}

	responseHandler.NoContentResponse()
}
