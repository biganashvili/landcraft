// Package land holds the repository and the implementations for a LandRepository
package land

import (
	"errors"

	"github.com/google/uuid"
)

var (
	//ErrLandNotFound is returned when a land is not found
	ErrLandNotFound = errors.New("the land was not found")
	//ErrLandAlreadyExist is returned when trying to add a land that already exists
	ErrLandAlreadyExist = errors.New("the land already exists")
)

// LandRepository is the repository interface to fulfill to use the land aggregate
type LandRepository interface {
	GetAll() ([]Land, error)
	GetByID(uuid.UUID) (Land, error)
	Add(Land) error
	Update(Land) error
	Delete(uuid.UUID) error
}
