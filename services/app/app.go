package app

import (
	"github.com/biganashvili/landcraft/services/order"
	"github.com/biganashvili/landcraft/services/user"
	"github.com/google/uuid"
)

// LandcraftConfiguration is an alias that takes a pointer and modifies the Landcraft
type LandcraftConfiguration func(os *Landcraft) error

type Landcraft struct {
	// orderservice is used to handle orders
	OrderService *order.OrderService
	UserService  *user.UserService
}

// NewLandcraft takes a variable amount of LandcraftConfigurations and builds a Landcraft
func NewLandcraft(cfgs ...LandcraftConfiguration) (*Landcraft, error) {
	// Create the Landcraft
	t := &Landcraft{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// WithOrderService applies a given OrderService to the Landcraft
func WithOrderService(os *order.OrderService) LandcraftConfiguration {
	// return a function that matches the LandcraftConfiguration signature
	return func(t *Landcraft) error {
		t.OrderService = os
		return nil
	}
}

// WithOrderService applies a given OrderService to the Landcraft
func WithUserService(us *user.UserService) LandcraftConfiguration {
	// return a function that matches the LandcraftConfiguration signature
	return func(t *Landcraft) error {
		t.UserService = us
		return nil
	}
}

// Order performs an order for a customer
func (t *Landcraft) AddOrder(userID uuid.UUID, landID uuid.UUID) (uuid.UUID, error) {
	return t.OrderService.CreateOrder(userID, landID)
}
