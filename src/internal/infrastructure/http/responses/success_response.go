package responses

import (
	"encoding/json"
	"net/http"
)

type SuccessResponseData struct {
	Data any `json:"data"`
}

func SuccessResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(SuccessResponseData{
		Data: data,
	})
}
