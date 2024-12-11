package main

import (
	"fmt"
	"log"

	"github.com/your-username/get-starships-stops/logic"
)

func main() {
	var input string
	fmt.Print("Enter distance in mega lights (MGLT): ")
	fmt.Scanln(&input)

	distance, err := logic.ParseDistance(input)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Valid distance entered: %d MGLT\n", distance)
}
