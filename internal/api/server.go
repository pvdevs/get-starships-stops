package server

import (
	"net/http"

	"github.com/pvdevs/get-starships-stops/internal/api/handlers"
	"github.com/pvdevs/get-starships-stops/internal/api/middleware"
	"github.com/pvdevs/get-starships-stops/internal/config"
)

// NewServer creates and configures an HTTP server with routes and middleware.
func NewServer(cfg *config.Config) *http.Server {
	mux := http.NewServeMux()

	handler := handlers.NewStopsHandler(cfg)

	// Register routes with middleware
	mux.HandleFunc("/calculate-stops/", middleware.Common(handler.HandleCalculate))

	return &http.Server{
		Addr:    cfg.Port,
		Handler: mux,
	}
}
