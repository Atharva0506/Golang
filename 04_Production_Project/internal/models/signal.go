package models

import (
	"time"

	"github.com/google/uuid"
)

// Symbol represents a tradable asset.
type Symbol string

const (
	SOL Symbol = "SOL"
	ETH Symbol = "ETH"
)

// Action represents whether to buy or sell.
type Action string

const (
	ActionBuy  Action = "buy"
	ActionSell Action = "sell"
)

type Signal struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Symbol    Symbol    `json:"symbol" db:"symbol"`
	Action    Action    `json:"action" db:"action"`
	Price     uint64    `json:"price" db:"price"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}
