package logic

import (
	"errors"
	"testing"
)

func TestParseDistance(t *testing.T) {
	tests := []struct {
		name        string // Descriptive name for the test case
		input       string
		expected    int64
		expectedErr error
	}{
		// Valid cases
		{"Valid input: positive integer", "1000000", 1000000, nil},
		{"Valid input: small positive integer", "50000", 50000, nil},
		{"Valid input: zero", "0", 0, nil},

		// Invalid cases
		{"Invalid input: negative number", "-1", 0, ErrNotPositiveInteger},
		{"Invalid input: non-numeric string", "abc", 0, ErrNotPositiveInteger},
		{"Invalid input: float value", "12.34", 0, ErrNotPositiveInteger},
		{"Invalid input: overflow value", "100000000000000000000000000", 0, ErrInputTooLarge},
		{"Invalid input: empty string", "", 0, ErrNotPositiveInteger},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ParseDistance(test.input)

			if !errors.Is(err, test.expectedErr) {
				t.Errorf("expected error %v, got %v for input: %s", test.expectedErr, err, test.input)
			}
			if result != test.expected {
				t.Errorf("expected %d, got %d for input: %s", test.expected, result, test.input)
			}
		})
	}
}
