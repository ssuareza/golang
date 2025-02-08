package wise

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ssuareza/golang/bank/pkg/config"
	"github.com/ssuareza/golang/bank/pkg/csv"
)

var (
	dateLayout        = "2006-01-02"
	errInvalidRequest = errors.New("invalid request")
	wiseAPIEndpoint   = "https://api.wise.com"
	wiseDateLayout    = "2006-01-02T15:04:05.000Z"
)

// Client contains the client configuration
type Client struct {
	ApiEndpoint string
	ApiKey      string
	ProfileID   string
	Token       string
	Client      *http.Client
}

// New creates a new client
func New(c config.Config) *Client {
	return &Client{
		ApiEndpoint: wiseAPIEndpoint,
		ApiKey:      c.ApiKey,
		ProfileID:   c.ProfileID,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Process processes the data.
func (c *Client) Process(date time.Time) error {
	// get transactions
	since := csv.BeginningOfMonth(date).Format(dateLayout) + "T00:00:00.000Z"
	fmt.Println(since)
	until := csv.EndOfMonth(date).Format(dateLayout) + "T23:59:00.000Z"
	fmt.Println(until)

	t, err := c.getTransactions(until, since)
	if err != nil {
		return err
	}

	// filter rows
	rows, err := c.filter(t)
	if err != nil {
		return err
	}

	// create csv
	if err := csv.New("wise", rows); err != nil {
		return err
	}

	// process transactions
	return nil
}

// GetTransactions retrieves the transactions by range.
func (c *Client) getTransactions(until, since string) (transactions, error) {
	var t transactions

	// set url
	url := fmt.Sprintf("%s%s%s%s%s%s%s%s", c.ApiEndpoint, "/v1/profiles/", c.ProfileID, "/activities?status=COMPLETED&until=", until, "&since=", since, "&size=100")

	fmt.Println(url)
	// new request
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)

	// make request
	res, err := c.Client.Do(req)
	if err != nil {
		return t, err
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		return t, errInvalidRequest
	}

	// get body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	// marshal json
	if err := json.Unmarshal(body, &t); err != nil {
		return t, err
	}

	return t, nil
}

// filter filters the transactions.
func (c *Client) filter(t transactions) (csv.Rows, error) {
	var rows csv.Rows

	for _, activity := range t.Activities {
		// set date
		date, _ := time.Parse(wiseDateLayout, activity.Date)

		// set description
		description := strings.Replace(activity.Title, "<strong>", "", -1)
		description = strings.Replace(description, "</strong>", "", -1)
		description = strings.Replace(description, "</strong>", "", -1)
		description = strings.Replace(description, "</strong>", "", -1)

		// set amount
		amount := strings.Replace(activity.Amount, ".", ",", -1)
		amount = strings.Replace(amount, " EUR", "", -1)

		// convert amount into negative
		if !strings.Contains(amount, "positive") {
			amount = fmt.Sprintf("-%s", amount)
		}

		// clean positive entries
		if strings.Contains(amount, "positive") {
			amount = strings.Replace(amount, "<positive>+ ", "", -1)
			amount = strings.Replace(amount, "</positive>", "", -1)
		}

		// add row
		rows = append(rows, csv.Row{
			Date:        date,
			Amount:      amount,
			Description: description,
		})
	}

	return rows, nil
}
