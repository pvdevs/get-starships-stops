package service

import (
	"context"
	"fmt"

	"github.com/pvdevs/get-starships-stops/internal/domain"
	"github.com/pvdevs/get-starships-stops/internal/parser"
)

// StarshipClient defines the interface for fetching starship data
// This allows to easily mock the client in tests
type StarshipClient interface {
	GetStarships(ctx context.Context) ([]domain.Starship, error)
}

// CalculatorService defines the interface for calculating starship stops
type CalculatorService interface {
	CalculateStops(ctx context.Context, distance int64) (map[string]int, error)
}

// Calculator handles the business logic for calculating required stops
type Calculator struct {
	client StarshipClient
}

// NewCalculator creates a new instance of Calculator with the provided client
func NewCalculator(client StarshipClient) *Calculator {
	return &Calculator{
		client: client,
	}
}

// CalculateStops determines how many stops each starship needs to make for a given distance
// It returns a map of starship names to their required number of stops
func (c *Calculator) CalculateStops(ctx context.Context, distance int64) (map[string]int, error) {
	starships, err := c.client.GetStarships(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch starships: %w", err)
	}

	results := make(map[string]int)

	for _, ship := range starships {
		if ship.MGLT <= 0 {
			results[ship.Name] = 0
			continue
		}

		hours, err := parser.ParseConsumables(ship.Consumables)
		if err != nil {
			continue
		}

		maxDistance := ship.MGLT * hours
		stops := int(distance) / maxDistance
		if int(distance)%maxDistance == 0 && stops > 0 {
			stops--
		}

		results[ship.Name] = stops
	}

	return results, nil
}
