package domain

import "github.com/google/uuid"

type Photo struct {
	ID     uuid.UUID `json:"id"`
	IDPlay uuid.UUID `json:"id_play"`
	Photo  []byte    `json:"photo"`
}
