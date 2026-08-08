package domain

import (
	"fmt"
	"time"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

type Advertisement struct {
	ID        int
	Title     string
	Image     *string
	Pdf       *string
	Url       *string
	CreatedAt time.Time
}

func NewAdvertisement(id int, title string, image *string, pdf *string, url *string, createdAt time.Time) Advertisement {
	return Advertisement{
		ID:        id,
		Title:     title,
		Image:     image,
		Pdf:       pdf,
		Url:       url,
		CreatedAt: createdAt,
	}
}

func NewAdvertisementUninitialized(title string, image *string, pdf *string, url *string) Advertisement {
	return Advertisement{
		ID:        UninitializedID,
		Title:     title,
		Image:     image,
		Pdf:       pdf,
		Url:       url,
		CreatedAt: time.Now(),
	}
}

func (a *Advertisement) Validate() error {
	titleLen := len([]rune(a.Title))
	if titleLen < 1 || titleLen > 200 {
		return fmt.Errorf("title must be between 1 and 200 characters: %d, %w", titleLen, core_errors.ErrInvalidArgument)
	}

	if a.Image != nil {
		imageLen := len([]rune(*a.Image))
		if imageLen < 1 || imageLen > 250 {
			return fmt.Errorf("image path must be between 1 and 250 characters: %d, %w", imageLen, core_errors.ErrInvalidArgument)
		}
	}

	if a.Pdf != nil {
		pdfLen := len([]rune(*a.Pdf))
		if pdfLen < 1 || pdfLen > 250 {
			return fmt.Errorf("pdf path must be between 1 and 250 characters: %d, %w", pdfLen, core_errors.ErrInvalidArgument)
		}
	}

	if a.Url != nil {
		urlLen := len([]rune(*a.Url))
		if urlLen < 1 || urlLen > 250 {
			return fmt.Errorf("url must be between 1 and 250 characters: %d, %w", urlLen, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

type AdvertisementPatch struct {
	Title Nullable[string]
	Image Nullable[string]
	Pdf   Nullable[string]
	Url   Nullable[string]
}

func NewAdvertisementPatch(title Nullable[string], image Nullable[string], pdf Nullable[string], url Nullable[string]) AdvertisementPatch {
	return AdvertisementPatch{
		Title: title,
		Image: image,
		Pdf:   pdf,
		Url:   url,
	}
}

func (p *AdvertisementPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("title cannot be null: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (a *Advertisement) ApplyPatch(patch AdvertisementPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid patch: %w", err)
	}

	tmp := *a

	if patch.Title.Set {
		if patch.Title.Value == nil {
			return fmt.Errorf("title cannot be null: %w", core_errors.ErrInvalidArgument)
		}
		tmp.Title = *patch.Title.Value
	}

	if patch.Image.Set {
		tmp.Image = patch.Image.Value
	}

	if patch.Pdf.Set {
		tmp.Pdf = patch.Pdf.Value
	}

	if patch.Url.Set {
		tmp.Url = patch.Url.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid advertisement after patch: %w", err)
	}

	*a = tmp
	return nil
}
