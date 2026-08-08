package authRepository

import (
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

type AuthRepository struct {
	pool postgres.Pool
}

func NewAuthRepository(pool postgres.Pool) *AuthRepository {
	return &AuthRepository{pool: pool}
}

func (r *AuthRepository) convertEntityToDomain(userEntity UserEntity) domain.User {
	return domain.NewUser(
		userEntity.Id, userEntity.Email, userEntity.Password,
		userEntity.Role, userEntity.CreateAt,
	)
}
