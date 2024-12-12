package domain

// Starship represents the simplified internal structure used for business logic.
type Starship struct {
	Name        string // Name of the starship
	MGLT        int    // Distance the starship can travel in mega lights per hour
	Consumables string // Time the starship can travel without resupplying (e.g., "2 months")
}

type APIStarship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               string   `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 string   `json:"crew"`
	Passengers           string   `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     string   `json:"hyperdrive_rating"`
	MGLT                 string   `json:"MGLT"`
	StarshipClass        string   `json:"starship_class"`
	Pilots               []string `json:"pilots"`
	Films                []string `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  string   `json:"url"`
}
