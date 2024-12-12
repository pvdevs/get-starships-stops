package swapi

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestClient_GetStarships tests the basic functionality of the GetStarships method.
// It verifies different response scenarios from the SWAPI API including:
// - Successful responses with starship data
// - Server errors
// - Invalid JSON responses
// - Empty result sets
func TestClient_GetStarships(t *testing.T) {
	// Define test cases using table-driven test pattern
	tests := []struct {
		name         string // Description of the test case
		mockResponse string // JSON response the mock server should return
		mockStatus   int    // HTTP status code the mock server should return
		wantErr      bool   // Whether we expect an error
		wantCount    int    // Expected number of starships in the response
	}{
		{
			name:         "successful response",
			mockResponse: `{"count":1,"next":null,"previous":null,"results":[{"name":"X-wing","model":"T-65 X-wing","MGLT":"100","consumables":"1 week"}]}`,
			mockStatus:   http.StatusOK,
			wantErr:      false,
			wantCount:    1,
		},
		{
			name:         "server error",
			mockResponse: `{"detail": "Internal server error"}`,
			mockStatus:   http.StatusInternalServerError,
			wantErr:      true,
			wantCount:    0,
		},
		{
			name:         "invalid json response",
			mockResponse: `{invalid json}`,
			mockStatus:   http.StatusOK,
			wantErr:      true,
			wantCount:    0,
		},
		{
			name:         "empty response",
			mockResponse: `{"count":0,"next":null,"previous":null,"results":[]}`,
			mockStatus:   http.StatusOK,
			wantErr:      false,
			wantCount:    0,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server that returns our mock response
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatus)
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close() // Ensure server is closed after test

			// Initialize client with test server URL
			client := NewClient(server.URL)

			// Execute the method being tested
			starships, err := client.GetStarships(context.Background())

			// Verify error expectations
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStarships() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify result count if no error occurred
			if err == nil && len(starships) != tt.wantCount {
				t.Errorf("GetStarships() got %d starships, want %d", len(starships), tt.wantCount)
			}
		})
	}
}

// TestClient_handlePagination specifically tests the client's ability to handle
// paginated responses from the SWAPI API. It verifies that:
// - The client correctly follows the "next" URL
// - All starships from all pages are collected
// - The starships are returned in the correct order
func TestClient_handlePagination(t *testing.T) {
	currentPage := 0
	var responses []string // Holds our mock responses for each page

	// Create test server that will serve different responses based on the page number
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responses[currentPage]))
		currentPage++ // Increment to serve next page on next request
	}))
	defer server.Close()

	// Define paginated responses after server creation so we can use its URL
	responses = []string{
		// First page with link to second page
		fmt.Sprintf(`{
            "count": 2,
            "next": "%s/api/starships/?page=2",
            "previous": null,
            "results": [{"name":"X-wing","MGLT":"100","consumables":"1 week"}]
        }`, server.URL),
		// Second (final) page
		`{
            "count": 2,
            "next": null,
            "previous": "page1",
            "results": [{"name":"Y-wing","MGLT":"80","consumables":"1 week"}]
        }`,
	}

	client := NewClient(server.URL)
	starships, err := client.GetStarships(context.Background())

	// Verify no errors occurred
	if err != nil {
		t.Errorf("GetStarships() error = %v", err)
		return
	}

	// Verify total number of starships
	if len(starships) != 2 {
		t.Errorf("Expected 2 starships, got %d", len(starships))
	}

	// Verify starships are in correct order
	expectedNames := []string{"X-wing", "Y-wing"}
	for i, ship := range starships {
		if ship.Name != expectedNames[i] {
			t.Errorf("Expected starship name %s, got %s", expectedNames[i], ship.Name)
		}
	}
}
