package swapi

import "fmt"

// A Film is an single film.
type Film struct {
	title         string   `json:"title,omitempty"`
	EpisodeID     int      `json:"episode_id,omitempty"`
	OpeningCrawl  string   `json:"opening_crawl,omitempty"`
	Director      string   `json:"director,omitempty"`
	Producer      string   `json:"producer,omitempty"`
	Character     Person   `json:"characters,omitempty"`
	Planet        Planet   `json:"planets,omitempty"`
	Starship      Starship `json:"starships,omitempty"`
	Vehicle       Vehicle  `json:"vehicles,omitempty"`
	Species       Species  `json:"species,omitempty"`
}

// Film retrieves the film with the given id
func (c *Client) Film(id string) (Film, error) {
	req, err := c.newRequest(fmt.Sprintf("films/%s", id))
	if err != nil {
		return Film{}, err
	}

	var film Film

	if _, err = c.do(req, &film); err != nil {
		return Film{}, err
	}

	return film, nil
}
