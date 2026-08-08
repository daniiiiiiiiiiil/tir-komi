package authService

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/domain"
)

type AuthService struct {
	authRepository AuthRepository
	hasher         PasswordHasher
}

func NewAuthService(repo AuthRepository, hasher PasswordHasher) *AuthService {
	return &AuthService{
		authRepository: repo,
		hasher:         hasher,
	}
}

//go:generate go run github.com/golang/mock/mockgen@v1.6.0 -source=service.go -destination=mocks/mock_auth_service.go -package=mocks
type AuthRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(hashedPassword, plainPassword string) error
}
