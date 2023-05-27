// Package order holds all the services that connects repositories into a business flow related to ordering lands
package order

import (
	"errors"

	"github.com/biganashvili/landcraft/domain/land"
	landMemRepo "github.com/biganashvili/landcraft/domain/land/memory"
	"github.com/biganashvili/landcraft/domain/order"
	orderMemRepo "github.com/biganashvili/landcraft/domain/order/memory"
	"github.com/biganashvili/landcraft/domain/user"
	userMemRepo "github.com/biganashvili/landcraft/domain/user/memory"
	"github.com/google/uuid"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type OrderConfiguration func(os *OrderService) error

// OrderService is a implementation of the OrderService
type OrderService struct {
	users  user.UserRepository
	lands  land.LandRepository
	orders order.OrderRepository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	// Create the orderservice
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithUserRepository applies a given user repository to the OrderService
func WithUserRepository(cr user.UserRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.users = cr
		return nil
	}
}

// WithOrderRepository applies a given user repository to the OrderService
func WithOrderRepository(cr order.OrderRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.orders = cr
		return nil
	}
}

// WithMemoryUserRepository applies a memory user repository to the OrderService
func WithMemoryUserRepository(rep user.UserRepository) OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := rep
	if rep == nil {
		cr = userMemRepo.New()
	}

	return WithUserRepository(cr)
}

// WithMemoryUserRepository applies a memory user repository to the OrderService
func WithMemoryOrderRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := orderMemRepo.New()
	return WithOrderRepository(cr)
}

// WithMemoryLandRepository adds a in memory land repo and adds all input lands
func WithMemoryLandRepository(lands []land.Land) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := landMemRepo.New()

		// Add Items to repo
		for _, p := range lands {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.lands = pr
		return nil
	}
}

// CreateOrder will chaintogether all repositories to create a order for a user
// will return the collected price of all Lands
func (o *OrderService) CreateOrder(userID uuid.UUID, landID uuid.UUID) (uuid.UUID, error) {
	// Get the user
	u, err := o.users.Get(userID)
	if err != nil {
		return uuid.Nil, err
	}
	if u.GetID() == uuid.Nil {
		return uuid.Nil, errors.New("user not found")
	}

	l, err := o.lands.GetByID(landID)
	if err != nil {
		return uuid.Nil, err
	}
	if l.GetID() == uuid.Nil {
		return uuid.Nil, errors.New("land not found")
	}

	newOrder, err := order.NewOrder(u.GetID(), l.GetID())
	if err != nil {
		return uuid.Nil, err
	}
	err = o.orders.Add(newOrder)
	if err != nil {
		return uuid.Nil, err
	}
	return newOrder.GetID(), nil
}

// AddUser will add a new user and return the userID
func (o *OrderService) List(userID uuid.UUID) ([]uuid.UUID, error) {
	orderIDs := []uuid.UUID{}
	ords, err := o.orders.GetByUserID(userID)
	if err != nil {
		return orderIDs, err
	}
	for _, ord := range ords {
		orderIDs = append(orderIDs, ord.GetID())
	}
	return orderIDs, nil
}
