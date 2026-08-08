package authRepository

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type UserEntity struct {
	Id       int
	Email    string
	Password string
	Role     domain.Role
	CreateAt time.Time
}
