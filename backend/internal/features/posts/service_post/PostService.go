package service_post

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post domain.Post) (domain.Post, error)
	GetPost(ctx context.Context, id int) (domain.Post, error)
	GetPostsByType(ctx context.Context, postType domain.PostType) ([]domain.Post, error)
	UpdatePost(ctx context.Context, id int, patch domain.PostPatch) (domain.Post, error)
	DeletePost(ctx context.Context, id int) error
}

type PostService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}
