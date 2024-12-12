package ui

import (
	"fmt"
	"sort"
	"strings"
)

type Result struct {
	Name  string
	Stops int
}

func FormatResults(distance int64, stops map[string]int) []string {
	// Convert map to slice for sorting
	var results []Result
	for name, numStops := range stops {
		results = append(results, Result{
			Name:  name,
			Stops: numStops,
		})
	}

	// Sort by stops and then name
	sort.Slice(results, func(i, j int) bool {
		if results[i].Stops != results[j].Stops {
			return results[i].Stops < results[j].Stops
		}
		return strings.ToLower(results[i].Name) < strings.ToLower(results[j].Name)
	})

	// Find longest name for alignment
	maxNameLength := 0
	for _, r := range results {
		if len(r.Name) > maxNameLength {
			maxNameLength = len(r.Name)
		}
	}

	// Build output
	var output []string
	output = append(output, fmt.Sprintf("\nResults for %d MGLT:", distance))
	output = append(output, strings.Repeat("=", maxNameLength+20))

	// Add context message if all stops are 0
	allZero := true
	for _, r := range results {
		if r.Stops > 0 {
			allZero = false
			break
		}
	}
	if allZero {
		output = append(output, "\nNote: Distance is too short for any ship to require a resupply stop.")
		output = append(output, "All ships can complete this journey without stopping.\n")
	}

	// Format the result line with aligned columns
	for _, r := range results {
		stopWord := "stops"
		if r.Stops == 1 {
			stopWord = "stop "
		}

		// %-*s  : Left-align the starship name with a width of maxNameLength
		// |     : Adds a vertical bar separator
		// %3d   : Right-align the number of stops with a width of 3
		// %s    : Append the word "stop" or "stops" based on the number of stops
		line := fmt.Sprintf("%-*s | %3d %s", maxNameLength, r.Name, r.Stops, stopWord)
		output = append(output, line)
	}

	output = append(output, strings.Repeat("=", maxNameLength+20))
	return output
}
