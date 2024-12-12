package service

import (
	"context"
	"testing"

	"github.com/pvdevs/get-starships-stops/internal/domain"
)

// mockStarshipClient implements a mock version of our SWAPI client interface
type mockStarshipClient struct {
	starships []domain.Starship
	err       error
}

func (m *mockStarshipClient) GetStarships(ctx context.Context) ([]domain.Starship, error) {
	return m.starships, m.err
}

// TestCalculateStops verifies the stop calculation logic for different scenarios.
// It tests various cases including:
// - Multiple starships with different speeds and consumables
// - Edge cases like MGLT = 0
// - Long distance calculations
func TestCalculateStops(t *testing.T) {
	// Define test cases using table-driven test pattern
	tests := []struct {
		name          string            // Description of the test case
		distance      int64             // Input distance to travel
		starships     []domain.Starship // Mock starships for testing
		expectedStops map[string]int    // Expected number of stops for each ship
		wantErr       bool              // Whether we expect an error
	}{
		{
			name:     "successful calculation for multiple ships",
			distance: 1000000,
			starships: []domain.Starship{
				{
					Name:        "Millennium Falcon",
					MGLT:        75,
					Consumables: "2 months",
				},
				{
					Name:        "Y-wing",
					MGLT:        80,
					Consumables: "1 week",
				},
			},
			expectedStops: map[string]int{
				"Millennium Falcon": 9,
				"Y-wing":            74,
			},
			wantErr: false,
		},
		{
			name:     "handle ship with MGLT zero",
			distance: 1000000,
			starships: []domain.Starship{
				{
					Name:        "Death Star",
					MGLT:        0,
					Consumables: "3 years",
				},
			},
			expectedStops: map[string]int{
				"Death Star": 0,
			},
			wantErr: false,
		},
		{
			name:     "handle very long distance",
			distance: 10000000,
			starships: []domain.Starship{
				{
					Name:        "X-wing",
					MGLT:        100,
					Consumables: "1 week",
				},
			},
			expectedStops: map[string]int{
				"X-wing": 595, // Updated expected value based on correct calculation
			},
			wantErr: false,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client with test data
			mockClient := &mockStarshipClient{
				starships: tt.starships,
				err:       nil,
			}

			// Initialize calculator with mock client
			calculator := NewCalculator(mockClient)

			// Execute the method being tested
			stops, err := calculator.CalculateStops(context.Background(), tt.distance)

			// Verify error expectations
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateStops() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify results match expectations
			if len(stops) != len(tt.expectedStops) {
				t.Errorf("Expected %d results, got %d", len(tt.expectedStops), len(stops))
			}

			// Check each ship's calculated stops
			for shipName, expectedStops := range tt.expectedStops {
				if gotStops, exists := stops[shipName]; !exists {
					t.Errorf("Missing result for ship %s", shipName)
				} else if gotStops != expectedStops {
					t.Errorf("For ship %s: expected %d stops, got %d", shipName, expectedStops, gotStops)
				}
			}
		})
	}
}
