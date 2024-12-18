package parser

import (
	"errors"
	"fmt"
	"testing"
)

// TestParseConsumables tests the ParseConsumables function.
// It verifies the conversion of various consumables strings to hours, including:
// - Valid input strings with units like "years", "months", "weeks", "days"
// - Invalid input strings, such as empty strings, wrong formats, or unsupported formats
func TestParseConsumables(t *testing.T) {
	// Define test cases using table-driven test pattern
	tests := []struct {
		name        string // Description of the test case
		input       string // Consumables string to parse
		expected    int    // Expected result in hours
		expectedErr error  // Expected error
	}{
		{
			name:        "Valid input: 3 years",
			input:       "3 years",
			expected:    26280, // 3 * 365 * 24 hours
			expectedErr: nil,
		},
		{
			name:        "Valid input: 1 year",
			input:       "1 year",
			expected:    8760, // 365 * 24 hours
			expectedErr: nil,
		},
		{
			name:        "Valid input: 2 months",
			input:       "2 months",
			expected:    1440, // 2 * 30 * 24 hours
			expectedErr: nil,
		},
		{
			name:        "Valid input: 1 week",
			input:       "1 week",
			expected:    168, // 7 * 24 hours
			expectedErr: nil,
		},
		{
			name:        "Valid input: 6 days",
			input:       "6 days",
			expected:    144, // 6 * 24 hours
			expectedErr: nil,
		},
		{
			name:        "Invalid input: empty string",
			input:       "",
			expected:    0,
			expectedErr: ErrEmptyConsumables,
		},
		{
			name:        "Invalid input: wrong format",
			input:       "invalid",
			expected:    0,
			expectedErr: ErrInvalidConsumables,
		},
		{
			name:        "Invalid input: numbers only",
			input:       "123",
			expected:    0,
			expectedErr: ErrInvalidConsumables,
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseConsumables(tt.input)

			// Verify error expectations
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v for input: %s", tt.expectedErr, err, tt.input)
			}

			// Verify result
			if result != tt.expected {
				t.Errorf("expected %d, got %d for input: %s", tt.expected, result, tt.input)
			}
		})
	}
}

func TestBasicCalculation(t *testing.T) {
	// Basic example from SWAPI:
	// X-wing with MGLT: 100, consumables: "1 week"
	// For distance 1000000

	consumables := "1 week"
	hours, err := ParseConsumables(consumables)
	if err != nil {
		t.Fatalf("ParseConsumables error: %v", err)
	}

	// Hours in a week: 7 * 24 = 168
	if hours != 168 {
		t.Errorf("Expected 168 hours, got %d", hours)
	}

	// X-wing travels 100 MGLT per hour
	// So in 168 hours it can travel: 100 * 168 = 16800 MGLT

	mglt := 100
	maxDistance := mglt * hours // should be 16800

	// For total distance 1000000:
	distance := 1000000
	stops := distance / maxDistance // should be ~59.52, rounded to 59

	fmt.Printf("Test calculation:\n")
	fmt.Printf("Hours: %d\n", hours)
	fmt.Printf("Max distance: %d\n", maxDistance)
	fmt.Printf("Stops needed: %d\n", stops)
}
