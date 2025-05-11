package domain

type Modality struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Amount_play int    `json:"amount_play"`
	Active      bool   `json:"active"`
}

type CreateModalityRequest struct {
	Name        string `json:"name"`
	Amount_play int    `json:"amount_play"`
	Active      bool   `json:"active"`
}
