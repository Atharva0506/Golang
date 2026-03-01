package models

import (
	"time"

	"github.com/google/uuid"
)

// Status represents the delivery status of a notification.
type Status string

const (
	StatusPending Status = "PENDING"
	StatusSent    Status = "SENT"
	StatusFailed  Status = "FAILED"
)

type Notification struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	Type      string    `json:"message_type" db:"message_type"`
	Status    Status    `json:"notification_status" db:"notification_status"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}
