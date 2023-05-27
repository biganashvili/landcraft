package landcraft

import "github.com/google/uuid"

// Land represents a Item for all sub domains
type Land struct {
	ID           uuid.UUID
	Name         string
	OwnerID      *string
	OwnerAddress *string
}
