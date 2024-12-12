package logic

import (
	"errors"
	"testing"
)

func TestMapAPIStarshipToStarship(t *testing.T) {
	tests := []struct {
		name        string
		apiStarship APIStarship
		expected    Starship
		expectedErr error
	}{
		// Valid case
		{
			"Valid starship",
			APIStarship{
				Name:                 "Death Star",
				Model:                "DS-1 Orbital Battle Station",
				Manufacturer:         "Imperial Department of Military Research, Sienar Fleet Systems",
				CostInCredits:        "1000000000000",
				Length:               "120000",
				MaxAtmospheringSpeed: "n/a",
				Crew:                 "342,953",
				Passengers:           "843,342",
				CargoCapacity:        "1000000000000",
				Consumables:          "3 years",
				HyperdriveRating:     "4.0",
				MGLT:                 "10",
				StarshipClass:        "Deep Space Mobile Battlestation",
				Created:              "2014-12-10T16:36:50.509000Z",
				Edited:               "2014-12-20T21:26:24.783000Z",
				URL:                  "https://swapi.dev/api/starships/9/",
			},
			Starship{
				Name:        "CR90 corvette",
				MGLT:        60,
				Consumables: "1 year",
			},
			nil,
		},
		// Invalid MGLT
		{
			"Invalid MGLT",
			APIStarship{
				Name:        "CR90 corvette",
				MGLT:        "invalid",
				Consumables: "1 year",
			},
			Starship{},
			ErrInvalidMGLT,
		},
		// Missing Consumables
		{
			"Missing consumables",
			APIStarship{
				Name:        "CR90 corvette",
				MGLT:        "60",
				Consumables: "",
			},
			Starship{},
			ErrInvalidConsumables,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := MapAPIStarshipToStarship(test.apiStarship)

			// Check for errors
			if !errors.Is(err, test.expectedErr) {
				t.Errorf("expected error %v, got %v", test.expectedErr, err)
			}

			// Check for correct result
			if result != test.expected {
				t.Errorf("expected result %v, got %v", test.expected, result)
			}
		})
	}
}
