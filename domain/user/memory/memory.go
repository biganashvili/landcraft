// Package memory is a in-memory implementation of the user repository
package memory

import (
	"fmt"
	"sync"

	"github.com/biganashvili/landcraft/domain/user"
	"github.com/google/uuid"
)

// MemoryRepository fulfills the UserRepository interface
type MemoryRepository struct {
	users map[uuid.UUID]user.User
	sync.Mutex
}

// New is a factory function to generate a new repository of users
func New() *MemoryRepository {
	return &MemoryRepository{
		users: make(map[uuid.UUID]user.User),
	}
}

// Get finds a user by ID
func (mr *MemoryRepository) Get(id uuid.UUID) (user.User, error) {
	if user, ok := mr.users[id]; ok {
		return user, nil
	}

	return user.User{}, user.ErrUserNotFound
}

// Get finds a user by ID
func (mr *MemoryRepository) List() ([]user.User, error) {
	users := []user.User{}
	for _, u := range mr.users {
		users = append(users, u)
	}

	return users, nil
}

// Add will add a new user to the repository
func (mr *MemoryRepository) Add(c user.User) error {
	if mr.users == nil {
		// Saftey check if users is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.users = make(map[uuid.UUID]user.User)
		mr.Unlock()
	}
	// Make sure User isn't already in the repository
	if _, ok := mr.users[c.GetID()]; ok {
		return fmt.Errorf("user already exists: %w", user.ErrFailedToAddUser)
	}
	mr.Lock()
	mr.users[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing user information with the new user information
func (mr *MemoryRepository) Update(c user.User) error {
	// Make sure User is in the repository
	if _, ok := mr.users[c.GetID()]; !ok {
		return fmt.Errorf("user does not exist: %w", user.ErrUpdateUser)
	}
	mr.Lock()
	mr.users[c.GetID()] = c
	mr.Unlock()
	return nil
}
