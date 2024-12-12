package swapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pvdevs/get-starships-stops/internal/domain"
)

// TestClient_GetStarships tests the basic functionality of the GetStarships method.
func TestClient_GetStarships(t *testing.T) {
	tests := []struct {
		name         string
		mockResponse string
		mockStatus   int
		wantErr      bool
		wantCount    int
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatus)
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(tt.mockResponse))
			}))
			defer server.Close()

			client := NewClient(ClientConfig{
				BaseURL: server.URL,
			})
			starships, err := client.GetStarships(context.Background())

			if (err != nil) != tt.wantErr {
				t.Errorf("GetStarships() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && len(starships) != tt.wantCount {
				t.Errorf("GetStarships() got %d starships, want %d", len(starships), tt.wantCount)
			}
		})
	}
}

// TestClient_handlePagination tests the client's pagination handling.
func TestClient_handlePagination(t *testing.T) {
	currentPage := 0
	var responses []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responses[currentPage]))
		currentPage++
	}))
	defer server.Close()

	responses = []string{
		fmt.Sprintf(`{
            "count": 2,
            "next": "%s/api/starships/?page=2",
            "previous": null,
            "results": [{"name":"X-wing","MGLT":"100","consumables":"1 week"}]
        }`, server.URL),
		`{
            "count": 2,
            "next": null,
            "previous": "page1",
            "results": [{"name":"Y-wing","MGLT":"80","consumables":"1 week"}]
        }`,
	}

	client := NewClient(ClientConfig{
		BaseURL: server.URL,
	})
	starships, err := client.GetStarships(context.Background())

	if err != nil {
		t.Errorf("GetStarships() error = %v", err)
		return
	}

	if len(starships) != 2 {
		t.Errorf("Expected 2 starships, got %d", len(starships))
	}

	expectedNames := []string{"X-wing", "Y-wing"}
	for i, ship := range starships {
		if ship.Name != expectedNames[i] {
			t.Errorf("Expected starship name %s, got %s", expectedNames[i], ship.Name)
		}
	}
}

// TestAPIToDomainStarship tests the conversion from API response to domain model
func TestAPIToDomainStarship(t *testing.T) {
	tests := []struct {
		name    string
		input   APIStarship
		want    domain.Starship
		wantErr error
	}{
		{
			name: "valid starship",
			input: APIStarship{
				Name:        "X-wing",
				MGLT:        "100",
				Consumables: "1 week",
			},
			want: domain.Starship{
				Name:        "X-wing",
				MGLT:        100,
				Consumables: "1 week",
			},
			wantErr: nil,
		},
		{
			name: "unknown MGLT",
			input: APIStarship{
				Name:        "Unknown Ship",
				MGLT:        "unknown",
				Consumables: "1 month",
			},
			wantErr: ErrSkipShip,
		},
		{
			name: "n/a MGLT",
			input: APIStarship{
				Name:        "NA Ship",
				MGLT:        "n/a",
				Consumables: "1 month",
			},
			wantErr: ErrSkipShip,
		},
		{
			name: "invalid MGLT number",
			input: APIStarship{
				Name:        "Broken Ship",
				MGLT:        "not a number",
				Consumables: "1 month",
			},
			wantErr: fmt.Errorf("invalid MGLT format"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := apiToDomainStarship(tt.input)

			if tt.wantErr != nil {
				if err == nil {
					t.Error("apiToDomainStarship() expected error, got nil")
					return
				}
				if tt.wantErr == ErrSkipShip && !errors.Is(err, ErrSkipShip) {
					t.Errorf("expected ErrSkipShip, got %v", err)
				}
				return
			}

			if err != nil {
				t.Errorf("apiToDomainStarship() unexpected error: %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("apiToDomainStarship() = %v, want %v", got, tt.want)
			}
		})
	}
}
