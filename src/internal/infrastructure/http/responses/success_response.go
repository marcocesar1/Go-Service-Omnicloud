package responses

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]any{
		"data": data,
	})
}
