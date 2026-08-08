package http

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
)

type MaterialHandler struct {
	materialService MaterialService
	fileStorage     FileStorage
}

type MaterialService interface {
	CreateMethodologicalMaterial(ctx context.Context, material domain.MethodologicalMaterial) (domain.MethodologicalMaterial, error)
	DeleteMethodologicalMaterial(ctx context.Context, id int) error
	GetMethodologicalMaterial(ctx context.Context, id int) (domain.MethodologicalMaterial, error)
	GetMethodologicalMaterials(ctx context.Context, limit, offset int) ([]domain.MethodologicalMaterial, int, error)
	UpdateMethodologicalMaterial(ctx context.Context, id int, patch domain.MethodologicalMaterialPatch) (domain.MethodologicalMaterial, error)
}

func NewMaterialHandler(materialService MaterialService, fileStorage FileStorage) *MaterialHandler {
	return &MaterialHandler{
		materialService: materialService,
		fileStorage:     fileStorage,
	}
}

type FileStorage interface {
	Save(ctx context.Context, data []byte) (string, error)
}

func (handler *MaterialHandler) Routers() []server.Route {
	return []server.Route{
		{
			Method:  "POST",
			Path:    "/materials",
			Handler: handler.CreateMaterial,
		},
		{
			Method:  "GET",
			Path:    "/materials",
			Handler: handler.GetMaterials,
		},
		{
			Method:  "GET",
			Path:    "/materials/{id}",
			Handler: handler.GetMaterial,
		},
		{
			Method:  "DELETE",
			Path:    "/materials/{id}",
			Handler: handler.DeleteMaterial,
		},
		{
			Method:  "PATCH",
			Path:    "/materials/{id}",
			Handler: handler.UpdateMaterial,
		},
	}
}
