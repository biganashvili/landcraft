package landcraft

import (
	"math/big"
	"time"

	"github.com/google/uuid"
)

// Transaction is a payment between two parties
type Transaction struct {
	amount     big.Int
	customerID uuid.UUID
	landID     uuid.UUID
	address    string
	createdAt  time.Time
}
