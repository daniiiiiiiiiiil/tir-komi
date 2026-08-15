package postgres

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type PostModel struct {
	ID          int
	Title       string
	Description *string
	Pdf         []byte
	Image       []byte
	Date        *time.Time
	Type        string
}

func postDomainFromModel(model PostModel) domain.Post {
	return domain.NewPost(
		model.ID,
		model.Title,
		model.Description,
		model.Pdf,
		model.Image,
		model.Date,
		domain.PostType(model.Type),
	)
}

func postModelsFromDomains(posts []domain.Post) []PostModel {
	models := make([]PostModel, len(posts))
	for i, post := range posts {
		models[i] = PostModel{
			ID:          post.ID,
			Title:       post.Title,
			Description: post.Description,
			Pdf:         post.Pdf,
			Image:       post.Image,
			Date:        post.Date,
			Type:        string(post.Type),
		}
	}
	return models
}
