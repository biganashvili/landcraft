// Package Order holds aggregates that combines many entities into a full object
package order

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidUserID = errors.New("a Order has to have an valid user")
	ErrInvalidLandID = errors.New("a Order has to have an valid land")
)

// Order is a aggregate that combines all entities needed to represent a Order
type Order struct {
	id        uuid.UUID
	userID    uuid.UUID
	landID    uuid.UUID
	createdAt time.Time
}

// NewOrder is a factory to create a new Order aggregate
// It will validate that the name is not empty
func NewOrder(userID, landID uuid.UUID) (Order, error) {
	// Validate that the Name is not empty
	if userID == uuid.Nil {
		return Order{}, ErrInvalidUserID
	}
	if landID == uuid.Nil {
		return Order{}, ErrInvalidLandID
	}
	// Create a Order object and initialize all the values to avoid nil pointer exceptions
	return Order{
		id:        uuid.New(),
		userID:    userID,
		landID:    landID,
		createdAt: time.Now(),
	}, nil
}

// GetID returns the orders root entity ID
func (c *Order) GetID() uuid.UUID {
	return c.id
}

// SetName changes the name of the Order
func (c *Order) GetUserID() uuid.UUID {
	return c.userID
}

// SetName changes the name of the Order
func (c *Order) GetLandID() uuid.UUID {
	return c.landID
}
