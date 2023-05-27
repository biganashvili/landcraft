// Package order holds all the services that connects repositories into a business flow related to ordering lands
package order

import (
	"github.com/biganashvili/landcraft/domain/land"
	landMemRep "github.com/biganashvili/landcraft/domain/land/memory"
	"github.com/biganashvili/landcraft/domain/order"
	orderMemRep "github.com/biganashvili/landcraft/domain/order/memory"
	"github.com/biganashvili/landcraft/domain/user"
	userMemRep "github.com/biganashvili/landcraft/domain/user/memory"
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
func WithMemoryUserRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := userMemRep.New()
	return WithUserRepository(cr)
}

// WithMemoryUserRepository applies a memory user repository to the OrderService
func WithMemoryOrderRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := orderMemRep.New()
	return WithOrderRepository(cr)
}

// WithMemoryLandRepository adds a in memory land repo and adds all input lands
func WithMemoryLandRepository(lands []land.Land) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := landMemRep.New()

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
	l, err := o.lands.GetByID(landID)
	if err != nil {
		return uuid.Nil, err
	}

	newOrder, err := order.NewOrder(u.GetID(), l.GetID())
	if err != nil {
		return uuid.Nil, err
	}
	return newOrder.GetID(), o.orders.Add(newOrder)
}

// AddUser will add a new user and return the userID
func (o *OrderService) AddUser(name string) (uuid.UUID, error) {
	c, err := user.NewUser(name)
	if err != nil {
		return uuid.Nil, err
	}
	// Add to Repo
	err = o.users.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
