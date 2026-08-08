package domain

import (
	"fmt"
	"time"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

type MethodologicalMaterial struct {
	ID          int
	Title       string
	Description *string
	Date        *time.Time
	Pdf         *string
}

func NewMethodologicalMaterial(id int, title string, description *string, date *time.Time, pdf *string) MethodologicalMaterial {
	return MethodologicalMaterial{
		ID:          id,
		Title:       title,
		Description: description,
		Date:        date,
		Pdf:         pdf,
	}
}

func NewMethodologicalMaterialUninitialized(title string, description *string, date *time.Time, pdf *string) MethodologicalMaterial {
	return MethodologicalMaterial{
		ID:          UninitializedID,
		Title:       title,
		Description: description,
		Date:        date,
		Pdf:         pdf,
	}
}

func (m *MethodologicalMaterial) Validate() error {
	titleLen := len([]rune(m.Title))
	if titleLen < 1 || titleLen > 200 {
		return fmt.Errorf("title must be between 1 and 200 characters: %d, %w", titleLen, core_errors.ErrInvalidArgument)
	}

	if m.Pdf != nil {
		pdfLen := len([]rune(*m.Pdf))
		if pdfLen > 200 {
			return fmt.Errorf("pdf path too long: %d, %w", pdfLen, core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

type MethodologicalMaterialPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Date        Nullable[time.Time]
	Pdf         Nullable[string]
}

func NewMethodologicalMaterialPatch(title Nullable[string], description Nullable[string], date Nullable[time.Time], pdf Nullable[string]) MethodologicalMaterialPatch {
	return MethodologicalMaterialPatch{
		Title:       title,
		Description: description,
		Date:        date,
		Pdf:         pdf,
	}
}

func (p *MethodologicalMaterialPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("title cannot be null: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (m *MethodologicalMaterial) ApplyPatch(patch MethodologicalMaterialPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid patch: %w", err)
	}

	tmp := *m

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

	if patch.Pdf.Set {
		tmp.Pdf = patch.Pdf.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid methodological material after patch: %w", err)
	}

	*m = tmp
	return nil
}
