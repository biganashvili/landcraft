package handler

import (
	"encoding/json"
	"net/http"

	"github.com/biganashvili/landcraft/cmd/api/config"
	"github.com/biganashvili/landcraft/cmd/api/request"
	"github.com/biganashvili/landcraft/cmd/api/response"
)

func AddOrder(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		response.WrapResponse(w, nil, "Method is not supported.", http.StatusNotFound)
		return
	}

	var req request.AddOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WrapResponse(w, nil, "Bad request", http.StatusBadRequest)
		return
	}

	id, err := config.App.AddOrder(req.UserID, req.LandID)
	if err != nil {
		response.WrapResponse(w, nil, "Order not created", http.StatusNotFound)
		return
	}

	response.WrapResponse(w, response.AddOrderResponse{ID: id}, "", http.StatusCreated)

}

func OrdersList(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		response.WrapResponse(w, nil, "Method is not supported.", http.StatusNotFound)
		return
	}
	var req request.ListOrdersRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WrapResponse(w, nil, "Bad request", http.StatusBadRequest)
		return
	}

	ids, err := config.App.OrderService.List(req.UserID)
	if err != nil {
		response.WrapResponse(w, nil, "Internal Error", http.StatusInternalServerError)
		return
	}
	response.WrapResponse(w, response.OrdersListResponse{Orders: ids}, "", http.StatusOK)

}
