package domain

import (
	"fmt"
	"time"

	core_errors "github.com/daniiiiiiiiiiil/tir-komi/internal/core/errors"
)

type Review struct {
	ID          int
	Name        string
	Email       *string
	Description *string
	Rating      int
	CreatedAt   time.Time
}

func NewReview(id int, name string, email *string, description *string, rating int, createdAt time.Time) Review {
	return Review{
		ID:          id,
		Name:        name,
		Email:       email,
		Description: description,
		Rating:      rating,
		CreatedAt:   createdAt,
	}
}

func NewReviewUninitialized(name string, email *string, description *string, rating int) Review {
	return Review{
		ID:          UninitializedID,
		Name:        name,
		Email:       email,
		Description: description,
		Rating:      rating,
		CreatedAt:   time.Now(),
	}
}

func (r *Review) Validate() error {
	nameLen := len([]rune(r.Name))
	if nameLen < 1 || nameLen > 100 {
		return fmt.Errorf("name must be between 1 and 100 characters: %d, %w", nameLen, core_errors.ErrInvalidArgument)
	}

	if r.Rating < 1 || r.Rating > 5 {
		return fmt.Errorf("rating must be between 1 and 5: %d, %w", r.Rating, core_errors.ErrInvalidArgument)
	}

	return nil
}

type ReviewPatch struct {
	Name        Nullable[string]
	Email       Nullable[string]
	Description Nullable[string]
	Rating      Nullable[int]
}

func NewReviewPatch(name Nullable[string], email Nullable[string], description Nullable[string], rating Nullable[int]) ReviewPatch {
	return ReviewPatch{
		Name:        name,
		Email:       email,
		Description: description,
		Rating:      rating,
	}
}

func (p *ReviewPatch) Validate() error {
	if p.Name.Set && p.Name.Value == nil {
		return fmt.Errorf("name cannot be null: %w", core_errors.ErrInvalidArgument)
	}
	if p.Rating.Set && p.Rating.Value != nil && (*p.Rating.Value < 1 || *p.Rating.Value > 5) {
		return fmt.Errorf("rating must be between 1 and 5: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (r *Review) ApplyPatch(patch ReviewPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid patch: %w", err)
	}

	tmp := *r

	if patch.Name.Set {
		if patch.Name.Value == nil {
			return fmt.Errorf("name cannot be null: %w", core_errors.ErrInvalidArgument)
		}
		tmp.Name = *patch.Name.Value
	}

	if patch.Email.Set {
		tmp.Email = patch.Email.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Rating.Set {
		if patch.Rating.Value == nil {
			return fmt.Errorf("rating cannot be null: %w", core_errors.ErrInvalidArgument)
		}
		tmp.Rating = *patch.Rating.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("invalid review after patch: %w", err)
	}

	*r = tmp
	return nil
}
