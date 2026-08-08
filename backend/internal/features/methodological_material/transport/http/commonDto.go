package http

import "time"

type MaterialDto struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
	Pdf         *string   `json:"pdf,omitempty"`
}
