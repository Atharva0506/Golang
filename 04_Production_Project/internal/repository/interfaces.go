package repository

import (
	"context"

	"github.com/Atharva0506/trading_bot/internal/models"
	"github.com/google/uuid"
)

// UserRepository defines the persistence interactions for a User.
type UserRepository interface {
	Create(ctx context.Context, defaultUser *models.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)    // ID should be fully capitalized. Must return the model.
	GetByEmail(ctx context.Context, email string) (*models.User, error) // Must return the model.
}

// SignalRepository defines the persistence interactions for Trade Signals.
type SignalRepository interface {
	Create(ctx context.Context, signal *models.Signal) error
	GetAll(ctx context.Context) ([]*models.Signal, error) // Must return a slice of models.
	GetAllBySymbol(ctx context.Context, symbol models.Symbol) ([]*models.Signal, error)
}

// NotificationRepository defines the persistence interactions for Notifications.
type NotificationRepository interface {
	Create(ctx context.Context, notification *models.Notification) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status models.Status) error // Name it UpdateStatus for clarity
	GetAll(ctx context.Context) ([]*models.Notification, error)                 // Return data and error
	GetByStatus(ctx context.Context, status models.Status) ([]*models.Notification, error)
}
