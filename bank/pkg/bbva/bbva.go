package bbva

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ssuareza/golang/bank/pkg/csv"
	"github.com/xuri/excelize/v2"
)

var (
	errNoSuchFile = errors.New("no such file")
	errRows       = errors.New("error getting rows")
)

// BBVA
type BBVA struct {
	file  string
	sheet string
	month string
}

// New creates a new BBVA instance.
func New(file, sheet, month string) (*BBVA, error) {
	return &BBVA{file: file, sheet: sheet, month: month}, nil
}

// Process processes the file.
func (b *BBVA) Process() error {
	// get file
	f, err := b.getFile()
	if err != nil {
		return err
	}

	// get rows
	rows, err := b.getRows(f)
	if err != nil {
		return err
	}

	// filter rows
	filtered, err := b.filter(rows, b.month)
	if err != nil {
		return err
	}

	// // order rows
	// ordered, err := b.order(filtered)
	// if err != nil {
	// 	return err
	// }

	// create csv
	if err := csv.New("bbva", filtered); err != nil {
		return err
	}

	return nil
}

// getFile returns the file.
func (b *BBVA) getFile() (*excelize.File, error) {
	// open file
	f, err := excelize.OpenFile(b.file)
	if err != nil {
		return nil, errNoSuchFile
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
func (b *BBVA) getRows(file *excelize.File) ([][]string, error) {
	rows, err := file.GetRows(b.sheet)
	if err != nil {
		return nil, errRows
	}

	return rows, nil
}

// filter returns the filtered rows.
func (b *BBVA) filter(rows [][]string, filter string) (csv.Rows, error) {
	var filteredRows csv.Rows

	// loop through rows
	for _, row := range rows {
		for _, colCel := range row {
			// check if line starts with a date
			r := regexp.MustCompile(`^\d{2}/\d{2}/\d{4}`)
			if !r.MatchString(colCel) {
				continue
			}

			// filter by month
			if strings.Contains(colCel, fmt.Sprintf("/%s/", filter)) {
				// set date
				date, _ := time.Parse("02/01/2006", row[1])

				// set description
				filteredRow := csv.Row{
					Date:        date,
					Description: row[3],
					Amount:      row[5],
				}
				filteredRows = append(filteredRows, filteredRow)
				break
			}
		}
	}

	return filteredRows, nil
}

// order returns the ordered rows.
func (b *BBVA) order(rows csv.Rows) (csv.Rows, error) {
	for i, j := 0, len(rows)-1; i < j; i, j = i+1, j-1 {
		rows[i], rows[j] = rows[j], rows[i] //reverse the slice
	}

	return rows, nil
}
