package postgres

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type MediaModel struct {
	ID          int
	Type        string
	Title       *string
	Description *string
	FileData    []byte
	FileName    string
	MimeType    string
	FileSize    int64
	SortOrder   int
	CreatedAt   time.Time
}

func mediaDomainFromModel(model MediaModel) domain.Media {
	return domain.NewMedia(
		model.ID,
		domain.MediaType(model.Type),
		model.Title,
		model.Description,
		model.FileData,
		model.FileName,
		model.MimeType,
		model.FileSize,
		model.SortOrder,
		model.CreatedAt,
	)
}

func mediaModelFromDomain(m domain.Media) MediaModel {
	return MediaModel{
		ID:          m.ID,
		Type:        string(m.Type),
		Title:       m.Title,
		Description: m.Description,
		FileData:    m.FileData,
		FileName:    m.FileName,
		MimeType:    m.MimeType,
		FileSize:    m.FileSize,
		SortOrder:   m.SortOrder,
		CreatedAt:   m.CreatedAt,
	}
}

func mediaModelsFromDomains(items []domain.Media) []MediaModel {
	if len(items) == 0 {
		return nil
	}
	models := make([]MediaModel, len(items))
	for i, m := range items {
		models[i] = mediaModelFromDomain(m)
	}
	return models
}

func mediaDomainsFromModels(models []MediaModel) []domain.Media {
	if len(models) == 0 {
		return nil
	}
	items := make([]domain.Media, len(models))
	for i, model := range models {
		items[i] = mediaDomainFromModel(model)
	}
	return items
}
