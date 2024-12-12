package parser

import (
	"errors"
	"testing"
)

type distanceTestCase struct {
	name     string
	input    string
	expected int64
	wantErr  error
}

func TestParseDistance(t *testing.T) {
	tests := []distanceTestCase{
		{
			name:     "Valid input: positive integer",
			input:    "1000000",
			expected: 1000000,
			wantErr:  nil,
		},
		{
			name:     "Valid input: small positive integer",
			input:    "50000",
			expected: 50000,
			wantErr:  nil,
		},
		{
			name:     "Valid input: zero",
			input:    "0",
			expected: 0,
			wantErr:  nil,
		},
		{
			name:     "Invalid input: negative number",
			input:    "-1",
			expected: 0,
			wantErr:  ErrNotPositiveInteger,
		},
		{
			name:     "Invalid input: non-numeric string",
			input:    "abc",
			expected: 0,
			wantErr:  ErrNotPositiveInteger,
		},
		{
			name:     "Invalid input: float value",
			input:    "12.34",
			expected: 0,
			wantErr:  ErrNotPositiveInteger,
		},
		{
			name:     "Invalid input: overflow value",
			input:    "100000000000000000000000000",
			expected: 0,
			wantErr:  ErrInputTooLarge,
		},
		{
			name:     "Invalid input: empty string",
			input:    "",
			expected: 0,
			wantErr:  ErrNotPositiveInteger,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDistance(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error %v, got %v for input: %s", tt.wantErr, err, tt.input)
			}
			if got != tt.expected {
				t.Errorf("expected %d, got %d for input: %s", tt.expected, got, tt.input)
			}
		})
	}
}
