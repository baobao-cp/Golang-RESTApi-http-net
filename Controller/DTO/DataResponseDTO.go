package DTO

type DataResponseDTO struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
