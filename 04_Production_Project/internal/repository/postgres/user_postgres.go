package postgres

import (
	"context"
	"database/sql"

	"github.com/Atharva0506/trading_bot/internal/models"
	"github.com/google/uuid"
)

type UserPostgresRepo struct {
	db *sql.DB
}

func NewUserPostgresRepo(db *sql.DB) *UserPostgresRepo {
	return &UserPostgresRepo{
		db: db,
	}
}

func (r *UserPostgresRepo) Create(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (id, email, password, role, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.PasswordHash, user.Role, user.Status, user.CreatedAt)
	return err
}

func (r *UserPostgresRepo) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := "SELECT id, email, password, role, status, created_at FROM users WHERE id = $1"
	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.Status, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgresRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := "SELECT id, email, password, role, status, created_at FROM users WHERE email = $1"
	var user models.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role, &user.Status, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
