package http

type CreateAdvertisementRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       *[]byte `json:"image,omitempty"`
	Pdf         *[]byte `json:"pdf,omitempty"`
	Url         *string `json:"url"`
}
