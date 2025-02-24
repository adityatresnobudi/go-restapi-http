package internal_http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adityatresnobudi/go-restapi-http/internal/dto"
)

func NewOKStatusCommonBaseResponseDTO() dto.CommonBaseResponseDTO {
	return dto.CommonBaseResponseDTO{
		Status:     "OK",
		StatusCode: http.StatusOK,
	}
}

func NewCreatedStatusCommonBaseResponseDTO() dto.CommonBaseResponseDTO {
	return dto.CommonBaseResponseDTO{
		Status:     "OK",
		StatusCode: http.StatusCreated,
	}
}

func NewDeletedStatusCommonBaseResponseDTO() dto.CommonBaseResponseDTO {
	return dto.CommonBaseResponseDTO{
		Status:     "No Content",
		StatusCode: http.StatusNoContent,
	}
}

func NewAPIPath(method string, path string) string {
	return fmt.Sprintf("%s %s", method, path)
}

func SendResponse(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
