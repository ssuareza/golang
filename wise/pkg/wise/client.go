package wise

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ssuareza/golang/wise/pkg/config"
)

var (
	errInvalidCredentials = errors.New("invalid credentials")
)

// Client contains the client configuration
type Client struct {
	ApiEndpoint string
	ApiKey      string
	ProfileID   string
	Token       string
	Client      *http.Client
}

// NewClient creates a new client
func NewClient(c config.Config) *Client {
	return &Client{
		ApiEndpoint: c.ApiEndpoint,
		ApiKey:      c.ApiKey,
		ProfileID:   c.ProfileID,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetBalance retrieves the balance of the account
func (c *Client) GetBalance() (float64, error) {
	// set url
	balanceURL := fmt.Sprintf("%s%s%s%s", c.ApiEndpoint, "/v4/profiles/", c.ProfileID, "/balances?types=STANDARD")

	// new request
	req, _ := http.NewRequest(http.MethodGet, balanceURL, nil)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)

	// make request
	res, err := c.Client.Do(req)
	if err != nil {
		return 0, err
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		return 0, errInvalidCredentials
	}

	// get body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	// marshal json
	var balances Balances
	if err := json.Unmarshal(body, &balances); err != nil {
		return 0, err
	}

	return balances[0].Amount.Value, nil
}

// GetRate retrieves the rate
func (c *Client) GetRate(source, target string) (float64, error) {
	// set url
	rateURL := fmt.Sprintf("%s%s%s%s%s", c.ApiEndpoint, "/v1/rates?source=", source, "&target=", target)

	// new request
	req, _ := http.NewRequest(http.MethodGet, rateURL, nil)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)

	// make request
	res, err := c.Client.Do(req)
	if err != nil {
		return 0, err
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		return 0, errInvalidCredentials
	}

	// get body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	// marshal json
	var rates Rates
	if err := json.Unmarshal(body, &rates); err != nil {
		return 0, err
	}

	return rates[0].Rate, nil
}

// GetTransactionsByRange retrieves the transactions by range
func (c *Client) GetTransactionsByRange(until, since string) (Transactions, error) {
	// set url
	transactionsURL := fmt.Sprintf("%s%s%s%s%s%s%s%s", c.ApiEndpoint, "/v1/profiles/", c.ProfileID, "/activities?status=COMPLETED&until=", until, "&since=", since, "&size=100")

	// new request
	req, _ := http.NewRequest(http.MethodGet, transactionsURL, nil)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)

	// make request
	res, err := c.Client.Do(req)
	if err != nil {
		return Transactions{}, err
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		return Transactions{}, errInvalidCredentials
	}

	// get body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	// marshal json
	var Transactions Transactions
	if err := json.Unmarshal(body, &Transactions); err != nil {
		return Transactions, err
	}

	return Transactions, nil
}

// SumTransactions sums the transactions
func (c *Client) SumTransactions(Transactions Transactions) float64 {
	var sum float64

	// loop activities
	for _, activity := range Transactions.Activities {
		// process only card payments completed
		if activity.Type != "CARD_PAYMENT" {
			continue
		}

		amount := strings.Fields(activity.Amount)
		amountFloat, _ := strconv.ParseFloat(amount[0], 64)
		sum += float64(amountFloat)
	}

	return sum
}

// GetMaxTransaction retrieves the max transaction
func (c *Client) GetMaxTransaction(transactions Transactions) Transaction {
	var max Transaction

	// loop activities
	for i, activity := range transactions.Activities {
		// process only payments completed
		if activity.Type != "CARD_PAYMENT" {
			continue
		}

		// check if is the first value
		if i == 0 {
			max = activity
			continue
		}

		// ignore positive entries
		if strings.Contains(activity.Amount, "+") {
			continue
		}

		// check if max
		activityAmount := strings.Fields(activity.Amount)
		activityAmountFloat, _ := strconv.ParseFloat(activityAmount[0], 64)

		maxAmount := strings.Fields(max.Amount)
		maxAmountFloat, _ := strconv.ParseFloat(maxAmount[0], 64)

		if activityAmountFloat > maxAmountFloat {
			max = activity
		}
	}

	return max
}

// FilterTransactions filters the transactions by label
func (c *Client) FilterTransactionsByLabel(transactions Transactions, label string) Transactions {
	var filtered Transactions

	for _, activity := range transactions.Activities {
		if strings.Contains(activity.Title, label) {
			filtered.Activities = append(filtered.Activities, activity)
		}
	}

	return filtered
}
