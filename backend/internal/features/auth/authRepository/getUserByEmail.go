package authRepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"
)

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, email, password, role, created_at FROM users
    WHERE email = $1
`

	exec := r.pool.ExecutorFromContext(ctx)
	row := exec.QueryRow(ctx, query, email)

	var userEntity UserEntity
	err := row.Scan(&userEntity.Id, &userEntity.Email, &userEntity.Password, &userEntity.Role, &userEntity.CreateAt)

	if err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with email=%s: %w", email, core_errors.ErrNotFound)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	return r.convertEntityToDomain(userEntity), nil
}
