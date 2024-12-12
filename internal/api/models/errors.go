package models

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a standard error response for the API.
type ErrorResponse struct {
	Error   string `json:"error"`   // HTTP status text (e.g., "Bad Request")
	Code    int    `json:"code"`    // HTTP status code (e.g., 400)
	Message string `json:"message"` // Detailed error message
}

// WriteError sends a JSON-formatted error response to the client.
// Sets the HTTP status code and encodes the ErrorResponse.
//
// Example:
//
//	models.WriteError(w, http.StatusBadRequest, "Invalid distance format")
func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error:   http.StatusText(code),
		Code:    code,
		Message: message,
	})
}
