package http

import "github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"

type UpdateMaterialRequest struct {
	Title       domain.Nullable[string] `json:"title"`
	Description domain.Nullable[string] `json:"description"`
	Pdf         domain.Nullable[[]byte] `json:"pdf"`
}
