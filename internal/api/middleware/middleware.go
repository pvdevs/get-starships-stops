package middleware

import "net/http"

// Common applies common headers to all HTTP responses.
// Sets the "Content-Type" header to "application/json".
//
// Usage:
//
//	http.HandleFunc("/route", middleware.Common(handler))
func Common(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
