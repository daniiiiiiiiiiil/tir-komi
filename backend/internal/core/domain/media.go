package domain

import (
	"fmt"
	"time"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

type MediaType string

const (
	MediaTypePhoto MediaType = "photo"
	MediaTypeVideo MediaType = "video"
)

func (t MediaType) Valid() bool {
	switch t {
	case MediaTypePhoto, MediaTypeVideo:
		return true
	default:
		return false
	}
}

type Media struct {
	ID          int
	Type        MediaType
	Title       *string
	Description *string
	FileData    []byte
	FileName    string
	MimeType    string
	FileSize    int64
	SortOrder   int
	CreatedAt   time.Time
}

func NewMedia(id int, mediaType MediaType, title *string, description *string, fileData []byte, fileName string, mimeType string, fileSize int64, sortOrder int, createdAt time.Time) Media {
	return Media{
		ID:          id,
		Type:        mediaType,
		Title:       title,
		Description: description,
		FileData:    fileData,
		FileName:    fileName,
		MimeType:    mimeType,
		FileSize:    fileSize,
		SortOrder:   sortOrder,
		CreatedAt:   createdAt,
	}
}

func NewMediaUninitialized(mediaType MediaType, title *string, description *string, fileData []byte, fileName string, mimeType string, sortOrder int) Media {
	return Media{
		ID:          UninitializedID,
		Type:        mediaType,
		Title:       title,
		Description: description,
		FileData:    fileData,
		FileName:    fileName,
		MimeType:    mimeType,
		FileSize:    int64(len(fileData)),
		SortOrder:   sortOrder,
		CreatedAt:   time.Now(),
	}
}

func (m *Media) Validate() error {
	if !m.Type.Valid() {
		return fmt.Errorf("invalid media type: %q, %w", m.Type, core_errors.ErrInvalidArgument)
	}

	if m.Title != nil {
		titleLen := len([]rune(*m.Title))
		if titleLen < 1 || titleLen > 200 {
			return fmt.Errorf("title must be between 1 and 200 characters: %d, %w", titleLen, core_errors.ErrInvalidArgument)
		}
	}

	if m.Description != nil {
		descLen := len([]rune(*m.Description))
		if descLen > 1000 {
			return fmt.Errorf("description too long: %d, max 1000 characters, %w", descLen, core_errors.ErrInvalidArgument)
		}
	}

	if len(m.FileData) == 0 {
		return fmt.Errorf("file is required: %w", core_errors.ErrInvalidArgument)
	}

	fileNameLen := len([]rune(m.FileName))
	if fileNameLen < 1 || fileNameLen > 255 {
		return fmt.Errorf("file_name must be between 1 and 255 characters: %d, %w", fileNameLen, core_errors.ErrInvalidArgument)
	}

	if m.MimeType == "" {
		return fmt.Errorf("mime_type is required: %w", core_errors.ErrInvalidArgument)
	}

	if m.FileSize <= 0 {
		return fmt.Errorf("file_size must be greater than 0: %w", core_errors.ErrInvalidArgument)
	}

	switch m.Type {
	case MediaTypePhoto:
		if len(m.MimeType) >= 6 && m.MimeType[:6] != "image/" && m.MimeType != "application/octet-stream" {
			return fmt.Errorf("mime_type %q does not look like an image for photo media: %w", m.MimeType, core_errors.ErrInvalidArgument)
		}
	case MediaTypeVideo:
		if len(m.MimeType) >= 6 && m.MimeType[:6] != "video/" && m.MimeType != "application/octet-stream" {
			return fmt.Errorf("mime_type %q does not look like a video for video media: %w", m.MimeType, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

type MediaPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	FileData    Nullable[[]byte]
	FileName    Nullable[string]
	MimeType    Nullable[string]
	SortOrder   Nullable[int]
}

func NewMediaPatch(title Nullable[string], description Nullable[string], fileData Nullable[[]byte], fileName Nullable[string], mimeType Nullable[string], sortOrder Nullable[int]) MediaPatch {
	return MediaPatch{
		Title:       title,
		Description: description,
		FileData:    fileData,
		FileName:    fileName,
		MimeType:    mimeType,
		SortOrder:   sortOrder,
	}
}

func (p *MediaPatch) Validate() error {
	if p.SortOrder.Set && p.SortOrder.Value == nil {
		return fmt.Errorf("sort_order cannot be null: %w", core_errors.ErrInvalidArgument)
	}

	fileTouched := p.FileData.Set || p.FileName.Set || p.MimeType.Set
	if fileTouched {
		if !(p.FileData.Set && p.FileName.Set && p.MimeType.Set) {
			return fmt.Errorf("file_data, file_name and mime_type must be provided together: %w", core_errors.ErrInvalidArgument)
		}
		if p.FileData.Value == nil || p.FileName.Value == nil || p.MimeType.Value == nil {
			return fmt.Errorf("file_data, file_name and mime_type cannot be null: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

func (m *Media) ApplyPatch(patch MediaPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid patch: %w", err)
	}

	tmp := *m

	if patch.Title.Set {
		tmp.Title = patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.FileData.Set {
		tmp.FileData = *patch.FileData.Value
		tmp.FileName = *patch.FileName.Value
		tmp.MimeType = *patch.MimeType.Value
		tmp.FileSize = int64(len(tmp.FileData))
	}

	if patch.SortOrder.Set {
		tmp.SortOrder = *patch.SortOrder.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid media after patch: %w", err)
	}

	*m = tmp
	return nil
}
