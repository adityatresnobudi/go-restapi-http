package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-http/internal/dto"
	"github.com/adityatresnobudi/go-restapi-http/internal/entity"
	"github.com/adityatresnobudi/go-restapi-http/internal/repositories/transaction_repo"
	"github.com/adityatresnobudi/go-restapi-http/pkg/errors"
	"github.com/adityatresnobudi/go-restapi-http/pkg/internal_http"
	"github.com/google/uuid"
)

type TransactionService interface {
	GetAll(ctx context.Context) (*dto.GetAllTransactionsResponseDTO, errors.MessageErr)
	GetOne(ctx context.Context, id string) (*dto.GetOneTransactionResponseDTO, errors.MessageErr)
	Create(ctx context.Context, payload dto.CreateTransactionRequestDTO) (*dto.CreateTransactionResponseDTO, errors.MessageErr)
	UpdateById(ctx context.Context, id string, payload dto.UpdateTransactionRequestDTO) (*dto.UpdateByIdTransactionResponseDTO, errors.MessageErr)
	DeleteById(ctx context.Context, id string) (*dto.DeleteByIdTransactionResponseDTO, errors.MessageErr)
}

type transactionServiceIMPL struct {
	transactionRepo transaction_repo.Repository
}

func NewTransactionService(transactionRepo transaction_repo.Repository) TransactionService {
	return &transactionServiceIMPL{
		transactionRepo: transactionRepo,
	}
}

func (t *transactionServiceIMPL) GetAll(ctx context.Context) (*dto.GetAllTransactionsResponseDTO, errors.MessageErr) {
	transactions, err := t.transactionRepo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	result := dto.GetAllTransactionsResponseDTO{
		CommonBaseResponseDTO: internal_http.NewOKStatusCommonBaseResponseDTO(),
		Data:                  entity.Transactions(transactions).ToSliceOfTransactionResponseDTO(),
	}

	return &result, nil
}
func (t *transactionServiceIMPL) GetOne(ctx context.Context, id string) (*dto.GetOneTransactionResponseDTO, errors.MessageErr) {
	// uuid.Parse function merupakan function untuk merubah id yang tadinya string menjadi uuid.UUID
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errors.NewBadRequest("id has to be a valid uuid")
	}

	transaction, err := t.transactionRepo.GetOneById(ctx, parsedId)

	if err != nil {
		return nil, err
	}

	result := dto.GetOneTransactionResponseDTO{
		CommonBaseResponseDTO: internal_http.NewOKStatusCommonBaseResponseDTO(),
		Data:                  *transaction.ToTransactionResponseDTO(),
	}

	return &result, nil
}
func (t *transactionServiceIMPL) Create(
	ctx context.Context,
	payload dto.CreateTransactionRequestDTO,
) (*dto.CreateTransactionResponseDTO, errors.MessageErr) {
	if err := t.createValidator(payload); err != nil {
		return nil, err
	}

	existingTransaction, err := t.transactionRepo.GetOneByAccountNumber(
		ctx,
		payload.AccountNumber,
	)

	if err != nil && err.StatusCode() != http.StatusNotFound {
		return nil, err
	}

	if existingTransaction != nil {
		return nil, errors.NewConflictError("please choose another account number")
	}

	transaction := entity.Transaction{
		AccountNumber:   payload.AccountNumber,
		TransactionType: entity.TransactionVariant(payload.TransactionType),
		Amount:          payload.Amount,
	}

	newTransaction, err := t.transactionRepo.Create(ctx, transaction)

	if err != nil {
		return nil, err
	}

	result := dto.CreateTransactionResponseDTO{
		CommonBaseResponseDTO: internal_http.NewCreatedStatusCommonBaseResponseDTO(),
		Data:                  *newTransaction.ToTransactionResponseDTO(),
	}

	return &result, nil
}

func (t *transactionServiceIMPL) UpdateById(ctx context.Context, id string, payload dto.UpdateTransactionRequestDTO) (*dto.UpdateByIdTransactionResponseDTO, errors.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errors.NewBadRequest("id has to be a valid uuid")
	}

	if err := t.updateValidator(payload); err != nil {
		return nil, err
	}

	transaction, err := t.transactionRepo.GetOneById(ctx, parsedId)
	if err != nil {
		// fmt.Println(parsedId)
		return nil, err
	}

	txByAccNum, err := t.transactionRepo.GetOneByAccountNumber(ctx, payload.AccountNumber)
	if err != nil {
		fmt.Println(txByAccNum.Id, parsedId)
		return nil, err
	}
	fmt.Println(txByAccNum.Id, parsedId)
	if txByAccNum.Id != parsedId {
		return nil, errors.NewConflictError("duplicate account number")
	}

	transaction.Id = parsedId
	transaction.AccountNumber = payload.AccountNumber
	transaction.Amount = payload.Amount
	transaction.TransactionType = entity.TransactionVariant(payload.TransactionType)

	response, err := t.transactionRepo.UpdateById(ctx, *transaction)
	if err != nil {
		return nil, err
	}

	result := dto.UpdateByIdTransactionResponseDTO{
		CommonBaseResponseDTO: internal_http.NewOKStatusCommonBaseResponseDTO(),
		Data:                  *response.ToUpdateTransactionResponseDTO(),
	}

	return &result, nil
}

func (t *transactionServiceIMPL) DeleteById(ctx context.Context, id string) (*dto.DeleteByIdTransactionResponseDTO, errors.MessageErr) {
	parsedId, errParseId := uuid.Parse(id)

	if errParseId != nil {
		return nil, errors.NewBadRequest("id has to be a valid uuid")
	}

	_, err := t.transactionRepo.GetOneById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	err = t.transactionRepo.DeleteById(ctx, parsedId)
	if err != nil {
		return nil, err
	}

	result := dto.DeleteByIdTransactionResponseDTO{
		CommonBaseResponseDTO: internal_http.NewDeletedStatusCommonBaseResponseDTO(),
		Data: dto.DeleteTransactionResponseDTO{
			Message: "Transaction data deleted successfully",
		},
	}

	return &result, nil
}
