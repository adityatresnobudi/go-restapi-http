package service

import (
	"fmt"
	"strings"

	"github.com/adityatresnobudi/restapi/internal/dto"
	"github.com/adityatresnobudi/restapi/internal/entity"
	"github.com/adityatresnobudi/restapi/pkg/errors"
)

func (t *transactionServiceIMPL) createValidator(payload dto.CreateTransactionRequestDTO) errors.MessageErr {
	errArr := make([]errors.MessageErr, 0)

	if strings.TrimSpace(payload.AccountNumber) == "" {
		errArr = append(errArr, errors.NewBadRequest("account number cannot be empty"))
	}

	if !entity.TransactionVariants[entity.TransactionVariant(payload.TransactionType)] {
		errArr = append(errArr, errors.NewBadRequest(
			fmt.Sprintf(
				"transaction type has to be either %s or %s",
				entity.CREDIT,
				entity.DEBIT,
			),
		))
	}

	if payload.Amount < 1 {
		errArr = append(errArr, errors.NewBadRequest("amount cannot be less than 1"))
	}

	if len(errArr) > 0 {
		msgArr := make([]string, 0)
		for _, value := range errArr {
			msgArr = append(msgArr, value.Error())
		}
		errMsg := strings.Join(msgArr, ", ")
		err := errors.ErrorData{
			ErrCode:       errArr[0].Code(),
			ErrStatusCode: errArr[0].StatusCode(),
			ErrMessage:    errMsg,
		}
		return &err
	}

	return nil
}

func (t *transactionServiceIMPL) updateValidator(payload dto.UpdateTransactionRequestDTO) errors.MessageErr {
	errArr := make([]errors.MessageErr, 0)

	if strings.TrimSpace(payload.AccountNumber) == "" {
		errArr = append(errArr, errors.NewBadRequest("account number cannot be empty"))
	}

	if !entity.TransactionVariants[entity.TransactionVariant(payload.TransactionType)] {
		errArr = append(errArr, errors.NewBadRequest(
			fmt.Sprintf(
				"transaction type has to be either %s or %s",
				entity.CREDIT,
				entity.DEBIT,
			),
		))
	}

	if payload.Amount < 1 {
		errArr = append(errArr, errors.NewBadRequest("amount cannot be less than 1"))
	}

	if len(errArr) > 0 {
		msgArr := make([]string, 0)
		for _, value := range errArr {
			msgArr = append(msgArr, value.Error())
		}
		errMsg := strings.Join(msgArr, ", ")
		err := errors.ErrorData{
			ErrCode:       errArr[0].Code(),
			ErrStatusCode: errArr[0].StatusCode(),
			ErrMessage:    errMsg,
		}
		return &err
	}

	return nil
}
