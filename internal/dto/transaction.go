package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransactionResponseDTO struct {
	Id              uuid.UUID `json:"id" example:"d470a4f0-cd65-497d-9198-c16bbf670447"`
	AccountNumber   string    `json:"account_number" example:"233455011"`
	TransactionType string    `json:"transaction_type" example:"credit"`
	Amount          float64   `json:"amount" example:"10.3"`
	CreatedAt       time.Time `json:"created_at" example:"2025-02-22T15:11:19.25616+07:00"`
} // @name TransactionResponse

type CreateTransactionRequestDTO struct {
	AccountNumber   string  `json:"account_number" example:"233455011"`
	TransactionType string  `json:"transaction_type" example:"credit"`
	Amount          float64 `json:"amount" example:"10.3"`
} // @name CreateTransactionRequest

type UpdateTransactionRequestDTO struct {
	AccountNumber   string  `json:"account_number" validate:"required" example:"233455011"`
	TransactionType string  `json:"transaction_type" example:"credit"`
	Amount          float64 `json:"amount" example:"10.3"`
} // @name UpdateTransactionRequest

type UpdateTransactionResponseDTO struct {
	Id              uuid.UUID `json:"id" example:"d470a4f0-cd65-497d-9198-c16bbf670447"`
	AccountNumber   string    `json:"account_number" example:"233455011"`
	TransactionType string    `json:"transaction_type" example:"credit"`
	Amount          float64   `json:"amount" example:"10.3"`
	UpdatedAt       time.Time `json:"updated_at" example:"2025-02-22T15:11:19.25616+07:00"`
} // @name UpdateTransactionResponse

type DeleteTransactionResponseDTO struct {
	Message string `json:"message" example:"data deleted successfully"`
} // @name DeleteTransactionResponse

type CreateTransactionResponseDTO struct {
	CommonBaseResponseDTO
	Data TransactionResponseDTO `json:"data"`
} // @name CreateTransactionResponse

type GetOneTransactionResponseDTO struct {
	CommonBaseResponseDTO
	Data TransactionResponseDTO `json:"data"`
} // @name GetOneTransactionResponse

type GetAllTransactionsResponseDTO struct {
	CommonBaseResponseDTO
	Data []TransactionResponseDTO `json:"data"`
} // @name GetAllTransactionsResponse

type UpdateByIdTransactionResponseDTO struct {
	CommonBaseResponseDTO
	Data UpdateTransactionResponseDTO `json:"data"`
} // @name UpdateByIdTransactionResponse

type DeleteByIdTransactionResponseDTO struct {
	CommonBaseResponseDTO
	Data DeleteTransactionResponseDTO `json:"data"`
} // @name DeleteByIdTransactionResponse
