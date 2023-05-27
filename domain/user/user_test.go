package user_test

import (
	"testing"

	"github.com/biganashvili/landcraft/domain/user"
)

func TestUser_NewUser(t *testing.T) {
	// Build our needed testcase data struct
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	// Create new test cases
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: user.ErrInvalidPerson,
		}, {
			test:        "Valid Name",
			name:        "Sergi Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run Tests
		t.Run(tc.test, func(t *testing.T) {
			// Create a new user
			_, err := user.NewUser(tc.name)
			// Check if the error matches the expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
