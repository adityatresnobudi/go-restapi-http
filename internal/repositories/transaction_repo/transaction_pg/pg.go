package transaction_pg

import (
	"context"
	"database/sql"
	"log"

	"github.com/adityatresnobudi/go-restapi-http/internal/entity"
	"github.com/adityatresnobudi/go-restapi-http/internal/repositories/transaction_repo"
	"github.com/adityatresnobudi/go-restapi-http/pkg/errors"
	"github.com/google/uuid"
)

type transactionPG struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) transaction_repo.Repository {
	return &transactionPG{
		db: db,
	}
}

func (t *transactionPG) GetAll(ctx context.Context) ([]entity.Transaction, errors.MessageErr) {
	rows, err := t.db.QueryContext(ctx, GET_ALL_TRANSACTIONS)

	if err != nil {
		log.Printf("db get all transactions: %s\n", err.Error())
		return nil, errors.NewInternalServerError()
	}

	result := []entity.Transaction{}

	for rows.Next() {
		transaction := entity.Transaction{}

		if err = rows.Scan(
			&transaction.Id,
			&transaction.AccountNumber,
			&transaction.TransactionType,
			&transaction.Amount,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		); err != nil {
			log.Printf("db scan get all transactions: %s\n", err.Error())
			return nil, errors.NewInternalServerError()
		}

		result = append(result, transaction)
	}

	return result, nil
}
func (t *transactionPG) GetOneById(ctx context.Context, id uuid.UUID) (*entity.Transaction, errors.MessageErr) {
	transaction := entity.Transaction{}

	if err := t.db.QueryRowContext(
		ctx,
		GET_ONE_TRANSACTION_BY_ID,
		id,
	).Scan(
		&transaction.Id,
		&transaction.AccountNumber,
		&transaction.TransactionType,
		&transaction.Amount,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	); err != nil {
		log.Printf("db scan get one transaction by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("transaction was not found")
		}
		return nil, errors.NewInternalServerError()
	}

	return &transaction, nil
}
func (t *transactionPG) Create(ctx context.Context, transaction entity.Transaction) (*entity.Transaction, errors.MessageErr) {
	newTransaction := entity.Transaction{}

	if err := t.db.QueryRowContext(
		ctx,
		INSERT_TRANSACTION,
		transaction.AccountNumber,
		transaction.TransactionType,
		transaction.Amount,
	).Scan(
		&newTransaction.Id,
		&newTransaction.AccountNumber,
		&newTransaction.TransactionType,
		&newTransaction.Amount,
		&newTransaction.CreatedAt,
		&newTransaction.UpdatedAt,
	); err != nil {
		log.Printf("db scan create transaction: %s\n", err.Error())
		return nil, errors.NewInternalServerError()
	}

	return &newTransaction, nil
}
func (t *transactionPG) GetOneByAccountNumber(ctx context.Context, accountNumber string) (*entity.Transaction, errors.MessageErr) {
	transaction := entity.Transaction{}

	if err := t.db.QueryRowContext(
		ctx,
		GET_ONE_TRANSACTION_BY_ACCOUNT_NUMBER,
		accountNumber,
	).Scan(
		&transaction.Id,
		&transaction.AccountNumber,
		&transaction.TransactionType,
		&transaction.Amount,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	); err != nil {
		log.Printf("db scan get one transaction by account number: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("transaction was not found")
		}
		return nil, errors.NewInternalServerError()
	}

	return &transaction, nil
}
func (t *transactionPG) UpdateById(ctx context.Context, transaction entity.Transaction) (*entity.Transaction, errors.MessageErr) {
	response := entity.Transaction{}

	if err := t.db.QueryRowContext(
		ctx,
		UPDATE_TRANSACTION,
		transaction.AccountNumber,
		transaction.TransactionType,
		transaction.Amount,
		transaction.Id,
	).Scan(
		&response.Id,
		&response.AccountNumber,
		&response.TransactionType,
		&response.Amount,
		&response.UpdatedAt,
	); err != nil {
		log.Printf("db scan update transaction by id: %s\n", err.Error())
		if err == sql.ErrNoRows {
			return nil, errors.NewNotFoundError("transaction was not found")
		}
		return nil, errors.NewInternalServerError()
	}

	return &response, nil
}
func (t *transactionPG) DeleteById(ctx context.Context, id uuid.UUID) errors.MessageErr {
	if _, err := t.db.ExecContext(
		ctx,
		DELETE_TRANSACTION,
		id,
	); err != nil {
		log.Printf("db delete transaction by id: %s\n", err.Error())
		return errors.NewInternalServerError()
	}

	return nil
}
