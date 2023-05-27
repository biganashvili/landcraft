// Package User holds all the domain logic for the User domain.
package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrUserNotFound is returned when a User is not found.
	ErrUserNotFound = errors.New("the User was not found in the repository")
	// ErrFailedToAddUser is returned when the User could not be added to the repository.
	ErrFailedToAddUser = errors.New("failed to add the User to the repository")
	// ErrUpdateUser is returned when the User could not be updated in the repository.
	ErrUpdateUser = errors.New("failed to update the User in the repository")
)

// UserRepository is a interface that defines the rules around what a User repository
// Has to be able to perform
type UserRepository interface {
	Get(uuid.UUID) (User, error)
	List() ([]User, error)
	Add(User) error
	Update(User) error
}
