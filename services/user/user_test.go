package user

import (
	"testing"

	"github.com/biganashvili/landcraft/domain/user"
)

func TestUser_NewUserService(t *testing.T) {

	us, err := NewUserService(
		WithMemoryUserRepository(),
	)

	if err != nil {
		t.Error(err)
	}

	// Add User
	cust, err := user.NewUser("Sergi")
	if err != nil {
		t.Error(err)
	}

	err = us.users.Add(cust)
	if err != nil {
		t.Error(err)
	}
}
