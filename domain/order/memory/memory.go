// Package memory is a in-memory implementation of the order repository
package memory

import (
	"fmt"
	"sync"

	"github.com/biganashvili/landcraft/domain/order"
	"github.com/google/uuid"
)

// MemoryRepository fulfills the OrderRepository interface
type MemoryRepository struct {
	order map[uuid.UUID]order.Order
	sync.Mutex
}

// New is a factory function to generate a new repository of order
func New() *MemoryRepository {
	return &MemoryRepository{
		order: make(map[uuid.UUID]order.Order),
	}
}

// Get finds a order by ID
func (mr *MemoryRepository) Get(id uuid.UUID) (order.Order, error) {
	if order, ok := mr.order[id]; ok {
		return order, nil
	}

	return order.Order{}, order.ErrOrderNotFound
}

// Get finds a order by ID
func (mr *MemoryRepository) List() ([]order.Order, error) {
	order := []order.Order{}
	for _, u := range mr.order {
		order = append(order, u)
	}

	return order, nil
}

// GetByUserID finds a orders by UserID
func (mr *MemoryRepository) GetByUserID(userID uuid.UUID) ([]order.Order, error) {
	orders := []order.Order{}
	for _, o := range mr.order {
		if o.GetUserID() == userID {
			orders = append(orders, o)
		}
	}
	return orders, nil
}

// GetBylandD finds a orders by LandID
func (mr *MemoryRepository) GetByLandID(landID uuid.UUID) ([]order.Order, error) {
	orders := []order.Order{}
	for _, o := range mr.order {
		if o.GetLandID() == landID {
			orders = append(orders, o)
		}
	}
	return orders, nil
}

// Add will add a new order to the repository
func (mr *MemoryRepository) Add(c order.Order) error {
	if mr.order == nil {
		// Saftey check if order is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.order = make(map[uuid.UUID]order.Order)
		mr.Unlock()
	}
	// Make sure Order isn't already in the repository
	if _, ok := mr.order[c.GetID()]; ok {
		return fmt.Errorf("order already exists: %w", order.ErrFailedToAddOrder)
	}
	mr.Lock()
	mr.order[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing order information with the new order information
func (mr *MemoryRepository) Update(c order.Order) error {
	// Make sure Order is in the repository
	if _, ok := mr.order[c.GetID()]; !ok {
		return fmt.Errorf("order does not exist: %w", order.ErrUpdateOrder)
	}
	mr.Lock()
	mr.order[c.GetID()] = c
	mr.Unlock()
	return nil
}
