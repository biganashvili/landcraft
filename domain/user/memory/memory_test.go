package memory

import (
	"testing"

	"github.com/biganashvili/landcraft/domain/user"
	"github.com/google/uuid"
)

func TestMemory_GetUser(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// Create a fake user to add to repository
	cust, err := user.NewUser("Sergi")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	// Create the repo to use, and add some test Data to it for testing
	// Skip Factory for this
	repo := MemoryRepository{
		users: map[uuid.UUID]user.User{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "No User By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: user.ErrUserNotFound,
		}, {
			name:        "User By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestMemory_AddUser(t *testing.T) {
	type testCase struct {
		name        string
		cust        string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add User",
			cust:        "Sergi",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				users: map[uuid.UUID]user.User{},
			}

			cust, err := user.NewUser(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}
