package dto

import (
	"time"

	"github.com/google/uuid"
)

type TransactionLog struct {
	TransactionID uuid.UUID `json:"transaction_id"`
	Timestamp     time.Time `json:"timestamp"`
	Status        string    `json:"status"`
	Value         float64   `json:"value"`
}
