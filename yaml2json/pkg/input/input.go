package input

import (
	"errors"
	"io"
	"os"
)

var errMissingInput = errors.New("missing input")

// Read reads a input from a file or stdin.
func Read() (input []byte, err error) {
	// read from file if provided
	if len(os.Args) > 1 {
		if input, err = os.ReadFile(os.Args[1]); err != nil {
			return []byte{}, err
		}

		return input, nil
	}

	// read from stdin if no file provided
	if stdinStat, _ := os.Stdin.Stat(); (stdinStat.Mode() & os.ModeCharDevice) == 0 {
		if input, err = io.ReadAll(os.Stdin); err != nil {
			return []byte{}, err
		}
		return input, nil
	}

	return []byte{}, errMissingInput
}
