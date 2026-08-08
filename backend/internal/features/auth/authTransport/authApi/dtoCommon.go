package authApi

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type UserDTOResponse struct {
	Id       int         `json:"id"`
	Email    string      `json:"email"`
	Role     domain.Role `json:"role"`
	CreateAt time.Time   `json:"create_at"`
}

func convertUserDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		Id:       user.Id,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: user.CreateAt,
	}
}

func convertUserDTOsFromDomains(user []domain.User) []UserDTOResponse {
	userDTOs := make([]UserDTOResponse, len(user))
	for i, userDTO := range user {
		userDTOs[i] = convertUserDTOFromDomain(userDTO)
	}
	return userDTOs
}
