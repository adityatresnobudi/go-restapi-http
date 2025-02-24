package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-http/internal/domains/transaction/service"
	"github.com/adityatresnobudi/go-restapi-http/internal/dto"
	"github.com/adityatresnobudi/go-restapi-http/pkg/errors"
	"github.com/adityatresnobudi/go-restapi-http/pkg/internal_http"
)

type transactionHandler struct {
	mux     *http.ServeMux
	ctx     context.Context
	service service.TransactionService
}

func NewTransactionHandler(
	mux *http.ServeMux,
	ctx context.Context,
	service service.TransactionService,
) *transactionHandler {
	return &transactionHandler{
		mux:     mux,
		ctx:     ctx,
		service: service,
	}
}

// @Summary Get All Transactions
// @Tags transactions
// @Produce json
// @Success 200 {object}  GetAllTransactionsResponse
// @Router /transactions [get]
func (t *transactionHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		result, err := t.service.GetAll(t.ctx)

		if err != nil {
			internal_http.SendResponse(w, err.StatusCode(), err)
			return
		}

		internal_http.SendResponse(w, result.StatusCode, result)
	}
}

// @Summary Get One Transaction By ID
// @Tags transactions
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object}  GetOneTransactionResponse
// @Router /transactions/{id} [get]
func (t *transactionHandler) GetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		result, errData := t.service.GetOne(t.ctx, id)

		if errData != nil {
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		internal_http.SendResponse(w, result.StatusCode, result)
	}
}

// @Summary Create Transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param requestBody body CreateTransactionRequest true "Request Body"
// @Success 200 {object} CreateTransactionResponse
// @Router /transactions [post]
func (t *transactionHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		payload := dto.CreateTransactionRequestDTO{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			errData := errors.NewUnprocessibleEntityError(err.Error())
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		result, errData := t.service.Create(t.ctx, payload)

		if errData != nil {
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		internal_http.SendResponse(w, result.StatusCode, result)
	}
}

// @Summary Update Transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Param requestBody body UpdateTransactionRequest true "Request Body"
// @Success 200 {object} UpdateByIdTransactionResponse
// @Router /transactions [put]
func (t *transactionHandler) UpdateById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")
		payload := dto.UpdateTransactionRequestDTO{}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			errData := errors.NewUnprocessibleEntityError(err.Error())
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		result, errData := t.service.UpdateById(t.ctx, id, payload)

		if errData != nil {
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		internal_http.SendResponse(w, result.StatusCode, result)
	}
}

// @Summary Delete Transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 204 {object} DeleteByIdTransactionResponse
// @Router /transactions [delete]
func (t *transactionHandler) DeleteById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		result, errData := t.service.DeleteById(t.ctx, id)

		if errData != nil {
			internal_http.SendResponse(w, errData.StatusCode(), errData)
			return
		}

		internal_http.SendResponse(w, result.StatusCode, result)
	}
}
