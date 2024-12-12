package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/pvdevs/get-starships-stops/internal/parser"
	"github.com/pvdevs/get-starships-stops/internal/service"
	"github.com/pvdevs/get-starships-stops/internal/transport/swapi"
)

func main() {
	// Set up logger with timestamp
	log.SetFlags(log.Ldate | log.Ltime)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Initialize SWAPI client
	client := swapi.NewClient("https://swapi.dev")

	// Create calculator service
	calculator := service.NewCalculator(client)

	// Get distance from user input
	fmt.Print("\nEnter distance in mega lights (MGLT): ")
	var input string
	fmt.Scanln(&input)

	// Parse the distance
	distance, err := parser.ParseDistance(input)
	if err != nil {
		log.Fatalf("Error: %v\nPlease enter a valid positive number.", err)
	}

	fmt.Printf("\nCalculating stops for distance: %d MGLT...\n", distance)

	// Calculate stops for all starships
	stops, err := calculator.CalculateStops(ctx, distance)
	if err != nil {
		log.Fatalf("Error calculating stops: %v", err)
	}

	// Print results
	printResults(distance, stops)
}

// printResults formats and displays the calculation results
func printResults(distance int64, stops map[string]int) {
	if len(stops) == 0 {
		fmt.Println("\nNo starships found!")
		os.Exit(1)
	}

	// Convert map to slice for sorting
	type Result struct {
		Name  string
		Stops int
	}
	var results []Result

	for name, numStops := range stops {
		results = append(results, Result{
			Name:  name,
			Stops: numStops,
		})
	}

	// Sort by number of stops (and then by name for equal stops)
	sort.Slice(results, func(i, j int) bool {
		if results[i].Stops != results[j].Stops {
			return results[i].Stops < results[j].Stops
		}
		return strings.ToLower(results[i].Name) < strings.ToLower(results[j].Name)
	})

	// Find longest name for proper alignment
	maxNameLength := 0
	for _, result := range results {
		if len(result.Name) > maxNameLength {
			maxNameLength = len(result.Name)
		}
	}

	// Print header
	fmt.Printf("\nResults for %d MGLT:\n", distance)
	fmt.Println(strings.Repeat("=", maxNameLength+20))

	// Print results
	for _, result := range results {
		stopsText := "stops"
		if result.Stops == 1 {
			stopsText = "stop"
		}
		fmt.Printf("%-*s | %d %s\n",
			maxNameLength,
			result.Name,
			result.Stops,
			stopsText,
		)
	}
	fmt.Println(strings.Repeat("=", maxNameLength+20))
}
