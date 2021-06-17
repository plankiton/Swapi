package swapi

import "fmt"

// A Species is a type of person or character within the Star Wars Universe.
type Species struct {
	Name            string   `json:"name,omitempty"`
	Classification  string   `json:"classification,omitempty"`
	Designation     string   `json:"designation,omitempty"`
	AverageHeight   string   `json:"average_height,omitempty"`
	SkinColors      string   `json:"skin_colors,omitempty"`
	HairColors      string   `json:"hair_colors,omitempty"`
	EyeColors       string   `json:"eye_colors,omitempty"`
	AverageLifespan string   `json:"average_lifespan,omitempty"`
	Homeworld       string   `json:"homeworld,omitempty"`
	Language        string   `json:"language,omitempty"`
}

// Species retrieves the species with the given id
func (c *Client) Species(id string) (Species, error) {
	req, err := c.newRequest(fmt.Sprintf("species/%s", id))
	if err != nil {
		return Species{}, err
	}

	var species Species

	if _, err = c.do(req, &species); err != nil {
		return Species{}, err
	}

	return species, nil
}
