package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ssuareza/golang/bank/pkg/bbva"
	"github.com/ssuareza/golang/bank/pkg/config"
	"github.com/ssuareza/golang/bank/pkg/ing"
	"github.com/ssuareza/golang/bank/pkg/wise"
)

const (
	bbvaSheet = "Informe BBVA"
	ingSheet  = "Movimientos"
)

var (
	name *string
	file *string
	date *string

	errInvalidDate = fmt.Errorf("invalid date")
)

func init() {
	name = flag.String("name", "", "name of the bank")
	file = flag.String("file", "", "file to process")
	date = flag.String("date", "", "date to search")

	if flag.Arg(0) == "help" {
		usage()
		os.Exit(0)
	}

	flag.Parse()
}

// usage prints the usage of the program
func usage() {
	fmt.Println(`Usage: bank --name=<bbva|ing|wise> --file=<file> --date=<date>

Options:

  -name=<bbva|ing|wise>  name of the bank
  -file=file             file to process
  -date=date             date to search
  -help                  display this help and exit`)
}

func main() {
	// check params
	if *name == "" || *date == "" {
		usage()
		os.Exit(0)
	}

	// set date
	d, err := time.Parse("02/01/2006", *date)
	if err != nil {
		log.Fatal(errInvalidDate)
	}

	*date = d.Format("2006-01-02")

	// check file
	// process by name
	switch *name {
	case "bbva":
		checkFileParameter()

		bank, err := bbva.New(*file, bbvaSheet, d)
		if err != nil {
			log.Fatal(err)
		}

		if err := bank.Process(); err != nil {
			log.Fatal(err)
		}

	case "ing":
		checkFileParameter()

		ing, err := ing.New(*file, ingSheet, d)
		if err != nil {
			log.Fatal(err)
		}

		if err := ing.Process(); err != nil {
			log.Fatal(err)
		}

	case "wise":
		// get config
		config, err := config.New()
		if err != nil {
			log.Fatal(err)
		}

		wise := wise.New(config)

		if err := wise.Process(d); err != nil {
			log.Fatal(err)
		}

	default:
		usage()
	}
}

// checkFileParameter checks if the file parameter is valid
func checkFileParameter() {
	if *file == "" {
		usage()
		os.Exit(0)
	}
}
