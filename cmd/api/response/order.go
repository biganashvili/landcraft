package response

import "github.com/google/uuid"

type AddOrderResponse struct {
	ID uuid.UUID `json:"id"`
}

type OrdersListResponse struct {
	Orders []uuid.UUID `json:"orders"`
}
