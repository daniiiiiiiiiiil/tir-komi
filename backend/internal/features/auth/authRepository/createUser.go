package authRepository

import (
	"context"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

func (r *AuthRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO users (email,password,role,created_at)
	VALUES ($1,$2, $3, $4)
	RETURNING id, email, password, role, created_at;
`
	exec := r.pool.ExecutorFromContext(ctx)
	row := exec.QueryRow(ctx, query, user.Email, user.Password, user.Role, user.CreateAt)

	var userEntity UserEntity
	err := row.Scan(&userEntity.Id, &userEntity.Email, &userEntity.Password, &userEntity.Role, &userEntity.CreateAt)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan user error: %w", err)
	}

	return r.convertEntityToDomain(userEntity), nil
}
