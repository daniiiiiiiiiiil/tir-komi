package http

import "time"

type AdvertisementDto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       *string   `json:"image,omitempty"`
	Pdf         *string   `json:"pdf,omitempty"`
	Url         *string   `json:"url,omitempty"`
	CreateAt    time.Time `json:"create_at"`
}
