package ing

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ssuareza/golang/bank/pkg/csv"
	"github.com/xuri/excelize/v2"
)

const decimalSeparator = ","

var (
	errRows = errors.New("error getting rows")
)

// ING
type ING struct {
	file  string
	sheet string
	month string
}

// New creates a new ING instance.
func New(file string, sheet string, date time.Time) (*ING, error) {
	// get month
	month := date.Format("01")

	// create instance
	return &ING{file: file, sheet: sheet, month: month}, nil
}

// Process processes the file.
func (i *ING) Process() error {
	// get file
	f, err := i.getFile()
	if err != nil {
		return err
	}

	// get rows
	rows, err := i.getRows(f)
	if err != nil {
		return err
	}

	// filter rows
	filtered, err := i.filter(rows, i.month)
	if err != nil {
		return err
	}

	// create csv
	if err := csv.New("ing", filtered); err != nil {
		return err
	}

	return nil
}

// getFile returns the file.
func (i *ING) getFile() (*excelize.File, error) {
	// open file
	f, err := excelize.OpenFile(i.file)
	if err != nil {
		return nil, err
	}
	defer func() {
		// close the file
		if err := f.Close(); err != nil {
			return
		}
	}()

	return f, nil
}

// getRows returns the rows.
func (i *ING) getRows(file *excelize.File) ([][]string, error) {
	rows, err := file.GetRows(i.sheet)
	if err != nil {
		return nil, errRows
	}

	return rows, nil
}

// filter returns the filtered rows.
func (i *ING) filter(rows [][]string, filter string) (csv.Rows, error) {
	var filteredRows csv.Rows

	// loop through rows
	for _, row := range rows {
		for _, cell := range row {
			// check if line starts with a date
			r := regexp.MustCompile(`^\d{2}/\d{2}/\d{4}`)
			if !r.MatchString(cell) {
				continue
			}

			// filter by month
			if strings.Contains(cell, fmt.Sprintf("/%s/", filter)) {
				// set date
				date, _ := time.Parse("02/01/2006", row[0])

				// set amount
				amount := strings.Replace(row[5], ".", decimalSeparator, -1)

				// set description
				filteredRow := csv.Row{
					Date:        date,
					Description: row[3],
					Amount:      amount,
				}
				filteredRows = append(filteredRows, filteredRow)
				break
			}
		}
	}

	return filteredRows, nil
}
