package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "pending"
	PaymentCompleted PaymentStatus = "completed"
	PaymentFailed    PaymentStatus = "failed"
	PaymentRefunded  PaymentStatus = "refunded"
)

type PaymentTransaction struct {
	ID            uuid.UUID       `json:"id"`
	PurchaseID    *uuid.UUID      `json:"purchase_id,omitempty"`
	ItemID        uuid.UUID       `json:"item_id"`
	UserID        uuid.UUID       `json:"user_id"`
	Amount        decimal.Decimal `json:"amount"`
	Provider      string          `json:"provider"`        // "khalti"
	ProviderTxnID string          `json:"provider_txn_id"` // pidx from Khalti
	Status        PaymentStatus   `json:"status"`
	PaymentURL    string          `json:"payment_url"`
	CreatedAt     time.Time       `json:"created_at"`
	CompletedAt   *time.Time      `json:"completed_at,omitempty"`
}

