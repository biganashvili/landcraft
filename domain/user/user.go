// Package User holds aggregates that combines many entities into a full object
package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidPerson = errors.New("a User has to have an valid person")
)

// User is a aggregate that combines all entities needed to represent a User
type User struct {
	ID uuid.UUID
	// Name is the name of the User
	Name string
}

// NewUser is a factory to create a new User aggregate
// It will validate that the name is not empty
func NewUser(name string) (User, error) {
	// Validate that the Name is not empty
	if name == "" {
		return User{}, ErrInvalidPerson
	}
	// Create a User object and initialize all the values to avoid nil pointer exceptions
	return User{
		Name: name,
		ID:   uuid.New(),
	}, nil
}

// GetID returns the users root entity ID
func (c *User) GetID() uuid.UUID {
	return c.ID
}

// SetName changes the name of the User
func (c *User) SetName(name string) {
	c.Name = name
}

// SetName changes the name of the User
func (c *User) GetName() string {
	return c.Name
}
