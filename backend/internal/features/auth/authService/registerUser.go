package authService

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/domain"
)

func (s *AuthService) RegisterUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("invalid user: %w", err)
	}

	hashed, err := s.hasher.Hash(user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashed

	created, err := s.authRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return created, nil
}
