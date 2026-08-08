package postgres

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type VacantPositionModel struct {
	ID          int
	Title       string
	Description *string
	Date        *time.Time
}

func vacantPositionDomainFromModel(model VacantPositionModel) domain.VacantPosition {
	return domain.NewVacantPosition(
		model.ID,
		model.Title,
		model.Description,
		model.Date,
	)
}

func vacantPositionModelsFromDomains(positions []domain.VacantPosition) []VacantPositionModel {
	models := make([]VacantPositionModel, len(positions))
	for i, position := range positions {
		models[i] = VacantPositionModel{
			ID:          position.ID,
			Title:       position.Title,
			Description: position.Description,
			Date:        position.Date,
		}
	}
	return models
}
