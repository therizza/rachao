package domain

type Nation struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}

type CreateNationRequest struct {
	Name    string `json:"name"`
	Acronym string `json:"acronym"`
}
