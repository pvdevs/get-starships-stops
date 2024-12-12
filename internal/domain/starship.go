package logic

// Starship represents the simplified internal structure used for business logic.
type Starship struct {
	Name        string // Name of the starship
	MGLT        int    // Distance the starship can travel in mega lights per hour
	Consumables string // Time the starship can travel without resupplying (e.g., "2 months")
}
