package domain

import (
	"fmt"
	"time"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

type VacantPosition struct {
	ID          int
	Title       string
	Description *string
	Date        *time.Time
}

func NewVacantPosition(id int, title string, description *string, date *time.Time) VacantPosition {
	return VacantPosition{
		ID:          id,
		Title:       title,
		Description: description,
		Date:        date,
	}
}

func NewVacantPositionUninitialized(title string, description *string, date *time.Time) VacantPosition {
	return VacantPosition{
		ID:          UninitializedID,
		Title:       title,
		Description: description,
		Date:        date,
	}
}

func (v *VacantPosition) Validate() error {
	titleLen := len([]rune(v.Title))
	if titleLen < 1 || titleLen > 200 {
		return fmt.Errorf("title must be between 1 and 200 characters: %d, %w", titleLen, core_errors.ErrInvalidArgument)
	}

	return nil
}

type VacantPositionPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Date        Nullable[time.Time]
}

func NewVacantPositionPatch(title Nullable[string], description Nullable[string], date Nullable[time.Time]) VacantPositionPatch {
	return VacantPositionPatch{
		Title:       title,
		Description: description,
		Date:        date,
	}
}

func (p *VacantPositionPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("title cannot be null: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (v *VacantPosition) ApplyPatch(patch VacantPositionPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid patch: %w", err)
	}

	tmp := *v

	if patch.Title.Set {
		if patch.Title.Value == nil {
			return fmt.Errorf("title cannot be null: %w", core_errors.ErrInvalidArgument)
		}
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Date.Set {
		tmp.Date = patch.Date.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid vacant position after patch: %w", err)
	}

	*v = tmp
	return nil
}
