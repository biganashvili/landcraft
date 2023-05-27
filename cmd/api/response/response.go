package response

import (
	"encoding/json"
	"net/http"
)

type GeneralResponse struct {
	Data  any `json:"data"`
	Error any `json:"error"`
}

func WrapResponse(w http.ResponseWriter, data any, errorMessage string, statusCode int) error {
	res := GeneralResponse{Data: data, Error: errorMessage}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if errorMessage == "" {
		res.Error = nil
	}
	return json.NewEncoder(w).Encode(res)
}
