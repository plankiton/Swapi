package swapi

import "fmt"

// A Starship is a single transport craft that has hyperdrive capability.
type Starship struct {
	Name                 string   `json:"name,omitempty"`
	Model                string   `json:"model,omitempty"`
	Manufacturer         string   `json:"manufacturer,omitempty"`
	CostInCredits        string   `json:"cost_in_credits,omitempty"`
	Length               string   `json:"length,omitempty"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed,omitempty"`
	Crew                 string   `json:"crew,omitempty"`
	Passengers           string   `json:"passengers,omitempty"`
	CargoCapacity        string   `json:"cargo_capacity,omitempty"`
	Consumables          string   `json:"consumables,omitempty"`
	HyperdriveRating     string   `json:"hyperdrive_rating,omitempty"`
	MGLT                 string   `json:"MGLT,omitempty"`
	StarshipClass        string   `json:"starship_class,omitempty"`
}

// Starship retrieves the starship with the given id
func (c *Client) Starship(id string) (Starship, error) {
	req, err := c.newRequest(fmt.Sprintf("starships/%s", id))
	if err != nil {
		return Starship{}, err
	}

	var starship Starship

	if _, err = c.do(req, &starship); err != nil {
		return Starship{}, err
	}

	return starship, nil
}
