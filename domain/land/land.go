// Package land
// Land is an aggregate that represents a land.
package land

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrMissingValues is returned when a land is created without a name or description
	ErrMissingValues = errors.New("missing values")
)

// Land is a aggregate that combines Land with a price and quantity
type Land struct {
	ID           uuid.UUID
	Name         string
	OwnerID      *string
	OwnerAddress *string
}

// NewLand will create a new land
// will return error if name of description is empty
func NewLand(name string, price float64) (Land, error) {
	if name == "" {
		return Land{}, ErrMissingValues
	}

	return Land{
		ID:   uuid.New(),
		Name: name,
	}, nil
}

func (p Land) GetID() uuid.UUID {
	return p.ID
}

func (p Land) GetName() string {
	return p.Name
}
