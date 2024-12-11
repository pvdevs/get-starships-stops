package logic

import (
	"errors"
	"strconv"
)

// Predefined errors
var (
	ErrInvalidInput = errors.New("invalid input: must be a positive integer")
)

// ParseDistance parses the input and returns the distance in MGLT or an error
func ParseDistance(input string) (int64, error) {
	distance, err := strconv.ParseInt(input, 10, 64)
	if err != nil || distance < 0 {
		return 0, ErrInvalidInput
	}
	return distance, nil
}
