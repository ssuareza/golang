package themoviedb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Token contains TMDb token
type Token struct {
	Success    bool   `json:"success"`
	Expiration string `json:"expires_at"`
	Token      string `json:"request_token"`
}

// NewToken creates a temporary request token that can be used to validate a TMDb user
func (client *Client) NewToken() (string, error) {
	url := fmt.Sprintf("%s%s%s", client.URL, "/authentication/token/new?api_key=", client.APIKey)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("%s", "the request could not be satisfied.")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var t Token
	json.Unmarshal(body, &t)

	return t.Token, nil
}
