package logic

import (
	"errors"
	"strconv"
)

// Predefined errors for mapping validation
var (
	ErrInvalidMGLT        = errors.New("invalid MGLT: must be a numeric value")
	ErrInvalidConsumables = errors.New("invalid consumables: must follow the format '<value> <unit>'")
)

// MapAPIStarshipToStarship converts an APIStarship to a simplified Starship structure.
func MapAPIStarshipToStarship(apiStarship APIStarship) (Starship, error) {
	// Parse MGLT (string to int)
	mglt, err := strconv.Atoi(apiStarship.MGLT)
	if err != nil {
		return Starship{}, ErrInvalidMGLT
	}

	// Validate Consumables (not parsed here, but must be non-empty)
	if apiStarship.Consumables == "" {
		return Starship{}, ErrInvalidConsumables
	}

	// Return the simplified Starship structure
	return Starship{
		Name:        apiStarship.Name,
		MGLT:        mglt,
		Consumables: apiStarship.Consumables,
	}, nil
}
