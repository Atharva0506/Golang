package postgres

import (
	"context"
	"database/sql"

	"github.com/Atharva0506/trading_bot/internal/models"
	"github.com/google/uuid"
)

// NotificationPostgresRepo implements repository.NotificationRepository using Postgres.
type NotificationPostgresRepo struct {
	db *sql.DB
}

// NewNotificationPostgresRepo returns a new NotificationPostgresRepo.
func NewNotificationPostgresRepo(db *sql.DB) *NotificationPostgresRepo {
	return &NotificationPostgresRepo{db: db}
}

// Create inserts a new notification into the database.
func (r *NotificationPostgresRepo) Create(ctx context.Context, n *models.Notification) error {
	query := `INSERT INTO notifications (id, user_id, message_type, notification_status, timestamp)
	          VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, n.ID, n.UserID, n.Type, n.Status, n.Timestamp)
	return err
}

// UpdateStatus updates the delivery status of a notification identified by id.
func (r *NotificationPostgresRepo) UpdateStatus(ctx context.Context, id uuid.UUID, status models.Status) error {
	query := `UPDATE notifications SET notification_status = $1 WHERE id = $2`
	result, err := r.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

// GetAll retrieves all notifications ordered by most recent first.
func (r *NotificationPostgresRepo) GetAll(ctx context.Context) ([]*models.Notification, error) {
	query := `SELECT id, user_id, message_type, notification_status, timestamp
	          FROM notifications ORDER BY timestamp DESC`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanNotifications(rows)
}

// GetByStatus retrieves all notifications with a specific delivery status.
func (r *NotificationPostgresRepo) GetByStatus(ctx context.Context, status models.Status) ([]*models.Notification, error) {
	query := `SELECT id, user_id, message_type, notification_status, timestamp
	          FROM notifications WHERE notification_status = $1 ORDER BY timestamp ASC`

	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return scanNotifications(rows)
}

// scanNotifications is a helper that maps a *sql.Rows cursor into a slice of Notifications.
func scanNotifications(rows *sql.Rows) ([]*models.Notification, error) {
	var notifications []*models.Notification
	for rows.Next() {
		var n models.Notification
		if err := rows.Scan(&n.ID, &n.UserID, &n.Type, &n.Status, &n.Timestamp); err != nil {
			return nil, err
		}
		notifications = append(notifications, &n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notifications, nil
}
