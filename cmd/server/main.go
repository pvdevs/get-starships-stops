package main

import (
	"fmt"
	"log"

	"github.com/pvdevs/get-starships-stops/internal/parser/distance"
)

func main() {
	var input string
	fmt.Print("Enter distance in mega lights (MGLT): ")
	fmt.Scanln(&input)

	distance, err := distance.ParseDistance(input)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Valid distance entered: %d MGLT\n", distance)
}
