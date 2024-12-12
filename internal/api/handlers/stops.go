package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pvdevs/get-starships-stops/internal/api/models"
	"github.com/pvdevs/get-starships-stops/internal/config"
	"github.com/pvdevs/get-starships-stops/internal/parser"
	"github.com/pvdevs/get-starships-stops/internal/service"
	"github.com/pvdevs/get-starships-stops/internal/service/swapi"
)

// Handler holds dependencies for all handlers
type Handler struct {
	calculator service.CalculatorService
}

// NewHandler creates a new handler with required dependencies
func NewHandler(cfg *config.Config) *Handler {
	client := swapi.NewClient(swapi.ClientConfig{
		BaseURL: cfg.SWAPIURL,
	})
	return &Handler{
		calculator: service.NewCalculator(client),
	}
}

// CalculateStops handles the stop calculation endpoint
func (h *Handler) CalculateStops(w http.ResponseWriter, r *http.Request) {
	// Method validation
	if r.Method != http.MethodPost {
		models.WriteError(w, http.StatusMethodNotAllowed, "Only POST method is allowed")
		return
	}

	// Parse request body
	var req models.StopsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		models.WriteError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	// Validate request
	if req.Distance == "" {
		models.WriteError(w, http.StatusBadRequest, "Distance is required")
		return
	}

	// Parse distance
	distance, err := parser.ParseDistance(req.Distance)
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
