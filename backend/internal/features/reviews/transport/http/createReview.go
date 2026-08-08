package http

type CreateReviewRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}
