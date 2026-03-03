package postgres

import (
	"context"
	"database/sql"

	"github.com/Atharva0506/trading_bot/internal/models"
)

// SignalPostgresRepo implements repository.SignalRepository using Postgres.
type SignalPostgresRepo struct {
	db *sql.DB
}

// NewSignalPostgresRepo returns a new SignalPostgresRepo.
func NewSignalPostgresRepo(db *sql.DB) *SignalPostgresRepo {
	return &SignalPostgresRepo{
		db: db,
	}
}

// Create inserts a new trade signal into the database.
func (r *SignalPostgresRepo) Create(ctx context.Context, signal *models.Signal) error {
	query := "INSERT INTO signals (id, symbol, action, price, timestamp) VALUES ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, signal.ID, signal.Symbol, signal.Action, signal.Price, signal.Timestamp)
	return err
}

// GetAll retrieves all trade signals ordered by most recent first.
func (r *SignalPostgresRepo) GetAll(ctx context.Context) ([]*models.Signal, error) {
	query := "SELECT id, symbol, action, price, timestamp FROM signals ORDER BY timestamp DESC"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal
	for rows.Next() {
		var s models.Signal
		err = rows.Scan(&s.ID, &s.Symbol, &s.Action, &s.Price, &s.Timestamp)
		if err != nil {
			return nil, err
		}
		signals = append(signals, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return signals, nil
}

// GetAllBySymbol retrieves all trade signals for a specific symbol.
func (r *SignalPostgresRepo) GetAllBySymbol(ctx context.Context, symbol models.Symbol) ([]*models.Signal, error) {
	query := "SELECT id, symbol, action, price, timestamp FROM signals WHERE symbol = $1 ORDER BY timestamp DESC"

	rows, err := r.db.QueryContext(ctx, query, symbol)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var signals []*models.Signal
	for rows.Next() {
		var s models.Signal
		err = rows.Scan(&s.ID, &s.Symbol, &s.Action, &s.Price, &s.Timestamp)
		if err != nil {
			return nil, err
		}
		signals = append(signals, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return signals, nil
}
