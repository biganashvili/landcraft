// Package memory is a in memory implementation of the LandRepository interface.
package memory

import (
	"sync"

	"github.com/biganashvili/landcraft/domain/land"
	"github.com/google/uuid"
)

type MemoryLandRepository struct {
	lands map[uuid.UUID]land.Land
	sync.Mutex
}

// New is a factory function to generate a new repository of customers
func New() *MemoryLandRepository {
	return &MemoryLandRepository{
		lands: make(map[uuid.UUID]land.Land),
	}
}

// GetAll returns all lands as a slice
// Yes, it never returns an error, but
// A database implementation could return an error for instance
func (mpr *MemoryLandRepository) GetAll() ([]land.Land, error) {
	// Collect all Lands from map
	var lands []land.Land
	for _, land := range mpr.lands {
		lands = append(lands, land)
	}
	return lands, nil
}

// GetByID searches for a land based on it's ID
func (mpr *MemoryLandRepository) GetByID(id uuid.UUID) (land.Land, error) {
	if land, ok := mpr.lands[uuid.UUID(id)]; ok {
		return land, nil
	}
	return land.Land{}, land.ErrLandNotFound
}

// Add will add a new land to the repository
func (mpr *MemoryLandRepository) Add(newprod land.Land) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.lands[newprod.GetID()]; ok {
		return land.ErrLandAlreadyExist
	}

	mpr.lands[newprod.GetID()] = newprod

	return nil
}

// Update will change all values for a land based on it's ID
func (mpr *MemoryLandRepository) Update(upprod land.Land) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.lands[upprod.GetID()]; !ok {
		return land.ErrLandNotFound
	}

	mpr.lands[upprod.GetID()] = upprod
	return nil
}

// Delete remove an land from the repository
func (mpr *MemoryLandRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.lands[id]; !ok {
		return land.ErrLandNotFound
	}
	delete(mpr.lands, id)
	return nil
}
