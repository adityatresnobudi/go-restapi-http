package errors

import "net/http"

type MessageErr interface {
	Error() string
	StatusCode() int
	Code() string
}

type ErrorData struct {
	ErrCode       string `json:"code"`
	ErrStatusCode int    `json:"status_code"`
	ErrMessage    string `json:"message"`
}

func (e *ErrorData) Error() string {
	return e.ErrMessage
}

func (e *ErrorData) StatusCode() int {
	return e.ErrStatusCode
}

func (e *ErrorData) Code() string {
	return e.ErrCode
}

func NewUnauthorizedError(message string) MessageErr {
	return &ErrorData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusForbidden,
		ErrCode:       "FORBIDDEN_ACCESS",
	}
}

func NewUnauthenticatedError(message string) MessageErr {
	return &ErrorData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnauthorized,
		ErrCode:       "UNAUTHORIZED",
	}
}

func NewConflictError(message string) MessageErr {
	return &ErrorData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusConflict,
		ErrCode:       "CONFLICT",
	}
}

func NewNotFoundError(message string) MessageErr {
	return &ErrorData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusNotFound,
		ErrCode:       "NOT_FOUND",
	}
}

func NewBadRequest(message string) MessageErr {
	return &ErrorData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusBadRequest,
		ErrCode:       "BAD_REQUEST",
	}
}

func NewInternalServerError() MessageErr {
	return &ErrorData{
		ErrMessage:    "Something went wrong",
		ErrStatusCode: http.StatusInternalServerError,
		ErrCode:       "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessibleEntityError(message string) MessageErr {
	return &ErrorData{
		ErrMessage:    message,
		ErrStatusCode: http.StatusUnprocessableEntity,
		ErrCode:       "UNPROCESSABLE_ENTITY",
	}
}

func NewTimeOutError() MessageErr {
	return &ErrorData{
		ErrMessage:    "The request took too long to process. Please try again.",
		ErrStatusCode: http.StatusRequestTimeout,
		ErrCode:       "REQUEST_TIME_OUT",
	}
}
