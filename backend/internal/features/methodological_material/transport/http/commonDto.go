package http

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type MaterialDto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
	Pdf         []byte    `json:"pdf,omitempty"`
}

func convertMaterialDtoFromDomain(material domain.MethodologicalMaterial) MaterialDto {
	description := ""
	if material.Description != nil {
		description = *material.Description
	}

	createAt := time.Time{}
	if material.Date != nil {
		createAt = *material.Date
	}

	return MaterialDto{
		Id:          material.ID,
		Title:       material.Title,
		Description: description,
		CreateAt:    createAt,
		Pdf:         material.Pdf,
	}
}

func convertMaterialDtosFromDomains(materials []domain.MethodologicalMaterial) []MaterialDto {
	dtos := make([]MaterialDto, len(materials))
	for i, material := range materials {
		dtos[i] = convertMaterialDtoFromDomain(material)
	}
	return dtos
}
