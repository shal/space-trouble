package spacex

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/opencars/space-trouble/pkg/domain/model"
)

type Client struct {
	c *http.Client
}

func NewClient() *Client {
	return &Client{
		c: &http.Client{},
	}
}

func (c *Client) FindByLaunchpadID(ctx context.Context, launchpadID string) (*model.Launchpad, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.spacexdata.com/v4/launchpads/"+launchpadID, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	var result LaunchpadResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	// for _, r := range result.Launches {
	// 	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.spacexdata.com/v4/launch/"+r, nil)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	_, err := c.c.Do(req)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// TODO: Finish implementation.
	return nil, nil
}

type LaunchpadResult struct {
	Images struct {
		Large []string `json:"large"`
	} `json:"images"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Locality        string   `json:"locality"`
	Region          string   `json:"region"`
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	LaunchAttempts  int      `json:"launch_attempts"`
	LaunchSuccesses int      `json:"launch_successes"`
	Rockets         []string `json:"rockets"`
	Timezone        string   `json:"timezone"`
	Launches        []string `json:"launches"`
	Status          string   `json:"status"`
	Details         string   `json:"details"`
	ID              string   `json:"id"`
}
