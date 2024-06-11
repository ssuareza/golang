package csv

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	size      = 6   // size is the number of lines expected by day
	delimiter = ';' // delimiter is the csv separator
)

var errNoRows = errors.New("no rows found")

// Row represents a row in the file.
type Row struct {
	Date        time.Time
	Description string
	Amount      string
}

// Rows represents a slice of Rows.
type Rows []Row

// New generates a new csv file.
func New(name string, rows Rows) error {
	// ignore empty rows
	if len(rows) == 0 {
		return errNoRows
	}

	// set file
	file := fmt.Sprintf("/tmp/%s.csv", name)

	// create file
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func() {
		err := f.Close()
		if err != nil {
			return
		}
	}()

	// create writer
	w := csv.NewWriter(f)

	// set csv separator
	w.Comma = delimiter

	// write rows
	// loop the whole month
	for i := 1; i <= EndOfMonth(rows[len(rows)-1].Date).Day(); i++ {
		// set size
		s := size

		// write row
		for _, row := range rows {
			if row.Date.Day() == i {
				w.Write([]string{row.Description, row.Amount})
				s--
			}
		}

		// write empty lines
		if s != 0 {
			for i := 0; i < s; i++ {
				w.Write([]string{"", ""})
			}
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}

	// output
	fmt.Printf("CSV file created successfully on %s\n", file)
	return nil
}

// BeginningOfMonth returns the fist day of the month.
func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}

// EndOfMonth returns the last day of the month.
func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}
