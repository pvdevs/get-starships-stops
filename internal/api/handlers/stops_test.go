package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
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
// - Valid requests with mocked calculation results
// - Invalid request formats or missing fields
// - Error propagation from the calculator service
func TestCalculateStops(t *testing.T) {
	// Define test cases using table-driven test pattern
	tests := []struct {
		name           string         // Description of the test case
		requestBody    interface{}    // Input payload for the HTTP request
		mockStops      map[string]int // Mocked stops data from the calculator
		mockError      error          // Mocked error from the calculator
		expectedStatus int            // Expected HTTP status code
		expectedBody   bool           // Whether a response body is expected
	}{
		{
			name:        "valid request",
			requestBody: models.StopsRequest{Distance: "1000000"},
			mockStops: map[string]int{
				"X-wing": 50,
				"Y-wing": 74,
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   true,
		},
		{
			name:           "empty distance",
			requestBody:    models.StopsRequest{Distance: ""},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid json",
			requestBody:    "invalid",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "invalid distance format",
			requestBody:    models.StopsRequest{Distance: "invalid"},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "calculator error",
			requestBody:    models.StopsRequest{Distance: "1000000"},
			mockError:      fmt.Errorf("calculation error"),
			expectedStatus: http.StatusInternalServerError,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create handler with mock calculator
			h := &Handler{
				calculator: &mockCalculator{
					stops: tt.mockStops,
					err:   tt.mockError,
				},
			}

			// Create request
			var body bytes.Buffer
			if str, ok := tt.requestBody.(string); ok {
				body.WriteString(str) // Directly write string if input is raw JSON
			} else {
				json.NewEncoder(&body).Encode(tt.requestBody) // Encode struct to JSON
			}

			req := httptest.NewRequest(http.MethodPost, "/calculate-stops", &body)
			rec := httptest.NewRecorder()

			// Call handler
			h.CalculateStops(rec, req)

			// Verify HTTP status code
			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// For successful requests, verify response structure
			if tt.expectedBody {
				var response models.StopsResponse
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}

				// Verify distance matches request
				distance, _ := strconv.ParseInt(tt.requestBody.(models.StopsRequest).Distance, 10, 64)
				if response.Distance != distance {
					t.Errorf("expected distance %d, got %d", distance, response.Distance)
				}

				// Verify the number of results matches the mock data
				if len(response.Results) != len(tt.mockStops) {
					t.Errorf("expected %d results, got %d", len(tt.mockStops), len(response.Results))
				}
			}
		})
	}
}
