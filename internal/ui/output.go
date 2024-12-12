package ui

import "fmt"

// PrintResults prints each line of the provided results.
func PrintResults(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

// PrintError prints an error message.
func PrintError(err error) {
	fmt.Printf("Error: %v\n", err)
}

// PrintCalculating prints a message indicating that stop calculations are in progress.
func PrintCalculating(distance int64) {
	fmt.Printf("Calculating stops for distance: %d MGLT...\n", distance)
}