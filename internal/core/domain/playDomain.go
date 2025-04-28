package domain

import "github.com/google/uuid"

type Play struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	IDPosition int       `json:"id_position"`
	IDNation   int       `json:"id_nation"`
	Field      bool      `json:"field"`
	Active     bool      `json:"active"`
}

type CreatePlayRequest struct {
	Name       string `json:"name"`
	IDPosition int    `json:"id_position"`
	IDNation   int    `json:"id_nation"`
	Field      bool   `json:"field"`
	Active     bool   `json:"active"`
}
