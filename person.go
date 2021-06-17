package swapi

import "fmt"

// A Person is an individual person or character within the Star Wars universe.
type Person struct {
	Name         string   `json:"name,omitempty"`
	Height       string   `json:"height,omitempty"`
	Mass         string   `json:"mass,omitempty"`
	HairColor    string   `json:"hair_color,omitempty"`
	SkinColor    string   `json:"skin_color,omitempty"`
	EyeColor     string   `json:"eye_color,omitempty"`
	BirthYear    string   `json:"birth_year,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	Homeworld    string   `json:"homeworld,omitempty"`
}

// Person retrieves the person with the given id
func (c *Client) Person(id string) (Person, error) {
	req, err := c.newRequest(fmt.Sprintf("people/%s", id))
	if err != nil {
		return Person{}, err
	}

	var person Person

	if _, err = c.do(req, &person); err != nil {
		return Person{}, err
	}

	return person, nil
}
