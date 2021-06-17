package swapi

import "fmt"

// A Planet is a large mass, planet or planetoid in the Star Wars Universe, at the time of 0 ABY.
type Planet struct {
	Name           string   `json:"name,omitempty"`
	RotationPeriod string   `json:"rotation_period,omitempty"`
	OrbitalPeriod  string   `json:"orbital_period,omitempty"`
	Diameter       string   `json:"diameter,omitempty"`
	Climate        string   `json:"climate,omitempty"`
	Gravity        string   `json:"gravity,omitempty"`
	Terrain        string   `json:"terrain,omitempty"`
	SurfaceWater   string   `json:"surface_water,omitempty"`
	Population     string   `json:"population,omitempty"`
}

// Planet retrieves the planet with the given id
func (c *Client) Planet(id string) (Planet, error) {
	req, err := c.newRequest(fmt.Sprintf("planets/%s", id))
	if err != nil {
		return Planet{}, err
	}

	var planet Planet

	if _, err = c.do(req, &planet); err != nil {
		return Planet{}, err
	}

	return planet, nil
}
