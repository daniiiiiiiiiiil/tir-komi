package postgres

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type MethodologicalMaterialModel struct {
	ID          int
	Title       string
	Description *string
	Date        *time.Time
	Pdf         *string
}

func methodologicalMaterialDomainFromModel(model MethodologicalMaterialModel) domain.MethodologicalMaterial {
	return domain.NewMethodologicalMaterial(
		model.ID,
		model.Title,
		model.Description,
		model.Date,
		model.Pdf,
	)
}

func methodologicalMaterialModelsFromDomains(materials []domain.MethodologicalMaterial) []MethodologicalMaterialModel {
	models := make([]MethodologicalMaterialModel, len(materials))
	for i, material := range materials {
		models[i] = MethodologicalMaterialModel{
			ID:          material.ID,
			Title:       material.Title,
			Description: material.Description,
			Date:        material.Date,
			Pdf:         material.Pdf,
		}
	}
	return models
}
