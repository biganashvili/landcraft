package landcraft

import (
	"github.com/google/uuid"
)

// User is a entity that represents a person in all Domains
type User struct {
	// ID is the identifier of the Entity, the ID is shared for all sub domains
	ID uuid.UUID
	// Name is the name of the User
	Name string
	// Age is the age of the User
	Age int
}
