package http

type CreateVacantRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
