package response

import "github.com/google/uuid"

type AddUserResponse struct {
	ID uuid.UUID `json:"id"`
}

type UsersListResponse struct {
	Users []uuid.UUID `json:"users"`
}
