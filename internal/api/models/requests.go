package models

// StopsRequest represents the payload for a request to calculate starship stops.
type StopsRequest struct {
	Distance string `json:"distance"` // Distance to travel in mega lights (MGLT)
}
