package main

import (
	"context"
	"os"

	"github.com/pvdevs/get-starships-stops/internal/parser"
	"github.com/pvdevs/get-starships-stops/internal/service"
	"github.com/pvdevs/get-starships-stops/internal/transport/swapi"
	"github.com/pvdevs/get-starships-stops/internal/ui"
)

func main() {
	// Get user input
	input, err := ui.ReadDistance()
	if err != nil {
		ui.PrintError(err)
		os.Exit(1)
	}

	// Parse distance
	distance, err := parser.ParseDistance(input)
	if err != nil {
		ui.PrintError(err)
		os.Exit(1)
	}

	ui.PrintCalculating(distance)

	// Initialize services and calculate
	client := swapi.NewClient(swapi.ClientConfig{
		BaseURL: "https://swapi.dev",
	})
	calculator := service.NewCalculator(client)

	stops, err := calculator.CalculateStops(context.Background(), distance)
	if err != nil {
		ui.PrintError(err)
		os.Exit(1)
	}

	// Format and print results
	formattedResults := ui.FormatResults(distance, stops)
	ui.PrintResults(formattedResults)
}
