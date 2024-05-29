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

// GetCardTransactionsByRange retrieves the transactions by range
func (c *Client) GetCardTransactionsByRange(until, since string) (CardTransactions, error) {
	// set url
	cardTransactionsURL := fmt.Sprintf("%s%s%s%s%s%s%s%s", c.ApiEndpoint, "/v1/profiles/", c.ProfileID, "/activities?status=COMPLETED&until=", until, "&since=", since, "&size=100")

	// new request
	req, _ := http.NewRequest(http.MethodGet, cardTransactionsURL, nil)
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)

	// make request
	res, err := c.Client.Do(req)
	if err != nil {
		return CardTransactions{}, err
	}

	// check status code
	if res.StatusCode != http.StatusOK {
		return CardTransactions{}, errInvalidCredentials
	}

	// get body
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	// marshal json
	var cardTransactions CardTransactions
	if err := json.Unmarshal(body, &cardTransactions); err != nil {
		return cardTransactions, err
	}

	return cardTransactions, nil
}

// SumCardTransactions sums the card transactions
func (c *Client) SumCardTransactions(cardTransactions CardTransactions) float64 {
	var sum float64

	// loop activities
	for _, activity := range cardTransactions.Activities {
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

// GetMaxCardTransaction retrieves the max card transaction
func (c *Client) GetMaxCardTransaction(cardTransactions CardTransactions) CardTransaction {
	var max CardTransaction

	// loop activities
	for i, activity := range cardTransactions.Activities {
		// process only card payments completed
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
