package http

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type VacantDto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}

func convertVacantDtoFromDomain(vacant domain.VacantPosition) VacantDto {
	description := ""
	if vacant.Description != nil {
		description = *vacant.Description
	}

	createAt := time.Time{}
	if vacant.Date != nil {
		createAt = *vacant.Date
	}

	return VacantDto{
		Id:          vacant.ID,
		Title:       vacant.Title,
		Description: description,
		CreateAt:    createAt,
	}
}

func convertVacantDtosFromDomains(vacancies []domain.VacantPosition) []VacantDto {
	dtos := make([]VacantDto, len(vacancies))
	for i, vacant := range vacancies {
		dtos[i] = convertVacantDtoFromDomain(vacant)
	}
	return dtos
}
