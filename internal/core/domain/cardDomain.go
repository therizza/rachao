package domain

import "github.com/google/uuid"

type Card struct {
	ID     uuid.UUID `json:"id"`
	IDPlay uuid.UUID `json:"id_play"`
	PAC    int       `json:"pac"`
	SHO    int       `json:"sho"`
	PAS    int       `json:"pas"`
	DRI    int       `json:"dri"`
	DEF    int       `json:"def"`
	PHY    int       `json:"phy"`
}

type CardRequest struct {
	PAC int `json:"pac"`
	SHO int `json:"sho"`
	PAS int `json:"pas"`
	DRI int `json:"dri"`
	DEF int `json:"def"`
	PHY int `json:"phy"`
}
