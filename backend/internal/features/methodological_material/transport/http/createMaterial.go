package http

type CreateMaterialRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Pdf         *[]byte `json:"pdf"`
}
