package domain

import (
	"fmt"
	"regexp"
	"time"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/backend/internal/core/errors"
)

type User struct {
	Id       int
	Email    string
	Password string
	Role     Role
	CreateAt time.Time
}

func NewUser(id int, email string, password string, role Role, createAt time.Time) User {
	return User{
		Id:       id,
		Email:    email,
		Password: password,
		Role:     role,
		CreateAt: createAt,
	}
}

func NewUserUninitialized(email string, password string) User {
	return User{
		Id:       UninitializedID,
		Email:    email,
		Password: password,
		Role:     RoleUser,
		CreateAt: time.Now(),
	}
}

func (u *User) Validate() error {

	emailLen := len([]rune(u.Email))
	if emailLen > 128 {
		return fmt.Errorf("Invalid user email length %d,%w", emailLen, core_errors.ErrInvalidArgument)
	}
	re := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	if !re.MatchString(u.Email) {
		return fmt.Errorf("Invalid user email format %s,%w", u.Email, core_errors.ErrInvalidArgument)
	}

	passwordLen := len([]rune(u.Password))
	if passwordLen <= 7 {
		return fmt.Errorf("Invalid user password length %d,%w", passwordLen, core_errors.ErrInvalidArgument)
	}
	return nil
}

type Credentials struct {
	Email    string
	Password string
}

func NewCredentials(email string, password string) Credentials {
	return Credentials{
		Email:    email,
		Password: password,
	}
}

func (c *Credentials) Validate() error {
	emailLen := len([]rune(c.Email))

	if emailLen > 128 {
		return fmt.Errorf("Invalid user email length %d,%w", emailLen, core_errors.ErrInvalidArgument)
	}
	re := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
	if !re.MatchString(c.Email) {
		return fmt.Errorf("Invalid user email format %s,%w", c.Email, core_errors.ErrInvalidArgument)
	}

	passwordLen := len([]rune(c.Password))
	if passwordLen <= 7 {
		return fmt.Errorf("Invalid user password length %d,%w", passwordLen, core_errors.ErrInvalidArgument)
	}

	return nil
}
