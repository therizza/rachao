package domain

type Position struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}

type CreatePositionRequest struct {
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}
