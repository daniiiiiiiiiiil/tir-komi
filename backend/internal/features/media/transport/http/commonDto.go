package http

import (
	"fmt"
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type MediaDto struct {
	Id          int       `json:"id"`
	Type        string    `json:"type"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	FileName    string    `json:"file_name"`
	MimeType    string    `json:"mime_type"`
	FileSize    int64     `json:"file_size"`
	FileUrl     string    `json:"file_url"`
	SortOrder   int       `json:"sort_order"`
	CreateAt    time.Time `json:"create_at"`
}

func convertMediaDtoFromDomain(m domain.Media) MediaDto {
	title := ""
	if m.Title != nil {
		title = *m.Title
	}

	description := ""
	if m.Description != nil {
		description = *m.Description
	}

	return MediaDto{
		Id:          m.ID,
		Type:        string(m.Type),
		Title:       title,
		Description: description,
		FileName:    m.FileName,
		MimeType:    m.MimeType,
		FileSize:    m.FileSize,
		FileUrl:     fmt.Sprintf("/media/%d/file", m.ID),
		SortOrder:   m.SortOrder,
		CreateAt:    m.CreatedAt,
	}
}

func convertMediaDtosFromDomains(items []domain.Media) []MediaDto {
	dtos := make([]MediaDto, len(items))
	for i, m := range items {
		dtos[i] = convertMediaDtoFromDomain(m)
	}
	return dtos
}

func parseIntOrZero(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	return n, err
}
