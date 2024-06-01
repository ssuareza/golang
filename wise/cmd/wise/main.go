package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ssuareza/golang/wise/pkg/config"
	"github.com/ssuareza/golang/wise/pkg/wise"
)

var label *string

func init() {
	label = flag.String("label", "", "label to search")
}

func main() {
	// get config
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	// create client
	client := wise.NewClient(config)

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

		// get transactions
		transactions, err := client.GetTransactionsByRange(until, since)
		if err != nil {
			log.Fatal(err)
		}

		// filter transactions
		flag.Parse()
		if len(*label) != 0 {
			transactions = client.FilterTransactionsByLabel(transactions, *label)
		}

		// don't process transactions if is empty
		if len(transactions.Activities) == 0 {
			fmt.Printf("- %s: 0.0 EUR\n", date.Format("2006-01"))
			continue
		}

		// sum transactions
		sumTransactions := client.SumTransactions(transactions)

		// get max transaction
		maxTransaction := client.GetMaxTransaction(transactions)
		var maxTransactionTitle string
		maxTransactionTitle = strings.Replace(maxTransaction.Title, "<strong>", "", -1)
		maxTransactionTitle = strings.Replace(maxTransactionTitle, "</strong>", "", -1)
		maxTransactionParsed := fmt.Sprintf("%s %s", maxTransactionTitle, maxTransaction.Amount)

		// output
		fmt.Printf("- %s: %.2f EUR (%s)\n", date.Format("2006-01"), sumTransactions, maxTransactionParsed)
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
