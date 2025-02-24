package transaction_repo

import (
	"context"

	"github.com/adityatresnobudi/go-restapi-http/internal/entity"
	"github.com/adityatresnobudi/go-restapi-http/pkg/errors"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]entity.Transaction, errors.MessageErr)
	GetOneById(ctx context.Context, id uuid.UUID) (*entity.Transaction, errors.MessageErr)
	Create(ctx context.Context, transaction entity.Transaction) (*entity.Transaction, errors.MessageErr)
	GetOneByAccountNumber(ctx context.Context, accountNumber string) (*entity.Transaction, errors.MessageErr)
	UpdateById(ctx context.Context, transaction entity.Transaction) (*entity.Transaction, errors.MessageErr)
	DeleteById(ctx context.Context, id uuid.UUID) errors.MessageErr
}
