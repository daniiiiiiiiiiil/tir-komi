package http

import "github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"

type UpdateAdvertisementRequest struct {
	Title domain.Nullable[string] `json:"title"`
	Image domain.Nullable[[]byte] `json:"image"`
	Pdf   domain.Nullable[[]byte] `json:"pdf"`
	Url   domain.Nullable[string] `json:"url"`
}
