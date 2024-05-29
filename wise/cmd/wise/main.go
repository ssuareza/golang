package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ssuareza/golang/wise/pkg/wise"
)

const (
	apiKey = "dafe4c79-2163-4ea8-b019-6c22625ad635"
)

func main() {
	// cre
	client := wise.NewClient(apiKey)

	// get balance
	balance, err := client.GetBalance()
	if err != nil {
		log.Fatal(err)
	}

	// get rate
	rate, err := client.GetRate("USD", "EUR")
	if err != nil {
		log.Fatal(err)
	}

	// output
	fmt.Println("> Summary")
	fmt.Println("- Balance USD:", balance)
	fmt.Printf("- Balance EUR: %.2f\n", balance*rate)
	fmt.Println("- Rate USD to EUR:", rate)

	fmt.Printf("\n> Card transactions\n")

	// get transactions

	date := time.Now()

	// loop 6 months
	// layout := "2006-01-02T00:00:00.000Z"
	layout := "2006-01-02"
	for i := 0; i <= 5; i++ {
		date := date.AddDate(0, -i, 0)
		since := BeginningOfMonth(date).Format(layout) + "T00:00:00.000Z"
		until := EndOfMonth(date).Format(layout) + "T23:59:00.000Z"

		cardTransactions, err := client.GetCardTransactionsByRange(until, since)
		if err != nil {
			log.Fatal(err)
		}

		// sum transactions
		sumCardTransactions := client.SumCardTransactions(cardTransactions)

		// get max card transaction
		maxCardTransaction := client.GetMaxCardTransaction(cardTransactions)
		var maxCardTransactionTitle string
		maxCardTransactionTitle = strings.Replace(maxCardTransaction.Title, "<strong>", "", -1)
		maxCardTransactionTitle = strings.Replace(maxCardTransactionTitle, "</strong>", "", -1)
		maxCardTransactionParsed := fmt.Sprintf("%s %s", maxCardTransactionTitle, maxCardTransaction.Amount)

		// output
		fmt.Printf("- %s: %.2f EUR (%s)\n", date.Format("2006-01"), sumCardTransactions, maxCardTransactionParsed)
	}
}

// BeginningOfMonth returns the first day of the month
func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

// EndOfMonth returns the last day of the month
func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}
