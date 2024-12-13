package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/pvdevs/get-starships-stops/internal/api/models"
)

// mockCalculator implements the calculator interface for testing.
type mockCalculator struct {
	stops map[string]int // Mocked stops result for testing
	err   error          // Mocked error for testing
}

// CalculateStops simulates the calculation logic of the calculator.
func (m *mockCalculator) CalculateStops(ctx context.Context, distance int64) (map[string]int, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.stops, nil
}

// TestCalculateStops verifies the HTTP handler logic for calculating stops.
// It tests various scenarios including:
// - Valid URL paths with correct distance parameters
// - Invalid URL formats and missing parameters
// - Error propagation from the calculator service
func TestCalculateStops(t *testing.T) {
	// Define test cases using table-driven test pattern
	tests := []struct {
		name           string         // Description of the test case
		urlPath        string         // URL path including the distance parameter
		mockStops      map[string]int // Mocked stops data from the calculator
		mockError      error          // Mocked error from the calculator
		expectedStatus int            // Expected HTTP status code
	}{
		{
			name:    "valid request with proper distance",
			urlPath: "/calculate-stops/1000000",
			mockStops: map[string]int{
				"X-wing": 50,
				"Y-wing": 74,
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid distance format",
			urlPath:        "/calculate-stops/invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "calculator service error",
			urlPath:        "/calculate-stops/1000000",
			mockError:      fmt.Errorf("calculation error"),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "malformed URL path",
			urlPath:        "/calculate-stops/1000000/extra",
			expectedStatus: http.StatusBadRequest,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create handler with mock calculator
			h := &StopsHandler{
				calculator: &mockCalculator{
					stops: tt.mockStops,
					err:   tt.mockError,
				},
			}

			// Create request with URL path parameter
			req := httptest.NewRequest(http.MethodGet, tt.urlPath, nil)
			rec := httptest.NewRecorder()

			// Call handler
			h.HandleCalculate(rec, req)

			// Verify HTTP status code
			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// Verify response structure for successful requests
			if rec.Code == http.StatusOK {
				var response models.StopsResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				// Extract distance from URL path for comparison
				pathParts := strings.Split(tt.urlPath, "/")
				expectedDistance, _ := strconv.ParseInt(pathParts[2], 10, 64)
				if response.Distance != expectedDistance {
					t.Errorf("expected distance %d, got %d", expectedDistance, response.Distance)
				}

				// Verify results match mock data
				if len(response.Results) != len(tt.mockStops) {
					t.Errorf("expected %d results, got %d", len(tt.mockStops), len(response.Results))
				}

				// Verify each result matches the mock data
				resultMap := make(map[string]int)
				for _, result := range response.Results {
					resultMap[result.Name] = result.Stops
				}
				for name, stops := range tt.mockStops {
					if resultMap[name] != stops {
						t.Errorf("for ship %s: expected %d stops, got %d", name, stops, resultMap[name])
					}
				}
			}
		})
	}
}
