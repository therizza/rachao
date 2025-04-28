package domain

type Healthz struct {
	Message string  `json:"Message"`
	Version float32 `json:"version"`
	Status  string  `json:"Status"`
}
