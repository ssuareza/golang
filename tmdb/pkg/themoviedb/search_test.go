package themoviedb

import (
	"testing"
)

const (
	MovieResponse = `{
		"page": 1,
		"total_results": 54,
		"total_pages": 3,
		"results": [
		  {
			"popularity": 363.674,
			"vote_count": 6594,
			"video": false,
			"poster_path": "/udDclJoHjfjb8Ekgsd4FDteOkCU.jpg",
			"id": 475557,
			"adult": false,
			"backdrop_path": "/n6bUvigpRFqSwmPp1m2YADdbRBc.jpg",
			"original_language": "en",
			"original_title": "Joker",
			"genre_ids": [
			  80,
			  18,
			  53
			],
			"title": "Joker",
			"vote_average": 8.3,
			"overview": "During the 1980s, a failed stand-up comedian is driven insane and turns to a life of crime and chaos in Gotham City while becoming an infamous psychopathic crime figure.",
			"release_date": "2019-10-02"
		  },
		  {
			"popularity": 3.469,
			"id": 129507,
			"video": false,
			"vote_count": 23,
			"vote_average": 3.4,
			"title": "Joker",
			"release_date": "2012-08-31",
			"original_language": "hi",
			"original_title": "Joker",
			"genre_ids": [
			  35,
			  878
			],
			"backdrop_path": "/7hsepp47nWm5cPgC1bEFJQK3Rk8.jpg",
			"adult": false,
			"overview": "In 1947 when the maps of India and Pakistan were being drawn, an oversight ensured that the village of Paglapur didn't find a place in either country. Over 60 years later, Paglapur is isolated, and in need of help. The residents seek help, but as a direct result of being left off the map, they find noone willing to accept Paglapur, and its problems, as their jurisdiction. The villagers decide they need to draw media attention to Paglapur, and in so doing, gain acceptance and help from those who had previously denied them.",
			"poster_path": "/ql1DUWjY6Y6zEbVFS83s20SBo8J.jpg"
		  }
		]
	}`
)

func TestSearchMovie(t *testing.T) {
	server := TestingHTTPServer(MovieResponse)
	defer server.Close()

	client := NewClient("xxxxxxxxxxxxxxxxxxxxxx")
	client.URL = server.URL

	movie, err := client.SearchMovie("Joker", "")
	if err != nil || movie[0].Title != "Joker" {
		t.Fail()
	}
}
