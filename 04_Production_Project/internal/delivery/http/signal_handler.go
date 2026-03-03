package http

import (
	"encoding/json"
	"net/http"

	"github.com/Atharva0506/trading_bot/internal/service"
	"github.com/Atharva0506/trading_bot/pkg/apperrors"
	"github.com/go-chi/chi/v5"
)

// CreateSignalRequest represents the JSON body for creating a signal.
type CreateSignalRequest struct {
	Symbol string `json:"symbol"`
	Action string `json:"action"`
	Price  uint64 `json:"price"`
}

// SignalHandler handles HTTP requests for trade signals.
type SignalHandler struct {
	service *service.SignalService
}

// NewSignalHandler returns a new SignalHandler.
func NewSignalHandler(service *service.SignalService) *SignalHandler {
	return &SignalHandler{
		service: service,
	}
}

// CreateSignal handles POST /signals — creates a new trade signal.
func (h *SignalHandler) CreateSignal(w http.ResponseWriter, r *http.Request) {
	var req CreateSignalRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		apperrors.HandleHTTPError(w, apperrors.NewBadRequest("invalid request body"))
		return
	}
	signal, err := h.service.CreateSignal(r.Context(), req.Symbol, req.Action, req.Price)
	if err != nil {
		apperrors.HandleHTTPError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signal)
}

// GetAllSignals handles GET /signals — retrieves all trade signals.
func (h *SignalHandler) GetAllSignals(w http.ResponseWriter, r *http.Request) {
	signals, err := h.service.GetAllSignals(r.Context())
	if err != nil {
		apperrors.HandleHTTPError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(signals)
}

// GetSignalsBySymbol handles GET /signals/{symbol} — retrieves signals filtered by symbol.
func (h *SignalHandler) GetSignalsBySymbol(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")

	signals, err := h.service.GetSignalsBySymbol(r.Context(), symbol)
	if err != nil {
		apperrors.HandleHTTPError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(signals)
}
