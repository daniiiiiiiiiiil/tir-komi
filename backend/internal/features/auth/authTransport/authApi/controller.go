package authApi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/cookies"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/jwt"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type AuthController struct {
	authService AuthService
	jwtProvider JwtProvider
}

type AuthService interface {
	RegisterUser(ctx context.Context, user domain.User) (domain.User, error)
	LoginUser(ctx context.Context, credentials domain.Credentials) (domain.User, error)
}

type JwtProvider interface {
	GenerateAccessToken(userID int, role domain.Role) (string, error)
	GenerateRefreshToken(userID int, role domain.Role) (string, error)
	ParseToken(tokenString string) (*jwt.Claims, error)
	ParseAccessToken(tokenString string) (*jwt.Claims, error)
	ParseRefreshToken(tokenString string) (*jwt.Claims, error)
	GetTokenFromCookie(r *http.Request, name string) (string, error)
}

func NewAuthController(authService AuthService, jwtProvider JwtProvider) *AuthController {
	return &AuthController{
		authService: authService,
		jwtProvider: jwtProvider,
	}
}

func (c *AuthController) Routers() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: c.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/register",
			Handler: c.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/logout",
			Handler: c.Logout,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/refresh",
			Handler: c.Refresh,
		},
	}
}

func (c *AuthController) issueAuthCookies(rw http.ResponseWriter, userId int, role domain.Role) error {
	accessToken, err := c.jwtProvider.GenerateAccessToken(userId, role)
	if err != nil {
		return fmt.Errorf("failed to generate access token %w", err)
	}

	refreshToken, err := c.jwtProvider.GenerateRefreshToken(userId, role)
	if err != nil {
		return fmt.Errorf("failed to generate refresh token, %w", err)
	}

	cookies.SetAuthCookies(
		rw,
		accessToken,
		refreshToken,
		15*time.Minute,
		30*24*time.Hour,
		cookies.Options{
			Secure:   false,
			Domain:   "localhost",
			SameSite: http.SameSiteLaxMode,
		},
	)
	return nil
}
