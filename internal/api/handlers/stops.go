package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/pvdevs/get-starships-stops/internal/api/models"
	"github.com/pvdevs/get-starships-stops/internal/config"
	"github.com/pvdevs/get-starships-stops/internal/parser"
	"github.com/pvdevs/get-starships-stops/internal/service"
	"github.com/pvdevs/get-starships-stops/internal/service/swapi"
)

// StopsHandler holds dependencies for all handlers
type StopsHandler struct {
	calculator service.CalculatorService
}

// NewStopsHandler creates a new handler with required dependencies
func NewStopsHandler(cfg *config.Config) *StopsHandler {
	client := swapi.NewClient(swapi.ClientConfig{
		BaseURL: cfg.SWAPIURL,
	})
	return &StopsHandler{
		calculator: service.NewCalculator(client),
	}
}

// HandleCalculate handles the stop calculation endpoint
func (h *StopsHandler) HandleCalculate(w http.ResponseWriter, r *http.Request) {
	// Method validation
	if r.Method != http.MethodGet {
		models.WriteError(w, http.StatusMethodNotAllowed, "Only GET method is allowed")
		return
	}

	// Get distance from path parameter
	// URL format: /calculate-stops/{distance}
	pathParts := strings.Split(r.URL.Path, "/")

	// If no distance provided, return helpful message
	if len(pathParts) <= 2 || pathParts[2] == "" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.HelpResponse{
			Message: "Please provide a distance in MGLT after /calculate-stops/",
			Example: "/calculate-stops/1000000",
			Usage:   "GET /calculate-stops/{distance}",
		})
		return
	}

	if len(pathParts) != 3 {
		models.WriteError(w, http.StatusBadRequest, "Invalid URL format")
		return
	}

	// Parse distance
	distance, err := parser.ParseDistance(pathParts[2])
	if err != nil {
		models.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Use the handler's calculator instance
	stops, err := h.calculator.CalculateStops(context.Background(), distance)
	if err != nil {
		models.WriteError(w, http.StatusInternalServerError, "Failed to calculate stops")
		return
	}

	// Convert map to slice and sort
	var results []models.Result
	for name, numStops := range stops {
		results = append(results, models.Result{
			Name:  name,
			Stops: numStops,
		})
	}
	models.SortResults(results)

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.StopsResponse{
		Distance: distance,
		Results:  results,
	})
}
