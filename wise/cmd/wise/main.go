package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/ssuareza/golang/wise/pkg/config"
	"github.com/ssuareza/golang/wise/pkg/wise"
)

var (
	label   *string
	verbose *bool
)

func init() {
	label = flag.String("label", "", "label to search")
	verbose = flag.Bool("verbose", false, "print transactions")

	if flag.Arg(0) == "help" {
		usage()
		os.Exit(0)
	}
}

// usage prints the usage of the program
func usage() {
	fmt.Println(`Usage: wise [options]

Options:
  -label=LABEL  label to search
  -verbose      print transactions
  -help         display this help and exit`)
}

func main() {
	flag.Parse()

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

		// Don't print maxTransactionParsed if we spend nothing
		if len(maxTransaction.Amount) > 0 {
			fmt.Printf("- %s: %.2f EUR (%s)\n", date.Format("2006-01"), sumTransactions, maxTransactionParsed)
		} else {
			fmt.Printf("- %s: %.2f EUR \n", date.Format("2006-01"), sumTransactions)
		}

		if *verbose {
			for _, transaction := range transactions.Activities {
				// process only payments completed
				if transaction.Type != "CARD_PAYMENT" {
					continue
				}

				var title string
				title = strings.Replace(transaction.Title, "<strong>", "", -1)
				title = strings.Replace(title, "</strong>", "", -1)
				date, _ := time.Parse(time.RFC3339, transaction.Date)
				fmt.Printf("-- %s: %s %s\n", date.Format(layout), title, transaction.Amount)
			}
		}
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
