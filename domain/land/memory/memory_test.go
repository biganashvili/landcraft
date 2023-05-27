package memory

import (
	"testing"

	"github.com/biganashvili/landcraft/domain/land"
	"github.com/google/uuid"
)

func TestMemoryLandRepository_Add(t *testing.T) {
	repo := New()
	land, err := land.NewLand("Beer", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(land)
	if len(repo.lands) != 1 {
		t.Errorf("Expected 1 land, got %d", len(repo.lands))
	}
}
func TestMemoryLandRepository_Get(t *testing.T) {
	repo := New()
	existingProd, err := land.NewLand("Beer", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProd)
	if len(repo.lands) != 1 {
		t.Errorf("Expected 1 land, got %d", len(repo.lands))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get land by id",
			id:          existingProd.GetID(),
			expectedErr: nil,
		}, {
			name:        "Get non-existing land by id",
			id:          uuid.New(),
			expectedErr: land.ErrLandNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
func TestMemoryLandRepository_Delete(t *testing.T) {
	repo := New()
	existingProd, err := land.NewLand("Beer", 1.99)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingProd)
	if len(repo.lands) != 1 {
		t.Errorf("Expected 1 land, got %d", len(repo.lands))
	}

	err = repo.Delete(existingProd.GetID())
	if err != nil {
		t.Error(err)
	}
	if len(repo.lands) != 0 {
		t.Errorf("Expected 0 lands, got %d", len(repo.lands))
	}
}
