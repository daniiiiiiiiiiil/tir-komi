package http

import (
	"time"

	"github.com/daniiiiiiiiiiil/tir-komi/internal/core/domain"
)

type AdvertisementDto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       *string   `json:"image,omitempty"`
	Pdf         *string   `json:"pdf,omitempty"`
	Url         *string   `json:"url,omitempty"`
	CreateAt    time.Time `json:"create_at"`
}

func convertAdvertisementDtoFromDomain(ad domain.Advertisement) AdvertisementDto {
	description := ""
	if ad.Description != nil {
		description = *ad.Description
	}

	return AdvertisementDto{
		Id:          ad.ID,
		Title:       ad.Title,
		Description: description,
		Image:       ad.Image,
		Pdf:         ad.Pdf,
		Url:         ad.Url,
		CreateAt:    ad.CreatedAt,
	}
}

func convertAdvertisementDtosFromDomains(ads []domain.Advertisement) []AdvertisementDto {
	dtos := make([]AdvertisementDto, len(ads))
	for i, ad := range ads {
		dtos[i] = convertAdvertisementDtoFromDomain(ad)
	}
	return dtos
}
