package logic

import "testing"

func TestParseDistance(t *testing.T) {
	tests := []struct {
		input       string
		expected    int64
		expectedErr error
	}{
		// Valid cases
		{"1000000", 1000000, nil},
		{"50000", 50000, nil},
		{"0", 0, nil},

		// Invalid cases
		{"-1", 0, ErrInvalidInput},
		{"abc", 0, ErrInvalidInput},
		{"12.34", 0, ErrInvalidInput},
		{"", 0, ErrInvalidInput},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := ParseDistance(test.input)

			if err != test.expectedErr {
				t.Errorf("expected error %v, got %v for input: %s", test.expectedErr, err, test.input)
			}
			if result != test.expected {
				t.Errorf("expected %d, got %d for input: %s", test.expected, result, test.input)
			}
		})
	}
}
