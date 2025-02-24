package dto

type CommonBaseResponseDTO struct {
	Status     string `json:"status" example:"OK"`
	StatusCode int    `json:"status_code" example:"200"`
}
