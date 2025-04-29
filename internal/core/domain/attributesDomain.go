package domain

type Attributes struct {
	ID         int `json:"id"`
	IDPosition int `json:"id_position"`
	PAC        int `json:"pac"`
	SHO        int `json:"sho"`
	PAS        int `json:"pas"`
	DRI        int `json:"dri"`
	DEF        int `json:"def"`
	PHY        int `json:"phy"`
}

type AttributesRequest struct {
	IDPosition int `json:"id_position"`
	PAC        int `json:"pac"`
	SHO        int `json:"sho"`
	PAS        int `json:"pas"`
	DRI        int `json:"dri"`
	DEF        int `json:"def"`
	PHY        int `json:"phy"`
}
