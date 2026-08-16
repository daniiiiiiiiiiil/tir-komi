package http_post

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/features/posts/service_post"
)

type PostHandler struct {
	postService PostService
}

type PostService interface {
	CreatePost(ctx context.Context, post domain.Post) (domain.Post, error)
	DeletePost(ctx context.Context, id int) error
	GetPost(ctx context.Context, id int) (domain.Post, error)
	GetPostsByType(ctx context.Context, postType domain.PostType) ([]domain.Post, error)
	UpdatePost(ctx context.Context, id int, patch domain.PostPatch) (domain.Post, error)
}

func NewPostHandler(postService *service_post.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (handler *PostHandler) Routers() []server.Route {
	return []server.Route{
		{
			Method:  "POST",
			Path:    "/posts",
			Handler: handler.CreatePost,
		},
		{
			Method:  "GET",
			Path:    "/posts/{}",
			Handler: handler.GetPostsByType,
		},
		{
			Method:  "GET",
			Path:    "/posts/{id}",
			Handler: handler.GetPost,
		},
		{
			Method:  "DELETE",
			Path:    "/posts/{id}",
			Handler: handler.DeletePost,
		},
		{
			Method:  "PATCH",
			Path:    "/posts/{id}",
			Handler: handler.UpdatePost,
		},
	}
}
