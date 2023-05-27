// Package Order holds all the domain logic for the Order domain.
package order

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrOrderNotFound is returned when a Order is not found.
	ErrOrderNotFound = errors.New("the Order was not found in the repository")
	// ErrFailedToAddOrder is returned when the Order could not be added to the repository.
	ErrFailedToAddOrder = errors.New("failed to add the Order to the repository")
	// ErrUpdateOrder is returned when the Order could not be updated in the repository.
	ErrUpdateOrder = errors.New("failed to update the Order in the repository")
)

// OrderRepository is a interface that defines the rules around what a Order repository
// Has to be able to perform
type OrderRepository interface {
	Get(uuid.UUID) (Order, error)
	List() ([]Order, error)
	Add(Order) error
	Update(Order) error
	GetByUserID(uuid.UUID) ([]Order, error)
	GetByLandID(uuid.UUID) ([]Order, error)
}
