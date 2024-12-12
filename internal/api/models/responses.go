package models

import (
	"sort"
	"strings"
)

// Result represents a single starship's calculation result
type Result struct {
	Name  string `json:"name"`
	Stops int    `json:"stops"`
}

// StopsResponse represents the complete API response
type StopsResponse struct {
	Distance int64    `json:"distance"`
	Results  []Result `json:"results"`
}

// SortResults sorts a slice of Results by stops (ascending), then alphabetically by name.
func SortResults(results []Result) {
	sort.Slice(results, func(i, j int) bool {
		if results[i].Stops != results[j].Stops {
			return results[i].Stops < results[j].Stops
		}
		return strings.ToLower(results[i].Name) < strings.ToLower(results[j].Name)
	})
}
