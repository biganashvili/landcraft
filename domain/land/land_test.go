package land_test

import (
	"testing"

	"github.com/biganashvili/landcraft/domain/land"
)

func TestLand_NewLand(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "should return error if name is empty",
			name:        "",
			expectedErr: land.ErrMissingValues,
		},
		{
			test:        "validvalues",
			name:        "test",
			description: "test",
			price:       1.0,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := land.NewLand(tc.name, tc.price)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
