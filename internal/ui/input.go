package ui

import (
	"fmt"
	"strings"
)

// ReadDistance prompts the user to enter a distance in MGLT and reads the input.
func ReadDistance() (string, error) {
	var input string

	// Prompt the user for input
	fmt.Print("\nEnter distance in mega lights (MGLT): ")

	// Read the input and handle potential errors
	_, err := fmt.Scanln(&input)
	if err != nil {
		return "", fmt.Errorf("failed to read input: %w", err)
	}

	// Trim whitespace from the input and return it
	return strings.TrimSpace(input), nil
}
