// Package service provides the business logic for calculating starship stops
package service

import (
	"context"
	"fmt"

	"github.com/pvdevs/get-starships-stops/internal/domain"
	"github.com/pvdevs/get-starships-stops/internal/parser"
)

// StarshipClient defines the interface for fetching starship data
// This allows us to easily mock the client in tests
type StarshipClient interface {
	GetStarships(ctx context.Context) ([]domain.Starship, error)
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
// CalculateStops determines how many stops each starship needs to make for a given distance
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

// Helper method to calculate stops for a single starship
// This is useful when we want to calculate stops for specific ships
func (c *Calculator) calculateStopsForShip(shipName string, MGLT int, consumables string, distance int64) (int, error) {
	// Skip ships that can't move
	if MGLT <= 0 {
		return 0, nil
	}

	// Parse consumables to get hours of travel time
	hours, err := parser.ParseConsumables(consumables)
	if err != nil {
		return 0, fmt.Errorf("parse consumables: %w", err)
	}

	// Calculate maximum distance ship can travel before needing to stop
	maxDistance := MGLT * hours

	// Calculate number of stops needed
	stops := int(distance) / maxDistance
	if int(distance)%maxDistance == 0 && stops > 0 {
		stops--
	}

	return stops, nil
}
