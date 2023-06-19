package themoviedb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Results holds the results from the request
type Results struct {
	Results []Movie `json:"results"`
}

// Movie holds the movie structure
type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
}

// SearchMovie search movies based in a query string
func (client *Client) SearchMovie(query string, year string) ([]Movie, error) {
	query = strings.Replace(query, " ", "%20", -1)
	url := fmt.Sprintf("%s%s%s%s%s%s%s", client.URL, "/search/movie?include_adult=false&page=1&query=", query, "&year=", year, "&language=en-US&api_key=", client.APIKey)

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Movie{}, err
	}
	if res.StatusCode != 200 {
		return []Movie{}, fmt.Errorf("%s", "the request could not be satisfied.")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Movie{}, err
	}

	var r Results
	json.Unmarshal(body, &r)

	return r.Results, nil
}
