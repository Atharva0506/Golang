package models

import (
	"time"

	"github.com/google/uuid"
)

// ==========================================
// STEP 1: DEFINE YOUR CORE DOMAIN MODEL
// ==========================================
// As a senior engineer, your first job is to define WHAT a User is,
// entirely independent of databases or APIs. This is a pure Go struct.
//
// TODO 1: Create a `User` struct below.
// Think about the fields an Authentication Provider / Trading Platform needs.
// For example:
// - ID (UUID or int64)
// - Email (string)
// - PasswordHash (string) - Never store plain text passwords!
// - Role (string) - e.g., "admin", "trader"
// - Status (string) - e.g., "active", "suspended", "pending_verification"
// - CreatedAt (time.Time)
//

// TODO 2: Add Struct Tags.
// Struct tags tell Go how to marshal this data to JSON (for API responses)
// and how to map it to Database columns.
// Example: `json:"id" db:"id"`
// Note: You should generally avoid returning `PasswordHash` in JSON responses,
// so you would use the tag `json:"-"` to hide it.
type User struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	Email        string     `json:"email" db:"email"`
	PasswordHash string     `json:"-" db:"password"`
	Role         string     `json:"role" db:"role"`
	Status       UserStatus `json:"status" db:"status"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
}

// TODO 3: Create Custom Types (Optional but Recommended)
// Instead of Raw strings for Status or Role, define specific custom types:
// type UserStatus string
// const (
//
//	StatusActive UserStatus = "active"
//	StatusSuspended UserStatus = "suspended"
//
// )
type UserStatus string

const (
	StatusActive    UserStatus = "active"
	StatusSuspended UserStatus = "suspended"
)
