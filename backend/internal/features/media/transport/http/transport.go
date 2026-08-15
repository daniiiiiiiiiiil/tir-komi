package http

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type MediaHandler struct {
	mediaService MediaService
}

type MediaService interface {
	CreateMedia(ctx context.Context, m domain.Media) (domain.Media, error)
	DeleteMedia(ctx context.Context, id int) error
	GetMedia(ctx context.Context, id int) (domain.Media, error)
	GetMedias(ctx context.Context, mediaType *domain.MediaType, limit, offset int) ([]domain.Media, int, error)
	UpdateMedia(ctx context.Context, id int, patch domain.MediaPatch) (domain.Media, error)
}

func NewMediaHandler(mediaService MediaService) *MediaHandler {
	return &MediaHandler{
		mediaService: mediaService,
	}
}

func (handler *MediaHandler) Routers() []server.Route {
	return []server.Route{
		{Method: "POST", Path: "/media", Handler: handler.CreateMedia},
		{Method: "GET", Path: "/media", Handler: handler.GetMedias},
		{Method: "GET", Path: "/media/{id}", Handler: handler.GetMedia},
		{Method: "GET", Path: "/media/{id}/file", Handler: handler.GetMediaFile},
		{Method: "DELETE", Path: "/media/{id}", Handler: handler.DeleteMedia},
		{Method: "PATCH", Path: "/media/{id}", Handler: handler.UpdateMedia},
	}
}
