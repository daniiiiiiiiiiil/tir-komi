package postgres_post

import "github.com/daniiiiiiiiiiil/tir-komi/internal/core/repository/pool/postgres"

type PostRepository struct {
	pool postgres.Pool
}

func NewPostRepository(p postgres.Pool) *PostRepository {
	return &PostRepository{pool: p}
}
