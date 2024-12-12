package parser

import (
	"errors"
	"testing"
)

func TestParseConsumables(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    int
		expectedErr error
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseConsumables(tt.input)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v for input: %s", tt.expectedErr, err, tt.input)
			}
			if result != tt.expected {
				t.Errorf("expected %d, got %d for input: %s", tt.expected, result, tt.input)
			}
		})
	}
}
