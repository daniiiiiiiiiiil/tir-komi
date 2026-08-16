package http_post

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type PostDto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Pdf         []byte    `json:"pdf,omitempty"`
	Image       []byte    `json:"image,omitempty"`
	CreateAt    time.Time `json:"create_at"`
	Type        string    `json:"type"`
}

func convertPostDtoFromDomain(post domain.Post) PostDto {
	description := ""
	if post.Description != nil {
		description = *post.Description
	}

	createAt := time.Time{}
	if post.Date != nil {
		createAt = *post.Date
	}

	return PostDto{
		Id:          post.ID,
		Title:       post.Title,
		Description: description,
		Pdf:         post.Pdf,
		Image:       post.Image,
		CreateAt:    createAt,
		Type:        string(post.Type),
	}
}

func convertPostDtosFromDomains(posts []domain.Post) []PostDto {
	dtos := make([]PostDto, len(posts))
	for i, post := range posts {
		dtos[i] = convertPostDtoFromDomain(post)
	}
	return dtos
}
