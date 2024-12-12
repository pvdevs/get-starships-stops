package swapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pvdevs/get-starships-stops/internal/domain"
)

// Client handles all communication with the SWAPI API
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new SWAPI client instance
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// // GetStarships fetches and returns all starships from the SWAPI API
// // Returns domain.Starship objects instead of API responses
func (c *Client) GetStarships(ctx context.Context) ([]domain.Starship, error) {
	var allStarships []domain.Starship
	nextURL := fmt.Sprintf("%s/api/starships/", c.baseURL)

	for nextURL != "" {
		response, err := c.fetchStarshipsPage(ctx, nextURL)
		if err != nil {
			return nil, fmt.Errorf("fetch starships page: %w", err)
		}

		for _, apiShip := range response.Results {
			ship, err := apiToDomainStarship(apiShip)
			if err != nil {
				if err.Error() != "skip" { // Only log real errors, not skipped ships
					fmt.Printf("Warning: could not process ship %s: %v\n", apiShip.Name, err)
				}
				continue
			}
			allStarships = append(allStarships, ship)
		}

		nextURL = response.Next
	}

	return allStarships, nil
}

// apiToDomainStarship converts an APIStarship to a domain.Starship
func apiToDomainStarship(apiShip APIStarship) (domain.Starship, error) {
	// Skip ships with non-numeric MGLT values
	if apiShip.MGLT == "unknown" || apiShip.MGLT == "n/a" {
		return domain.Starship{}, fmt.Errorf("skip")
	}

	mglt, err := strconv.Atoi(apiShip.MGLT)
	if err != nil {
		return domain.Starship{}, fmt.Errorf("invalid MGLT format")
	}

	return domain.Starship{
		Name:        apiShip.Name,
		MGLT:        mglt,
		Consumables: apiShip.Consumables,
	}, nil
}

// fetchStarshipsPage fetches a single page of starship data from the API
func (c *Client) fetchStarshipsPage(ctx context.Context, url string) (*StarshipsResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var starshipsResp StarshipsResponse
	if err := json.NewDecoder(resp.Body).Decode(&starshipsResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &starshipsResp, nil
}
