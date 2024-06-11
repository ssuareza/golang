package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ssuareza/golang/bank/pkg/bbva"
	"github.com/ssuareza/golang/bank/pkg/ing"
)

const (
	bbvaSheet = "Informe BBVA"
	ingSheet  = "Movimientos"
)

var (
	name  *string
	file  *string
	month *string
)

func init() {
	name = flag.String("name", "", "name of the bank")
	file = flag.String("file", "", "file to process")
	month = flag.String("month", "", "month to search")

	if flag.Arg(0) == "help" {
		usage()
		os.Exit(0)
	}

	flag.Parse()
}

// usage prints the usage of the program
func usage() {
	fmt.Println(`Usage: bank --name=<bbva|ing> --file=<file> --month=<month>

Options:

  -name=<bbva|ing>   name of the bank
  -file=file         file to process
  -month=month       month to search
  -help              display this help and exit`)
}

func main() {
	// check params
	if *name == "" || *file == "" || *month == "" {
		usage()
		os.Exit(0)
	}

	// process by name
	switch *name {
	case "bbva":
		bank, err := bbva.New(*file, bbvaSheet, *month)
		if err != nil {
			log.Fatal(err)
		}

		if err := bank.Process(); err != nil {
			log.Fatal(err)
		}

	case "ing":
		ing, err := ing.New(*file, ingSheet, *month)
		if err != nil {
			log.Fatal(err)
		}

		if err := ing.Process(); err != nil {
			log.Fatal(err)
		}

	default:
		usage()
	}
}
