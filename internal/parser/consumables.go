package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidConsumables = errors.New("invalid consumables format")
	ErrEmptyConsumables   = errors.New("empty consumables input")
)

// ParseConsumables converts a consumables string (e.g. "2 years", "6 months")
// into total hours of operation. Returns an error if the format is invalid.
func ParseConsumables(input string) (int, error) {
	if input == "" {
		return 0, ErrEmptyConsumables
	}

	parts := strings.Split(strings.TrimSpace(input), " ")
	if len(parts) != 2 {
		return 0, ErrInvalidConsumables
	}

	quantity, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, ErrInvalidConsumables
	}

	unit := strings.ToLower(parts[1])
	// Remove 's' from plural if present
	unit = strings.TrimSuffix(unit, "s")

	var hours int
	switch unit {
	case "year":
		hours = quantity * 365 * 24
	case "month":
		hours = quantity * 30 * 24 // simplified to 30 days per month
	case "week":
		hours = quantity * 7 * 24
	case "day":
		hours = quantity * 24
	default:
		return 0, fmt.Errorf("%w: unknown unit %s", ErrInvalidConsumables, unit)
	}

	return hours, nil
}
