package parser

import (
	"errors"
	"strconv"
)

var (
	ErrNotPositiveInteger = errors.New("input must be a positive integer")
	ErrInputTooLarge      = errors.New("input is too large to process")
)

// ParseDistance converts a string input to an int64 distance value.
// Returns an error if the input is not a positive integer or is too large.
func ParseDistance(input string) (int64, error) {
	distance, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		// Check if the error is due to an overflow
		if errors.Is(err, strconv.ErrRange) {
			return 0, ErrInputTooLarge
		}
		return 0, ErrNotPositiveInteger
	}
	if distance < 0 {
		return 0, ErrNotPositiveInteger
	}
	return distance, nil
}
