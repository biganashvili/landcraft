// Package user holds all the services that connects repositories into a business flow related to usering products
package user

import (
	"github.com/biganashvili/landcraft/domain/user"
	"github.com/biganashvili/landcraft/domain/user/memory"
	"github.com/google/uuid"
)

// UserConfiguration is an alias for a function that will take in a pointer to an UserService and modify it
type UserConfiguration func(os *UserService) error

// UserService is a implementation of the UserService
type UserService struct {
	users user.UserRepository
}

// NewUserService takes a variable amount of UserConfiguration functions and returns a new UserService
// Each UserConfiguration will be called in the user they are passed in
func NewUserService(cfgs ...UserConfiguration) (*UserService, error) {
	// Create the userservice
	os := &UserService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// AddUser will add a new user and return the userID
func (o *UserService) GetUserRepo() user.UserRepository {
	return o.users
}

// WithUserRepository applies a given user repository to the UserService
func WithUserRepository(ur user.UserRepository) UserConfiguration {
	// return a function that matches the UserConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *UserService) error {
		os.users = ur
		return nil
	}
}

func WithMongoUserRepository(connectionString string) UserConfiguration {
	return func(os *UserService) error {

		return nil
	}
}

// WithMemoryUserRepository applies a memory user repository to the UserService
func WithMemoryUserRepository() UserConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := memory.New()
	return WithUserRepository(cr)
}

// AddUser will add a new user and return the userID
func (o *UserService) AddUser(name string) (uuid.UUID, error) {
	c, err := user.NewUser(name)
	if err != nil {
		return uuid.Nil, err
	}
	// Add to Repo
	err = o.users.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}

// List of user ids
func (o *UserService) List() ([]uuid.UUID, error) {
	userIDs := []uuid.UUID{}
	usersList, err := o.users.List()
	if err != nil {
		return userIDs, err
	}

	for _, u := range usersList {
		userIDs = append(userIDs, u.GetID())
	}

	return userIDs, nil
}
