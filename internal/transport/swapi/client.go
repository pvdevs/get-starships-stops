package swapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetStarships(ctx context.Context) ([]APIStarship, error) {
	var allStarships []APIStarship
	nextURL := fmt.Sprintf("%s/api/starships/", c.baseURL)

	for nextURL != "" {
		response, err := c.fetchStarshipsPage(ctx, nextURL)
		if err != nil {
			return nil, fmt.Errorf("fetch starships page: %w", err)
		}

		allStarships = append(allStarships, response.Results...)

		// Update URL for next page
		nextURL = response.Next
	}

	return allStarships, nil
}

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
