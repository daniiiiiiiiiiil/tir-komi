package service

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type MediaRepository interface {
	CreateMedia(ctx context.Context, m domain.Media) (domain.Media, error)
	GetMedia(ctx context.Context, id int) (domain.Media, error)
	GetMedias(ctx context.Context, mediaType *domain.MediaType, limit, offset int) ([]domain.Media, int, error)
	UpdateMedia(ctx context.Context, id int, patch domain.MediaPatch) (domain.Media, error)
	DeleteMedia(ctx context.Context, id int) error
}

type MediaService struct {
	repo MediaRepository
}

func NewMediaService(repo MediaRepository) *MediaService {
	return &MediaService{
		repo: repo,
	}
}
