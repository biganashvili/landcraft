package request

import "github.com/google/uuid"

type AddOrderRequest struct {
	UserID uuid.UUID `json:"user_id"`
	LandID uuid.UUID `json:"land_id"`
}

type ListOrdersRequest struct {
	UserID uuid.UUID `json:"user_id"`
}
