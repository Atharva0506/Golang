package service

import (
	"context"
	"time"

	"github.com/Atharva0506/trading_bot/internal/models"
	"github.com/Atharva0506/trading_bot/internal/repository"
	"github.com/Atharva0506/trading_bot/pkg/apperrors"
	"github.com/google/uuid"
)

// SignalService handles business logic for trade signals.
type SignalService struct {
	repo repository.SignalRepository
}

// NewSignalService returns a new SignalService.
func NewSignalService(repo repository.SignalRepository) *SignalService {
	return &SignalService{
		repo: repo,
	}
}

// CreateSignal validates the input and persists a new trade signal.
func (s *SignalService) CreateSignal(ctx context.Context, symbol, action string, price uint64) (*models.Signal, error) {
	sym := models.Symbol(symbol)
	if sym != models.SOL && sym != models.ETH {
		return nil, apperrors.NewBadRequest("invalid symbol: must be SOL or ETH")
	}
	act := models.Action(action)
	if act != models.ActionBuy && act != models.ActionSell {
		return nil, apperrors.NewBadRequest("invalid action: must be buy or sell")
	}
	if price == 0 {
		return nil, apperrors.NewBadRequest("price must be greater than 0")
	}

	signal := models.Signal{
		ID:        uuid.New(),
		Symbol:    sym,
		Action:    act,
		Price:     price,
		Timestamp: time.Now(),
	}
	err := s.repo.Create(ctx, &signal)
	if err != nil {
		return nil, apperrors.NewInternal(err, "failed to create signal")
	}
	return &signal, nil
}

// GetAllSignals retrieves all trade signals.
func (s *SignalService) GetAllSignals(ctx context.Context) ([]*models.Signal, error) {
	signals, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, apperrors.NewInternal(err, "failed to fetch signals")
	}
	return signals, nil
}

// GetSignalsBySymbol retrieves all trade signals for a given symbol.
func (s *SignalService) GetSignalsBySymbol(ctx context.Context, symbol string) ([]*models.Signal, error) {
	sym := models.Symbol(symbol)
	if sym != models.SOL && sym != models.ETH {
		return nil, apperrors.NewBadRequest("invalid symbol: must be SOL or ETH")
	}
	signals, err := s.repo.GetAllBySymbol(ctx, sym)
	if err != nil {
		return nil, apperrors.NewInternal(err, "failed to fetch signals")
	}
	return signals, nil
}
