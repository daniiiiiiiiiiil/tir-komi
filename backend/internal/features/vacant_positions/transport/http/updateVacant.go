package http

import "github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"

type UpdateVacantRequest struct {
	Title       domain.Nullable[string] `json:"title"`
	Description domain.Nullable[string] `json:"description"`
}
