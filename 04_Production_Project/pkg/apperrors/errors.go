package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorType string

const (
	TypeBadRequest     ErrorType = "BAD_REQUEST"
	TypeNotFound       ErrorType = "NOT_FOUND"
	TypeValidation     ErrorType = "VALIDATION_ERROR"
	TypeUnauthorized   ErrorType = "UNAUTHORIZED"
	TypeConflict       ErrorType = "CONFLICT"
	TypeInternalServer ErrorType = "INTERNAL_SERVER_ERROR"
)

type AppError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
	Code    int       `json:"code"`
	Err     error     `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewBadRequest(message string) *AppError {
	return &AppError{
		Type:    TypeBadRequest,
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorized(message string) *AppError {
	return &AppError{
		Type:    TypeUnauthorized,
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewNotFound(message string) *AppError {
	return &AppError{
		Type:    TypeNotFound,
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewConflict(message string) *AppError {
	return &AppError{
		Type:    TypeConflict,
		Message: message,
		Code:    http.StatusConflict,
	}
}

func NewInternal(err error, message string) *AppError {
	return &AppError{
		Type:    TypeInternalServer,
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     err,
	}
}

func HandleHTTPError(w http.ResponseWriter, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		resp := map[string]interface{}{
			"type":    appErr.Type,
			"message": appErr.Message,
			"code":    appErr.Code,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(appErr.Code)

		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := map[string]interface{}{
		"type":    TypeInternalServer,
		"message": "internal server error",
		"code":    http.StatusInternalServerError,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(resp)
}
