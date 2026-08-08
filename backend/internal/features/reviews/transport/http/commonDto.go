package http

import "time"

type ReviewDto struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	Rating      int       `json:"rating"`
	CreateAt    time.Time `json:"create_at"`
}
