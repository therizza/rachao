package domain

import "github.com/google/uuid"

type Overall struct {
	ID      uuid.UUID `json:"id"`
	IDPlay  uuid.UUID `json:"id_play"`
	Overall int       `json:"overall"`
}

type OverallBodyRequest struct {
	Card       Card       `json:"card"`
	Attributes Attributes `json:"attributes"`
}

type OverallRequest struct {
	IDPlay  uuid.UUID `json:"id_play"`
	Overall int       `json:"overall"`
}
