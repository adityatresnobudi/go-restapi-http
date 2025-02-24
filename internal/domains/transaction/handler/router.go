package handler

import (
	"net/http"

	"github.com/adityatresnobudi/go-restapi-http/pkg/internal_http"
)

func (t *transactionHandler) MapRoutes() {
	t.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodGet, "/transactions"),
		t.GetAll(),
	)
	t.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodGet, "/transactions/{id}"),
		t.GetOne(),
	)
	t.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodPost, "/transactions"),
		t.Create(),
	)
	t.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodPut, "/transactions/{id}"),
		t.UpdateById(),
	)
	t.mux.HandleFunc(
		internal_http.NewAPIPath(http.MethodDelete, "/transactions/{id}"),
		t.DeleteById(),
	)
}
