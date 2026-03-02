package http

import (
	"encoding/json"
	"net/http"

	"github.com/Atharva0506/trading_bot/internal/service"
	"github.com/Atharva0506/trading_bot/pkg/apperrors"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apperrors.HandleHTTPError(w, apperrors.NewBadRequest("invalid request body"))
		return
	}
	user, err := h.service.RegisterUser(r.Context(), req.Email, req.Password)
	if err != nil {
		apperrors.HandleHTTPError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req AuthRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apperrors.HandleHTTPError(w, apperrors.NewBadRequest("invalid request body"))
		return
	}
	tokens, err := h.service.LoginUser(r.Context(), req.Email, req.Password)
	if err != nil {
		apperrors.HandleHTTPError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokens)

}
