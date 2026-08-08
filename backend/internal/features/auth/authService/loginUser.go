package authService

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/domain"
)

func (s *AuthService) LoginUser(ctx context.Context, credentials domain.Credentials) (domain.User, error) {
	if err := credentials.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("invalid credentials: %w", err)
	}

	user, err := s.authRepository.GetUserByEmail(ctx, credentials.Email)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user by email %w", err)
	}

	if err := s.hasher.Verify(user.Password, credentials.Password); err != nil {
		return domain.User{}, fmt.Errorf("invalid credentials: %w", err)
	}

	return user, nil
}
