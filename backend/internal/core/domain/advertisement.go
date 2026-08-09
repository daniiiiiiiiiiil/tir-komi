package domain

import (
	"fmt"
	"time"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

type Advertisement struct {
	ID          int
	Title       string
	Description *string
	Image       []byte
	Pdf         []byte
	Url         *string
	CreatedAt   time.Time
}

func NewAdvertisement(id int, title string, description *string, image []byte, pdf []byte, url *string, createdAt time.Time) Advertisement {
	return Advertisement{
		ID:          id,
		Title:       title,
		Description: description,
		Image:       image,
		Pdf:         pdf,
		Url:         url,
		CreatedAt:   createdAt,
	}
}

func NewAdvertisementUninitialized(title string, description *string, image []byte, pdf []byte, url *string) Advertisement {
	return Advertisement{
		ID:          UninitializedID,
		Title:       title,
		Description: description,
		Image:       image,
		Pdf:         pdf,
		Url:         url,
		CreatedAt:   time.Now(),
	}
}

func (a *Advertisement) Validate() error {
	titleLen := len([]rune(a.Title))
	if titleLen < 1 || titleLen > 200 {
		return fmt.Errorf("title must be between 1 and 200 characters: %d, %w", titleLen, core_errors.ErrInvalidArgument)
	}

	if a.Description != nil {
		descLen := len([]rune(*a.Description))
		if descLen > 1000 {
			return fmt.Errorf("description too long: %d, max 1000 characters, %w", descLen, core_errors.ErrInvalidArgument)
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
	Title       Nullable[string]
	Description Nullable[string]
	Image       Nullable[[]byte]
	Pdf         Nullable[[]byte]
	Url         Nullable[string]
}

func NewAdvertisementPatch(title Nullable[string], description Nullable[string], image Nullable[[]byte], pdf Nullable[[]byte], url Nullable[string]) AdvertisementPatch {
	return AdvertisementPatch{
		Title:       title,
		Description: description,
		Image:       image,
		Pdf:         pdf,
		Url:         url,
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

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Image.Set {
		if patch.Image.Value == nil {
			tmp.Image = []byte{}
		} else {
			tmp.Image = *patch.Image.Value
		}
	}

	if patch.Pdf.Set {
		if patch.Pdf.Value == nil {
			tmp.Pdf = []byte{}
		} else {
			tmp.Pdf = *patch.Pdf.Value
		}
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
