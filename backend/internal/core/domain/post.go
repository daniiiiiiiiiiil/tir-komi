package domain

import (
	"errors"
	"fmt"
	"time"
)

// PostType - тип поста (перечисление)
type PostType string

const (
	PostTypeInformation              PostType = "information"
	PostTypeStructuresAndBodies      PostType = "structures_and_bodies"
	PostTypeDocument                 PostType = "document"
	PostTypeEducation                PostType = "education"
	PostTypeEducationalStandards     PostType = "educational_standards"
	PostTypeManagement               PostType = "management"
	PostTypeMaterials                PostType = "materials"
	PostTypeScholarships             PostType = "scholarships"
	PostTypePaidServices             PostType = "paid_services"
	PostTypeFinancialAndEconomic     PostType = "financial_and_economic"
	PostTypeAccessibleEnvironment    PostType = "accessible_environment"
	PostTypeInternationalCooperation PostType = "international_cooperation"
)

// ValidPostTypes - валидные типы постов
var ValidPostTypes = map[PostType]bool{
	PostTypeInformation:              true,
	PostTypeStructuresAndBodies:      true,
	PostTypeDocument:                 true,
	PostTypeEducation:                true,
	PostTypeEducationalStandards:     true,
	PostTypeManagement:               true,
	PostTypeMaterials:                true,
	PostTypeScholarships:             true,
	PostTypePaidServices:             true,
	PostTypeFinancialAndEconomic:     true,
	PostTypeAccessibleEnvironment:    true,
	PostTypeInternationalCooperation: true,
}

// Post - доменная модель поста
type Post struct {
	ID          int
	Title       string
	Description *string
	Pdf         []byte
	Image       []byte
	Date        *time.Time
	Type        PostType
}

// NewPost - создание нового поста
func NewPost(id int, title string, description *string, pdf, image []byte, date *time.Time, postType PostType) Post {
	return Post{
		ID:          id,
		Title:       title,
		Description: description,
		Pdf:         pdf,
		Image:       image,
		Date:        date,
		Type:        postType,
	}
}

// NewPostUninitialized - создание поста без ID
func NewPostUninitialized(title string, description *string, pdf, image []byte, date *time.Time, postType PostType) Post {
	return Post{
		Title:       title,
		Description: description,
		Pdf:         pdf,
		Image:       image,
		Date:        date,
		Type:        postType,
	}
}

// Validate - валидация поста
func (p *Post) Validate() error {
	if len(p.Title) == 0 || len(p.Title) > 1000 {
		return errors.New("title length must be between 1 and 1000 characters")
	}

	if p.Description != nil && len(*p.Description) > 10000 {
		return errors.New("description length must not exceed 10000 characters")
	}

	if len(p.Pdf) > 50*1024*1024 { // 50 MB
		return errors.New("pdf size must not exceed 50 MB")
	}

	if len(p.Image) > 10*1024*1024 { // 10 MB
		return errors.New("image size must not exceed 10 MB")
	}

	if p.Date == nil {
		return errors.New("date is required")
	}

	if !ValidPostTypes[p.Type] {
		return fmt.Errorf("invalid post type: %s", p.Type)
	}

	return nil
}

type PostPatch struct {
	Title       Nullable[string]
	Description Nullable[*string]
	Pdf         Nullable[[]byte]
	Image       Nullable[[]byte]
	Date        Nullable[*time.Time]
	Type        Nullable[*PostType]
}

func NewPostPatch(
	title Nullable[string],
	description Nullable[*string],
	pdf Nullable[[]byte],
	image Nullable[[]byte],
	date Nullable[*time.Time],
	postType Nullable[*PostType],
) PostPatch {
	return PostPatch{
		Title:       title,
		Description: description,
		Pdf:         pdf,
		Image:       image,
		Date:        date,
		Type:        postType,
	}
}

// ApplyPatch - применение патча к посту
func (p *Post) ApplyPatch(patch PostPatch) error {
	if patch.Title.Set {
		if patch.Title.Value == nil {
			return errors.New("title cannot be null")
		}
		p.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		p.Description = *patch.Description.Value
	}

	if patch.Pdf.Set {
		p.Pdf = *patch.Pdf.Value
	}

	if patch.Image.Set {
		p.Image = *patch.Image.Value
	}

	if patch.Date.Set {
		p.Date = *patch.Date.Value
	}

	if patch.Type.Set {
		if patch.Type.Value == nil {
			return errors.New("type cannot be null")
		}
		if !ValidPostTypes[**patch.Type.Value] {
			return fmt.Errorf("invalid post type: %s", *patch.Type.Value)
		}
		rd := *patch.Type.Value
		p.Type = *rd
	}

	return nil
}
