package handler

import (
	"encoding/json"
	"net/http"

	"github.com/biganashvili/landcraft/cmd/api/config"
	"github.com/biganashvili/landcraft/cmd/api/request"
	"github.com/biganashvili/landcraft/cmd/api/response"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {

		response.WrapResponse(w, nil, "Method is not supported.", http.StatusNotFound)
		return
	}

	var req request.AddUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WrapResponse(w, nil, "Bad request", http.StatusBadRequest)
		return
	}

	id, err := config.App.UserService.AddUser(req.Name)
	if err != nil {
		response.WrapResponse(w, nil, "User not found", http.StatusNotFound)
	}

	response.WrapResponse(w, response.AddUserResponse{ID: id}, "", http.StatusCreated)

}

func UsersList(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		response.WrapResponse(w, nil, "Method is not supported.", http.StatusNotFound)
		return
	}

	ids, err := config.App.UserService.List()
	if err != nil {
		response.WrapResponse(w, nil, "Internal Error", http.StatusInternalServerError)
	}
	response.WrapResponse(w, response.UsersListResponse{Users: ids}, "", http.StatusOK)

}
