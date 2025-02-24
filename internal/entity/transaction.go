package entity

import (
	"time"

	"github.com/adityatresnobudi/go-restapi-http/internal/dto"
	"github.com/google/uuid"
)

type TransactionVariant string

const (
	CREDIT TransactionVariant = "credit"
	DEBIT  TransactionVariant = "debit"
)

var TransactionVariants map[TransactionVariant]bool = map[TransactionVariant]bool{
	CREDIT: true,
	DEBIT:  true,
}

type Transaction struct {
	Id              uuid.UUID
	AccountNumber   string
	TransactionType TransactionVariant
	Amount          float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Transactions []Transaction

func (t *Transaction) ToTransactionResponseDTO() *dto.TransactionResponseDTO {
	return &dto.TransactionResponseDTO{
		Id:              t.Id,
		AccountNumber:   t.AccountNumber,
		TransactionType: string(t.TransactionType),
		Amount:          t.Amount,
		CreatedAt:       t.CreatedAt,
	}
}

func (t Transactions) ToSliceOfTransactionResponseDTO() []dto.TransactionResponseDTO {
	result := []dto.TransactionResponseDTO{}
	for _, tr := range t {
		result = append(result, *tr.ToTransactionResponseDTO())
	}

	return result
}

func (t *Transaction) ToUpdateTransactionResponseDTO() *dto.UpdateTransactionResponseDTO {
	return &dto.UpdateTransactionResponseDTO{
		Id:              t.Id,
		AccountNumber:   t.AccountNumber,
		TransactionType: string(t.TransactionType),
		Amount:          t.Amount,
		UpdatedAt:       t.UpdatedAt,
	}
}
