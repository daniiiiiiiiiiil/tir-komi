package http

import "github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"

type UpdateReviewRequest struct {
	Name        domain.Nullable[string] `json:"name"`
	Email       domain.Nullable[string] `json:"email"`
	Description domain.Nullable[string] `json:"description"`
	Rating      domain.Nullable[int]    `json:"rating"`
}
