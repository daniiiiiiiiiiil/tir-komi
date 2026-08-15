package http

import (
	"context"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/transport/http/server"
	service2 "github.com/daniiiiiiiiiiil/tir-komi/internal/features/posts/service"
)

type PostHandler struct {
	postService PostService
}

type PostService interface {
	CreatePost(ctx context.Context, post domain.Post) (domain.Post, error)
	DeletePost(ctx context.Context, id int) error
	GetPost(ctx context.Context, id int) (domain.Post, error)
	GetPosts(ctx context.Context, limit, offset int) ([]domain.Post, error)
	UpdatePost(ctx context.Context, id int, patch domain.PostPatch) (domain.Post, error)
}

func NewPostHandler(postService *service2.PostService) *PostHandler {
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
			Path:    "/posts",
			Handler: handler.GetPosts,
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
