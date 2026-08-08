package http

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type AdvertisementHandler struct {
	advertisementService AdvertisementService
	fileStorage          FileStorage
}

type AdvertisementService interface {
	CreateAdvertisement(ctx context.Context, ad domain.Advertisement) (domain.Advertisement, error)
	DeleteAdvertisement(ctx context.Context, id int) error
	GetAdvertisement(ctx context.Context, id int) (domain.Advertisement, error)
	GetAdvertisements(ctx context.Context, limit, offset int) ([]domain.Advertisement, int, error)
	UpdateAdvertisement(ctx context.Context, id int, patch domain.AdvertisementPatch) (domain.Advertisement, error)
}

type FileStorage interface {
	Save(ctx context.Context, data []byte) (string, error)
}

func NewAdvertisementHandler(advertisementService AdvertisementService, fileStorage FileStorage) *AdvertisementHandler {
	return &AdvertisementHandler{
		advertisementService: advertisementService,
		fileStorage:          fileStorage,
	}
}

func (handler *AdvertisementHandler) Routers() []server.Route {
	return []server.Route{
		{
			Method:  "POST",
			Path:    "/advertisements",
			Handler: handler.CreateAdvertisement,
		},
		{
			Method:  "GET",
			Path:    "/advertisements",
			Handler: handler.GetAdvertisements,
		},
		{
			Method:  "GET",
			Path:    "/advertisements/{id}",
			Handler: handler.GetAdvertisement,
		},
		{
			Method:  "DELETE",
			Path:    "/advertisements/{id}",
			Handler: handler.DeleteAdvertisement,
		},
		{
			Method:  "PATCH",
			Path:    "/advertisements/{id}",
			Handler: handler.UpdateAdvertisement,
		},
	}
}
