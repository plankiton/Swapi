package swapi

import "fmt"

// A Vehicle is a single transport craft that does not have hyperdrive capability.
type Vehicle struct {
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
	VehicleClass         string   `json:"vehicle_class,omitempty"`
}

// Vehicle retrieves the vehicle with the given id
func (c *Client) Vehicle(id string) (Vehicle, error) {
	req, err := c.newRequest(fmt.Sprintf("vehicles/%s", id))
	if err != nil {
		return Vehicle{}, err
	}

	var vehicle Vehicle

	if _, err = c.do(req, &vehicle); err != nil {
		return Vehicle{}, err
	}

	return vehicle, nil
}
